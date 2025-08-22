#!/usr/bin/env python3
"""
Qdrant-based Knowledge Base with better embeddings.
Fast, efficient, and optimized for LLM context retrieval.
"""

import os
import sys
import json
import hashlib
import argparse
import time
from pathlib import Path
from typing import List, Dict, Any, Optional, Tuple
from datetime import datetime
from dataclasses import dataclass, asdict
from enum import Enum
from contextlib import contextmanager

# Qdrant client
from qdrant_client import QdrantClient
from qdrant_client.models import (
    Distance, VectorParams, PointStruct, 
    Filter, FieldCondition, MatchValue,
    SearchRequest, SearchParams
)

# Embeddings
import requests
import numpy as np
from typing import Union

# Text processing
import tiktoken


class SourceType(Enum):
    INTERNAL = "internal"
    HARVESTED = "harvested"
    GENERATED = "generated"
    USER = "user"


@dataclass
class DocumentMetadata:
    """Enhanced metadata for documents."""
    file_path: str
    chunk_index: int
    source_type: str
    document_category: Optional[str] = None
    agent_creator: Optional[str] = None
    confidence_score: float = 1.0
    indexed_at: Optional[str] = None
    last_modified: Optional[str] = None
    file_hash: Optional[str] = None
    tags: Optional[str] = None  # Comma-separated string for Qdrant
    
    def to_dict(self):
        """Convert to dictionary for Qdrant payload."""
        return {k: v for k, v in asdict(self).items() if v is not None}


class QdrantKB:
    """Fast, efficient knowledge base using Qdrant."""
    
    # Best quality embedding model for superior semantic understanding
    EMBEDDING_MODEL = "BAAI/bge-large-en-v1.5"  # 1024 dims, best quality
    EMBEDDING_DIM = 1024
    
    # Collections
    COLLECTIONS = {
        "docs": "General documentation",
        "architecture": "Architecture decisions and patterns",
        "zettelkasten": "Learning notes and insights",
        "code": "Source code and implementations",
        "harvested": "External content from web"
    }
    
    def __init__(self, qdrant_url: str = None, embeddings_url: str = None, project_root: str = None, read_only: bool = False):
        """Initialize Qdrant KB."""
        if project_root is None:
            project_root = Path(__file__).parent.parent
        
        self.project_root = Path(project_root)
        self.kb_path = self.project_root / ".vector_db"
        self.read_only = read_only
        
        # Qdrant server URL - can be overridden via environment variable
        if qdrant_url is None:
            qdrant_url = os.getenv('QDRANT_URL', 'http://localhost:6333')
        
        # Embeddings server URL - use TEI if available, fallback to local
        if embeddings_url is None:
            embeddings_url = os.getenv('EMBEDDINGS_URL', 'http://localhost:8080')
        self.embeddings_url = embeddings_url
        self.use_embeddings_server = False
        
        # Initialize Qdrant client (server mode)
        print(f"Connecting to Qdrant server at {qdrant_url}...")
        
        # Try to connect with retries
        max_retries = 5
        for attempt in range(max_retries):
            try:
                self.client = QdrantClient(url=qdrant_url)
                # Test connection
                self.client.get_collections()
                break
            except Exception as e:
                if attempt < max_retries - 1:
                    print(f"⚠️  Cannot connect to Qdrant server. Retrying in 2 seconds... ({attempt + 1}/{max_retries})")
                    time.sleep(2)
                else:
                    print("\n❌ Error: Cannot connect to Qdrant server.")
                    print(f"   URL: {qdrant_url}")
                    print("   Please ensure Qdrant is running: docker-compose up -d")
                    print(f"   Error: {e}\n")
                    sys.exit(1)
        
        # Try to use embeddings server first
        try:
            # Try Infinity health endpoint
            response = requests.get(f"{self.embeddings_url}/health", timeout=2)
            if response.status_code == 200:
                print(f"Using Infinity embedding server at {self.embeddings_url}")
                self.use_embeddings_server = True
        except:
            print(f"Embedding server not available, falling back to local model...")
        
        # Fallback to local model if server not available
        if not self.use_embeddings_server:
            from sentence_transformers import SentenceTransformer
            import torch
            device = 'cuda' if torch.cuda.is_available() else 'cpu'
            print(f"Loading embedding model ({self.EMBEDDING_MODEL}) locally on {device}...")
            self.encoder = SentenceTransformer(self.EMBEDDING_MODEL, device=device)
            if device == 'cuda':
                print(f"GPU: {torch.cuda.get_device_name(0)}, Memory: {torch.cuda.memory_allocated() / 1024**3:.2f} GB used")
        
        # Document tracking
        self.hash_file = self.kb_path / "qdrant_hashes.json"
        self.document_hashes = self._load_hashes()
        
        # Initialize collections
        self._init_collections()
        
        print("✅ Qdrant KB initialized successfully")
    
    def _init_collections(self):
        """Initialize Qdrant collections."""
        for name, description in self.COLLECTIONS.items():
            try:
                # Check if collection exists
                collections = self.client.get_collections().collections
                if not any(c.name == name for c in collections):
                    # Create collection with vector configuration
                    self.client.create_collection(
                        collection_name=name,
                        vectors_config=VectorParams(
                            size=self.EMBEDDING_DIM,
                            distance=Distance.COSINE
                        )
                    )
                    print(f"Created collection: {name}")
            except Exception as e:
                print(f"Collection {name} already exists or error: {e}")
    
    def _load_hashes(self) -> Dict[str, str]:
        """Load document hashes."""
        if self.hash_file.exists():
            with open(self.hash_file, 'r') as f:
                return json.load(f)
        return {}
    
    def _save_hashes(self):
        """Save document hashes."""
        with open(self.hash_file, 'w') as f:
            json.dump(self.document_hashes, f, indent=2)
    
    def _get_file_hash(self, file_path: Path) -> str:
        """Get hash of file contents."""
        with open(file_path, 'rb') as f:
            return hashlib.md5(f.read()).hexdigest()
    
    def _encode_text(self, texts: Union[str, List[str]]) -> np.ndarray:
        """Encode text using either embedding server or local model."""
        if isinstance(texts, str):
            texts = [texts]
        
        if self.use_embeddings_server:
            # Use embedding server (supports both TEI and Infinity)
            try:
                # Try Infinity API format first
                response = requests.post(
                    f"{self.embeddings_url}/embeddings",
                    json={
                        "model": "BAAI/bge-large-en-v1.5",
                        "input": texts
                    },
                    headers={"Content-Type": "application/json"},
                    timeout=30
                )
                if response.status_code == 200:
                    result = response.json()
                    # Infinity returns {"data": [{"embedding": [...]}]}
                    if "data" in result:
                        embeddings = np.array([item["embedding"] for item in result["data"]])
                    else:
                        # TEI format
                        embeddings = np.array(result)
                    return embeddings if len(texts) > 1 else embeddings[0]
                else:
                    raise Exception(f"Embedding server error: {response.status_code}")
            except Exception as e:
                print(f"Warning: Embedding server error ({e}), falling back to local model")
                # Fall back to local model
                if not hasattr(self, 'encoder'):
                    from sentence_transformers import SentenceTransformer
                    import torch
                    device = 'cuda' if torch.cuda.is_available() else 'cpu'
                    self.encoder = SentenceTransformer(self.EMBEDDING_MODEL, device=device)
                return self.encoder.encode(texts)
        else:
            # Use local model
            return self.encoder.encode(texts)
    
    def _chunk_text(self, text: str, chunk_size: int = 1000, overlap: int = 200) -> List[str]:
        """Chunk text with overlap."""
        chunks = []
        start = 0
        text_len = len(text)
        
        while start < text_len:
            end = min(start + chunk_size, text_len)
            chunk = text[start:end]
            
            # Try to break at sentence boundary
            if end < text_len:
                last_period = chunk.rfind('. ')
                if last_period > chunk_size * 0.5:
                    end = start + last_period + 1
                    chunk = text[start:end]
            
            chunks.append(chunk.strip())
            start = end - overlap if end < text_len else end
        
        return chunks
    
    def _determine_collection(self, file_path: Path) -> str:
        """Determine which collection a file belongs to."""
        path_str = str(file_path).lower()
        
        if 'architecture' in path_str or 'adr' in path_str:
            return 'architecture'
        elif 'zettelkasten' in path_str:
            return 'zettelkasten'
        elif any(ext in file_path.suffix for ext in ['.py', '.js', '.ts', '.go']):
            return 'code'
        elif 'harvested' in path_str:
            return 'harvested'
        else:
            return 'docs'
    
    def _extract_metadata(self, file_path: Path, chunk_index: int = 0) -> DocumentMetadata:
        """Extract metadata from file."""
        now = datetime.now().isoformat()
        file_hash = self._get_file_hash(file_path)
        
        # Detect source type
        path_lower = str(file_path).lower()
        if 'harvested' in path_lower:
            source_type = SourceType.HARVESTED.value
        elif 'generated' in path_lower:
            source_type = SourceType.GENERATED.value
        else:
            source_type = SourceType.INTERNAL.value
        
        # Detect agent creator
        agent_creator = None
        for agent in ['ana', 'archie', 'tina', 'paul', 'dave']:
            if agent in path_lower:
                agent_creator = agent
                break
        
        # Extract tags
        tags = []
        if 'test' in path_lower:
            tags.append('testing')
        if 'api' in path_lower:
            tags.append('api')
        if 'config' in path_lower:
            tags.append('config')
        
        return DocumentMetadata(
            file_path=str(file_path),
            chunk_index=chunk_index,
            source_type=source_type,
            document_category=self._determine_collection(file_path),
            agent_creator=agent_creator,
            indexed_at=now,
            file_hash=file_hash,
            tags=','.join(tags) if tags else None
        )
    
    def index_file(self, file_path: Path, force: bool = False) -> bool:
        """Index a single file."""
        file_path = Path(file_path).resolve()
        
        if not file_path.exists():
            return False
        
        # Check if already indexed
        file_hash = self._get_file_hash(file_path)
        if not force and str(file_path) in self.document_hashes:
            if self.document_hashes[str(file_path)] == file_hash:
                return False
        
        try:
            # Read file content
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            # Chunk the content
            chunks = self._chunk_text(content)
            
            # Determine collection
            collection_name = self._determine_collection(file_path)
            
            # Generate embeddings for all chunks at once (batch processing)
            if chunks:
                embeddings = self._encode_text(chunks)
                if len(chunks) == 1:
                    embeddings = [embeddings]  # Ensure it's a list of embeddings
            
            # Prepare points for Qdrant
            points = []
            for i, (chunk, embedding) in enumerate(zip(chunks, embeddings)):
                
                # Extract metadata
                metadata = self._extract_metadata(file_path, i)
                
                # Create point
                point_id = abs(hash(f"{file_path}_{i}")) % (10 ** 8)  # Generate stable ID
                point = PointStruct(
                    id=point_id,
                    vector=embedding.tolist() if hasattr(embedding, 'tolist') else list(embedding),
                    payload={
                        **metadata.to_dict(),
                        "content": chunk[:1000],  # Store first 1000 chars for preview
                        "full_content": chunk
                    }
                )
                points.append(point)
            
            # Upsert points to Qdrant
            if points:
                self.client.upsert(
                    collection_name=collection_name,
                    points=points
                )
            
            # Update hash
            self.document_hashes[str(file_path)] = file_hash
            self._save_hashes()
            
            print(f"✅ Indexed: {file_path.name} ({len(chunks)} chunks)")
            return True
            
        except Exception as e:
            print(f"❌ Error indexing {file_path}: {e}")
            return False
    
    def search(
        self, 
        query: str, 
        collection: Optional[str] = None,
        agent_context: Optional[str] = None,
        limit: int = 5,
        score_threshold: float = 0.0
    ) -> List[Dict[str, Any]]:
        """Search across collections."""
        # Generate query embedding
        query_embedding = self._encode_text(query)
        query_vector = query_embedding.tolist() if hasattr(query_embedding, 'tolist') else list(query_embedding)
        
        # Determine collections to search
        if collection:
            collections = [collection] if collection in self.COLLECTIONS else ['docs']
        else:
            # Search all collections
            collections = list(self.COLLECTIONS.keys())
        
        all_results = []
        
        for coll_name in collections:
            try:
                # Build filter
                filter_conditions = []
                if agent_context:
                    filter_conditions.append(
                        FieldCondition(
                            key="agent_creator",
                            match=MatchValue(value=agent_context)
                        )
                    )
                
                # Search in collection
                search_filter = Filter(must=filter_conditions) if filter_conditions else None
                
                results = self.client.query_points(
                    collection_name=coll_name,
                    query=query_vector,
                    query_filter=search_filter,
                    limit=limit,
                    score_threshold=score_threshold
                ).points
                
                # Format results
                for hit in results:
                    # Apply relevance boosting based on document path
                    score = hit.score
                    file_path = hit.payload.get('file_path', '').lower()
                    
                    # Boost high-value documentation
                    if 'architecture/adrs' in file_path:
                        score *= 1.3  # 30% boost for ADRs
                    elif 'architecture/patterns' in file_path:
                        score *= 1.25  # 25% boost for patterns
                    elif '/analysis/' in file_path:
                        score *= 1.2  # 20% boost for analysis
                    elif '/strategy/' in file_path:
                        score *= 1.2  # 20% boost for strategy
                    elif 'architecture/' in file_path:
                        score *= 1.15  # 15% boost for other architecture docs
                    elif '/research/' in file_path:
                        score *= 0.9  # 10% reduction for raw research
                    elif '/harvested/' in file_path:
                        score *= 0.85  # 15% reduction for external content (unless specific)
                    
                    # Apply agent-specific boosting
                    if agent_context:
                        if agent_context == 'architect' and 'architecture/' in file_path:
                            score *= 1.1  # Extra boost for architect viewing architecture docs
                        elif agent_context == 'analyst' and ('/analysis/' in file_path or '/strategy/' in file_path):
                            score *= 1.1  # Extra boost for analyst viewing analysis/strategy
                        elif agent_context == 'teacher' and 'zettelkasten/' in file_path:
                            score *= 1.15  # Extra boost for teacher viewing learning notes
                        elif agent_context == 'developer' and ('/code/' in file_path or '/api/' in file_path):
                            score *= 1.1  # Extra boost for developer viewing code
                    
                    # Cap score at 1.0
                    score = min(score, 1.0)
                    
                    result = {
                        'collection': coll_name,
                        'score': score,
                        'content': hit.payload.get('full_content', ''),
                        'metadata': {
                            k: v for k, v in hit.payload.items() 
                            if k not in ['content', 'full_content']
                        }
                    }
                    all_results.append(result)
                    
            except Exception as e:
                print(f"Error searching collection {coll_name}: {e}")
        
        # Sort by score and limit
        all_results.sort(key=lambda x: x['score'], reverse=True)
        return all_results[:limit]
    
    def index_directory(self, path: Path = None, recursive: bool = True, force: bool = False) -> int:
        """Index all files in a directory."""
        if path is None:
            path = self.project_root
        
        path = Path(path).resolve()
        
        if not path.exists():
            print(f"Error: Path {path} does not exist")
            return 0
        
        # Find all files
        files = []
        exclude_dirs = {'.git', 'node_modules', '__pycache__', '.venv', 'venv',
                       'build', 'dist', '.cache', 'coverage', '.vector_db'}
        
        if recursive:
            for item in path.rglob("*"):
                if item.is_file():
                    # Skip if in excluded directory
                    if not any(excluded in str(item) for excluded in exclude_dirs):
                        # Only index text files
                        if item.suffix in ['.md', '.py', '.js', '.ts', '.jsx', '.tsx', 
                                          '.go', '.rs', '.yaml', '.yml', '.json', '.txt']:
                            files.append(item)
        else:
            for item in path.iterdir():
                if item.is_file():
                    if item.suffix in ['.md', '.py', '.js', '.ts', '.jsx', '.tsx',
                                      '.go', '.rs', '.yaml', '.yml', '.json', '.txt']:
                        files.append(item)
        
        indexed_count = 0
        total_files = len(files)
        
        print(f"🔍 Indexing {total_files} files...")
        
        for i, file_path in enumerate(files, 1):
            if i % 10 == 0:
                print(f"  Progress: {i}/{total_files} files...")
            
            if self.index_file(file_path, force):
                indexed_count += 1
        
        print(f"✅ Indexed {indexed_count} files (examined {total_files} total)")
        return indexed_count
    
    def get_stats(self) -> Dict[str, Any]:
        """Get KB statistics."""
        stats = {
            'collections': {},
            'total_documents': len(self.document_hashes),
            'embedding_model': self.EMBEDDING_MODEL,
            'embedding_dim': self.EMBEDDING_DIM
        }
        
        # Get collection stats
        for name in self.COLLECTIONS:
            try:
                info = self.client.get_collection(name)
                stats['collections'][name] = {
                    'vectors_count': info.vectors_count,
                    'points_count': info.points_count
                }
            except:
                stats['collections'][name] = {'error': 'Could not retrieve'}
        
        return stats
    
    def clear(self, collection: Optional[str] = None) -> bool:
        """Clear collections."""
        try:
            if collection:
                self.client.delete_collection(collection)
                self.client.create_collection(
                    collection_name=collection,
                    vectors_config=VectorParams(
                        size=self.EMBEDDING_DIM,
                        distance=Distance.COSINE
                    )
                )
                print(f"Cleared collection: {collection}")
            else:
                # Clear all collections
                for name in self.COLLECTIONS:
                    self.client.delete_collection(name)
                    self.client.create_collection(
                        collection_name=name,
                        vectors_config=VectorParams(
                            size=self.EMBEDDING_DIM,
                            distance=Distance.COSINE
                        )
                    )
                self.document_hashes = {}
                self._save_hashes()
                print("Cleared all collections")
            return True
        except Exception as e:
            print(f"Error clearing: {e}")
            return False


def main():
    """CLI interface."""
    parser = argparse.ArgumentParser(description="Qdrant Knowledge Base")
    subparsers = parser.add_subparsers(dest="command", help="Commands")
    
    # Index command
    index_parser = subparsers.add_parser("index", help="Index files")
    index_parser.add_argument("--path", help="Path to index")
    index_parser.add_argument("--force", action="store_true", help="Force re-index")
    
    # Search command
    search_parser = subparsers.add_parser("search", help="Search KB")
    search_parser.add_argument("query", help="Search query")
    search_parser.add_argument("--collection", help="Specific collection")
    search_parser.add_argument("--agent", help="Agent context")
    search_parser.add_argument("--limit", type=int, default=5, help="Number of results")
    
    # Stats command
    subparsers.add_parser("stats", help="Show statistics")
    
    # Clear command
    clear_parser = subparsers.add_parser("clear", help="Clear KB")
    clear_parser.add_argument("--collection", help="Collection to clear")
    
    args = parser.parse_args()
    
    # Initialize KB - no need for different modes with server
    kb = QdrantKB()
    
    if args.command == "index":
        if args.path:
            # Resolve path relative to project root
            if Path(args.path).is_absolute():
                path = Path(args.path)
            else:
                path = kb.project_root / args.path
        else:
            path = None
        kb.index_directory(path, force=args.force)
    
    elif args.command == "search":
        results = kb.search(
            args.query,
            collection=args.collection,
            agent_context=args.agent,
            limit=args.limit
        )
        
        print(f"\n📊 Found {len(results)} results:\n")
        for i, result in enumerate(results, 1):
            print(f"{i}. Collection: {result['collection']}")
            print(f"   Score: {result['score']:.4f}")
            
            metadata = result.get('metadata', {})
            if metadata:
                # Show file path
                file_path = metadata.get('file_path', 'unknown')
                chunk_index = metadata.get('chunk_index', 0)
                print(f"   File: {file_path}:{chunk_index}")
                print(f"   Source: {metadata.get('source_type', 'unknown')}")
                if metadata.get('agent_creator'):
                    print(f"   Creator: {metadata['agent_creator']}")
            
            # Show preview
            content = result['content'][:200] + "..." if len(result['content']) > 200 else result['content']
            print(f"   Preview: {content}")
            print()
    
    elif args.command == "stats":
        stats = kb.get_stats()
        print("\n📈 Qdrant KB Statistics:")
        print(f"Model: {stats['embedding_model']} ({stats['embedding_dim']} dims)")
        print(f"Total indexed files: {stats['total_documents']}")
        print("\nCollections:")
        for name, info in stats['collections'].items():
            if 'error' not in info:
                print(f"  {name}: {info['points_count']} chunks")
    
    elif args.command == "clear":
        kb.clear(args.collection)
    
    else:
        parser.print_help()


if __name__ == "__main__":
    main()