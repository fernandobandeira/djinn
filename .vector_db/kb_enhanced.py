#!/usr/bin/env -S uv run python
"""
Enhanced Knowledge Base Manager for AI Agents
Optimized for local AI-assisted development with improved search and indexing
"""

import os
import sys
import hashlib
import json
import re
import argparse
import time
import threading
from pathlib import Path
from typing import List, Dict, Any, Optional, Tuple
from datetime import datetime, timedelta
from dataclasses import dataclass, asdict
from enum import Enum
from collections import defaultdict
import numpy as np

# Add parent directory to path for imports
sys.path.insert(0, str(Path(__file__).parent.parent))

try:
    import chromadb
    from chromadb.config import Settings
    from chromadb.utils import embedding_functions
    from sentence_transformers import SentenceTransformer, CrossEncoder
    from rank_bm25 import BM25Okapi
    import nltk
    from nltk.corpus import stopwords
    from nltk.tokenize import word_tokenize
    # Download required NLTK data
    try:
        nltk.data.find('tokenizers/punkt')
    except LookupError:
        nltk.download('punkt')
    try:
        nltk.data.find('corpora/stopwords')
    except LookupError:
        nltk.download('stopwords')
except ImportError as e:
    print(f"Installing required packages: {e}")
    os.system("pip install chromadb sentence-transformers rank-bm25 nltk watchdog")
    import chromadb
    from chromadb.config import Settings
    from chromadb.utils import embedding_functions
    from sentence_transformers import SentenceTransformer, CrossEncoder
    from rank_bm25 import BM25Okapi
    import nltk
    nltk.download('punkt')
    nltk.download('stopwords')
    from nltk.corpus import stopwords
    from nltk.tokenize import word_tokenize


class IndexMode(Enum):
    """Different indexing modes for various use cases."""
    ALL = "all"
    ARCHITECTURE = "arch"
    CODE = "code"
    DOCS = "docs"
    CONFIG = "config"
    ZETTELKASTEN = "zettel"


class SourceType(Enum):
    """Document source types for categorization."""
    INTERNAL = "internal"       # Created by agents
    HARVESTED = "harvested"     # Fetched from web
    GENERATED = "generated"     # AI-generated content
    USER = "user"              # User-created content


@dataclass
class DocumentMetadata:
    """Enhanced metadata schema for documents."""
    source_type: SourceType
    document_id: str
    chunk_index: int
    chunk_type: str  # content|header|table|code
    created_at: str
    updated_at: str
    language: str = "en"
    confidence_score: float = 0.95
    parent_document: Optional[str] = None
    section_path: Optional[str] = None
    source_url: Optional[str] = None  # For harvested docs
    agent_creator: Optional[str] = None  # archie|ana|tina|etc
    document_category: Optional[str] = None  # architecture|research|zettelkasten
    harvest_profile: Optional[str] = None  # deep_research|api_docs|troubleshooting
    last_validated: Optional[str] = None
    file_hash: Optional[str] = None
    tags: List[str] = None
    
    def to_dict(self):
        """Convert to dictionary for ChromaDB metadata."""
        data = asdict(self)
        # ChromaDB doesn't handle enums well
        data['source_type'] = self.source_type.value
        # Filter None values
        return {k: v for k, v in data.items() if v is not None}


@dataclass
class DocumentType:
    """Configuration for different document types."""
    name: str
    extensions: List[str]
    paths: List[str]
    chunk_size: int = 1000
    chunk_overlap: int = 200
    collection: str = "general"
    mode: IndexMode = IndexMode.ALL


class QueryPreprocessor:
    """Handles LLM query verbosity and optimization."""
    
    def __init__(self):
        self.stop_words = set(stopwords.words('english'))
        # Add common LLM filler words
        self.llm_filler_words = {
            'please', 'could', 'would', 'might', 'perhaps', 'maybe',
            'wondering', 'looking', 'trying', 'need', 'want', 'help',
            'find', 'search', 'locate', 'get', 'retrieve', 'show',
            'tell', 'explain', 'describe', 'about', 'regarding',
            'related', 'pertaining', 'concerning', 'also', 'additionally',
            'furthermore', 'moreover', 'basically', 'essentially'
        }
        self.stop_words.update(self.llm_filler_words)
    
    def preprocess(self, query: str, agent_context: str = None) -> Dict[str, Any]:
        """Preprocess LLM query for better retrieval."""
        # Original query for semantic search
        original = query.lower().strip()
        
        # Extract key terms for keyword search
        tokens = word_tokenize(original)
        key_terms = [t for t in tokens if t not in self.stop_words and len(t) > 2]
        
        # Generate query variants
        variants = {
            'original': original,
            'key_terms': ' '.join(key_terms),
            'compressed': self._compress_query(original),
            'expanded': self._expand_query(key_terms, agent_context)
        }
        
        # Agent-specific boosting
        boost_terms = []
        if agent_context:
            if agent_context == 'architect':
                boost_terms = ['architecture', 'design', 'pattern', 'adr', 'system']
            elif agent_context == 'teacher':
                boost_terms = ['learning', 'concept', 'explanation', 'zettelkasten']
            elif agent_context == 'developer':
                boost_terms = ['implementation', 'code', 'function', 'test', 'api']
            elif agent_context == 'product':
                boost_terms = ['requirement', 'user', 'story', 'feature', 'business']
        
        return {
            'variants': variants,
            'key_terms': key_terms,
            'boost_terms': boost_terms,
            'agent_context': agent_context
        }
    
    def _compress_query(self, query: str) -> str:
        """Remove redundant words while preserving meaning."""
        # Remove common question patterns
        patterns = [
            r'^(can you |could you |would you |please )?',
            r'^(i need |i want |i\'m looking for |i\'m trying to )',
            r'^(help me |show me |tell me |explain )',
            r'( please| thanks| thank you)$',
        ]
        
        compressed = query
        for pattern in patterns:
            compressed = re.sub(pattern, '', compressed, flags=re.IGNORECASE)
        
        return compressed.strip()
    
    def _expand_query(self, key_terms: List[str], agent_context: str) -> str:
        """Expand query with synonyms and related terms."""
        expansions = {
            'auth': 'authentication authorization security',
            'db': 'database storage persistence',
            'api': 'endpoint service interface',
            'test': 'testing unittest integration',
            'doc': 'documentation docs document',
            'config': 'configuration settings setup',
            'error': 'exception bug issue problem',
            'perf': 'performance optimization speed'
        }
        
        expanded_terms = []
        for term in key_terms:
            expanded_terms.append(term)
            # Check for common abbreviations
            for abbr, expansion in expansions.items():
                if abbr in term.lower():
                    expanded_terms.extend(expansion.split())
        
        return ' '.join(expanded_terms)


class HybridSearchEngine:
    """Implements hybrid BM25 + Vector search with reranking."""
    
    def __init__(self, embedding_model_name: str = "BAAI/bge-large-en-v1.5"):
        # Better embedding model
        self.embedding_model = SentenceTransformer(embedding_model_name)
        
        # Reranker for better precision
        self.reranker = CrossEncoder('BAAI/bge-reranker-base', max_length=512)
        
        # BM25 components
        self.bm25_index = {}  # collection -> BM25Okapi
        self.bm25_docs = {}   # collection -> List[str]
        self.bm25_metadata = {}  # collection -> List[metadata]
    
    def build_bm25_index(self, documents: List[str], metadata: List[Dict], collection: str):
        """Build BM25 index for keyword search."""
        # Tokenize documents
        tokenized_docs = []
        for doc in documents:
            tokens = word_tokenize(doc.lower())
            tokenized_docs.append(tokens)
        
        # Create BM25 index
        self.bm25_index[collection] = BM25Okapi(tokenized_docs)
        self.bm25_docs[collection] = documents
        self.bm25_metadata[collection] = metadata
    
    def hybrid_search(self, 
                     query: str,
                     vector_results: List[Dict],
                     collection: str,
                     alpha: float = 0.5) -> List[Dict]:
        """Combine BM25 and vector search results."""
        
        if collection not in self.bm25_index:
            # No BM25 index, return vector results
            return vector_results
        
        # BM25 search
        query_tokens = word_tokenize(query.lower())
        bm25_scores = self.bm25_index[collection].get_scores(query_tokens)
        
        # Get top BM25 results
        top_indices = np.argsort(bm25_scores)[-50:][::-1]
        bm25_results = []
        for idx in top_indices:
            if bm25_scores[idx] > 0:
                bm25_results.append({
                    'content': self.bm25_docs[collection][idx],
                    'metadata': self.bm25_metadata[collection][idx],
                    'score': float(bm25_scores[idx]),
                    'method': 'bm25'
                })
        
        # Reciprocal Rank Fusion (RRF)
        return self._reciprocal_rank_fusion(vector_results, bm25_results, alpha)
    
    def _reciprocal_rank_fusion(self, 
                               vector_results: List[Dict],
                               bm25_results: List[Dict],
                               alpha: float = 0.5,
                               k: int = 60) -> List[Dict]:
        """Combine results using Reciprocal Rank Fusion."""
        scores = defaultdict(float)
        result_map = {}
        
        # Score vector results
        for rank, result in enumerate(vector_results):
            doc_id = result.get('metadata', {}).get('document_id', str(rank))
            scores[doc_id] += alpha * (1 / (k + rank + 1))
            result_map[doc_id] = result
        
        # Score BM25 results  
        for rank, result in enumerate(bm25_results):
            doc_id = result.get('metadata', {}).get('document_id', str(rank))
            scores[doc_id] += (1 - alpha) * (1 / (k + rank + 1))
            if doc_id not in result_map:
                result_map[doc_id] = result
        
        # Sort by combined score
        sorted_ids = sorted(scores.keys(), key=lambda x: scores[x], reverse=True)
        
        # Return combined results
        combined = []
        for doc_id in sorted_ids:
            result = result_map[doc_id]
            result['hybrid_score'] = scores[doc_id]
            combined.append(result)
        
        return combined
    
    def rerank(self, query: str, results: List[Dict], top_k: int = 10) -> List[Dict]:
        """Rerank results using cross-encoder."""
        if not results:
            return results
        
        # Prepare query-document pairs
        pairs = []
        for result in results[:30]:  # Rerank top 30
            content = result.get('content', '')
            pairs.append([query, content])
        
        # Get reranking scores
        if pairs:
            scores = self.reranker.predict(pairs)
            
            # Add scores to results
            for i, score in enumerate(scores):
                results[i]['rerank_score'] = float(score)
            
            # Sort by rerank score
            results = sorted(results[:len(scores)], 
                           key=lambda x: x.get('rerank_score', 0), 
                           reverse=True)
        
        return results[:top_k]


class IndexStatusTracker:
    """Tracks indexing status and health."""
    
    def __init__(self, status_file: Path):
        self.status_file = status_file
        self.status = self._load_status()
    
    def _load_status(self) -> Dict:
        """Load status from file."""
        if self.status_file.exists():
            with open(self.status_file, 'r') as f:
                return json.load(f)
        return self._default_status()
    
    def _default_status(self) -> Dict:
        """Default status structure."""
        return {
            'last_full_index': None,
            'last_incremental_index': None,
            'pending_files': [],
            'failed_files': [],
            'index_health': 'unknown',
            'auto_index_enabled': True,
            'next_scheduled_index': None,
            'statistics': {
                'total_documents': 0,
                'total_chunks': 0,
                'avg_query_time_ms': 0,
                'last_query_time': None
            }
        }
    
    def save(self):
        """Save status to file."""
        with open(self.status_file, 'w') as f:
            json.dump(self.status, f, indent=2)
    
    def mark_indexed(self, file_path: str):
        """Mark file as indexed."""
        if file_path in self.status['pending_files']:
            self.status['pending_files'].remove(file_path)
        self.status['last_incremental_index'] = datetime.now().isoformat()
        self.update_health()
        self.save()
    
    def mark_pending(self, file_path: str):
        """Mark file as pending indexing."""
        if file_path not in self.status['pending_files']:
            self.status['pending_files'].append(file_path)
        self.update_health()
        self.save()
    
    def mark_failed(self, file_path: str, error: str):
        """Mark file as failed to index."""
        self.status['failed_files'].append({
            'path': file_path,
            'error': str(error),
            'timestamp': datetime.now().isoformat()
        })
        self.update_health()
        self.save()
    
    def update_health(self):
        """Update index health status."""
        pending_count = len(self.status['pending_files'])
        failed_count = len(self.status['failed_files'])
        
        if failed_count > 10:
            self.status['index_health'] = 'critical'
        elif failed_count > 5 or pending_count > 50:
            self.status['index_health'] = 'degraded'
        elif pending_count > 20:
            self.status['index_health'] = 'stale'
        else:
            self.status['index_health'] = 'healthy'
    
    def get_health_report(self) -> Dict:
        """Get comprehensive health report."""
        return {
            'health': self.status['index_health'],
            'pending_count': len(self.status['pending_files']),
            'failed_count': len(self.status['failed_files']),
            'last_index': self.status['last_incremental_index'],
            'auto_index': self.status['auto_index_enabled'],
            'statistics': self.status['statistics']
        }


class EnhancedKnowledgeBase:
    """Enhanced KB with all improvements for AI agent use."""
    
    # Document type configurations
    DOCUMENT_TYPES = [
        # Architecture documentation
        DocumentType("adr", [".md"], ["docs/architecture/adrs"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        DocumentType("design", [".md"], ["docs/architecture/designs"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        DocumentType("pattern", [".md"], ["docs/architecture/patterns"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        
        # Zettelkasten knowledge notes
        DocumentType("zettelkasten", [".md"], 
                    ["zettelkasten", "zettelkasten/permanent", "zettelkasten/literature", 
                     "zettelkasten/hubs", "zettelkasten/maps"],
                    chunk_size=800, collection="zettelkasten", mode=IndexMode.ZETTELKASTEN),
        
        # General documentation
        DocumentType("documentation", [".md", ".rst", ".txt"], 
                    ["docs", "README.md", "CLAUDE.md"], 
                    collection="documentation", mode=IndexMode.DOCS),
        
        # Harvested content
        DocumentType("harvested", [".md"], 
                    ["harvested"],
                    collection="harvested", mode=IndexMode.DOCS),
        
        # Source code
        DocumentType("python", [".py"], ["*.py", "src", "lib"], 
                    chunk_size=500, collection="code", mode=IndexMode.CODE),
        DocumentType("javascript", [".js", ".jsx", ".ts", ".tsx"], 
                    ["src", "lib", "app"], 
                    chunk_size=500, collection="code", mode=IndexMode.CODE),
        
        # Configuration
        DocumentType("config", [".json", ".yaml", ".yml", ".toml"], 
                    ["config", ".claude"], 
                    collection="config", mode=IndexMode.CONFIG),
        
        # API specifications
        DocumentType("api", [".yaml", ".yml", ".json"], 
                    ["api", "openapi", "swagger"], 
                    collection="api", mode=IndexMode.ALL),
        
        # Tests
        DocumentType("test", [".py", ".js", ".ts"], 
                    ["test", "tests", "__tests__"], 
                    collection="tests", mode=IndexMode.CODE),
    ]
    
    def __init__(self, project_root: str = None):
        """Initialize enhanced knowledge base."""
        if project_root is None:
            project_root = Path(__file__).parent.parent
        
        self.project_root = Path(project_root)
        self.kb_path = self.project_root / ".vector_db"
        self.kb_path.mkdir(parents=True, exist_ok=True)
        
        # Initialize ChromaDB
        self.client = chromadb.PersistentClient(
            path=str(self.kb_path / "chromadb"),
            settings=Settings(
                anonymized_telemetry=False,
                allow_reset=True
            )
        )
        
        # Better embedding model
        print("Loading enhanced embedding model (BAAI/bge-large-en-v1.5)...")
        self.embedding_function = embedding_functions.SentenceTransformerEmbeddingFunction(
            model_name="BAAI/bge-large-en-v1.5"
        )
        
        # Initialize components
        self.collections = self._init_collections()
        self.query_preprocessor = QueryPreprocessor()
        self.hybrid_search_engine = HybridSearchEngine()
        self.status_tracker = IndexStatusTracker(self.kb_path / "index_status.json")
        
        # Document tracking
        self.hash_file = self.kb_path / "document_hashes.json"
        self.document_hashes = self._load_hashes()
        
        print("✅ Enhanced KB initialized with improved search capabilities")
    
    def _init_collections(self) -> Dict[str, Any]:
        """Initialize all collections."""
        collections = {}
        
        # Get unique collection names including new 'harvested' collection
        collection_names = set(['harvested'])  # Add harvested collection
        for dt in self.DOCUMENT_TYPES:
            collection_names.add(dt.collection)
        collection_names.add("general")
        
        for name in collection_names:
            try:
                collections[name] = self.client.get_or_create_collection(
                    name=name,
                    embedding_function=self.embedding_function
                )
            except Exception as e:
                print(f"Warning: Could not create collection {name}: {e}")
        
        return collections
    
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
    
    def _detect_source_type(self, file_path: Path) -> SourceType:
        """Detect document source type."""
        path_str = str(file_path)
        
        if 'harvested' in path_str:
            return SourceType.HARVESTED
        elif '.claude/agents' in path_str or 'generated' in path_str:
            return SourceType.GENERATED
        elif any(agent in path_str for agent in ['archie', 'ana', 'tina', 'paul', 'dave']):
            return SourceType.INTERNAL
        else:
            return SourceType.USER
    
    def _extract_metadata(self, file_path: Path, chunk_index: int = 0) -> DocumentMetadata:
        """Extract enhanced metadata from file."""
        now = datetime.now().isoformat()
        file_hash = self._get_file_hash(file_path)
        
        # Detect source type
        source_type = self._detect_source_type(file_path)
        
        # Detect agent creator from path or content
        agent_creator = None
        path_lower = str(file_path).lower()
        if 'archie' in path_lower or 'architect' in path_lower:
            agent_creator = 'archie'
        elif 'ana' in path_lower or 'analyst' in path_lower:
            agent_creator = 'ana'
        elif 'tina' in path_lower or 'teacher' in path_lower:
            agent_creator = 'tina'
        elif 'paul' in path_lower or 'product' in path_lower:
            agent_creator = 'paul'
        elif 'dave' in path_lower or 'developer' in path_lower:
            agent_creator = 'dave'
        
        # Detect document category
        doc_category = None
        if 'architecture' in path_lower or 'adr' in path_lower:
            doc_category = 'architecture'
        elif 'zettelkasten' in path_lower:
            doc_category = 'zettelkasten'
        elif 'research' in path_lower:
            doc_category = 'research'
        elif 'brainstorm' in path_lower:
            doc_category = 'brainstorming'
        
        # Extract tags from content or filename
        tags = []
        if 'test' in path_lower:
            tags.append('testing')
        if 'api' in path_lower:
            tags.append('api')
        if 'config' in path_lower:
            tags.append('configuration')
        
        return DocumentMetadata(
            source_type=source_type,
            document_id=str(file_path),
            chunk_index=chunk_index,
            chunk_type='content',
            created_at=now,
            updated_at=now,
            file_hash=file_hash,
            agent_creator=agent_creator,
            document_category=doc_category,
            tags=tags if tags else None,
            confidence_score=0.95
        )
    
    def index_file(self, file_path: Path, force: bool = False) -> bool:
        """Index a single file with enhanced metadata."""
        file_path = Path(file_path).resolve()
        
        # Check if file exists
        if not file_path.exists():
            return False
        
        # Check if already indexed (unless forced)
        file_hash = self._get_file_hash(file_path)
        if not force and str(file_path) in self.document_hashes:
            if self.document_hashes[str(file_path)] == file_hash:
                return False
        
        try:
            # Read file content
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            # Determine collection
            collection_name = self._determine_collection(file_path)
            if collection_name not in self.collections:
                collection_name = "general"
            
            collection = self.collections[collection_name]
            
            # Chunk content
            chunks = self._chunk_content(content)
            
            # Prepare for indexing
            texts = []
            metadatas = []
            ids = []
            
            for i, chunk in enumerate(chunks):
                metadata = self._extract_metadata(file_path, i)
                
                texts.append(chunk)
                metadatas.append(metadata.to_dict())
                ids.append(f"{file_path}_{i}")
            
            # Add to collection
            if texts:
                collection.add(
                    documents=texts,
                    metadatas=metadatas,
                    ids=ids
                )
                
                # Update BM25 index for hybrid search
                self.hybrid_search_engine.build_bm25_index(texts, metadatas, collection_name)
            
            # Update hash
            self.document_hashes[str(file_path)] = file_hash
            self._save_hashes()
            
            # Update status
            self.status_tracker.mark_indexed(str(file_path))
            
            print(f"✅ Indexed: {file_path.name} ({len(chunks)} chunks)")
            return True
            
        except Exception as e:
            print(f"❌ Error indexing {file_path}: {e}")
            self.status_tracker.mark_failed(str(file_path), str(e))
            return False
    
    def _determine_collection(self, file_path: Path) -> str:
        """Determine which collection a file belongs to."""
        path_str = str(file_path).lower()
        
        # Check for specific patterns
        if 'harvested' in path_str:
            return 'harvested'
        elif 'zettelkasten' in path_str:
            return 'zettelkasten'
        elif 'architecture' in path_str or 'adr' in path_str:
            return 'architecture'
        elif file_path.suffix in ['.py', '.js', '.ts', '.jsx', '.tsx']:
            return 'code'
        elif file_path.suffix in ['.yaml', '.yml', '.json', '.toml']:
            return 'config'
        elif 'test' in path_str:
            return 'tests'
        elif 'api' in path_str or 'swagger' in path_str or 'openapi' in path_str:
            return 'api'
        else:
            return 'documentation'
    
    def _chunk_content(self, content: str, chunk_size: int = 1000, overlap: int = 200) -> List[str]:
        """Chunk content with overlap."""
        chunks = []
        start = 0
        
        while start < len(content):
            end = start + chunk_size
            chunk = content[start:end]
            
            # Try to break at sentence boundary
            if end < len(content):
                last_period = chunk.rfind('.')
                last_newline = chunk.rfind('\n')
                break_point = max(last_period, last_newline)
                
                if break_point > chunk_size * 0.8:
                    chunk = content[start:start + break_point + 1]
                    end = start + break_point + 1
            
            chunks.append(chunk)
            start = end - overlap
        
        return chunks
    
    def search(self, 
              query: str, 
              collection: str = None, 
              limit: int = 10,
              agent_context: str = None,
              use_hybrid: bool = True,
              use_reranking: bool = True) -> List[Dict]:
        """Enhanced search with preprocessing, hybrid search, and reranking."""
        
        start_time = time.time()
        
        # Preprocess query
        processed = self.query_preprocessor.preprocess(query, agent_context)
        search_query = processed['variants']['original']
        
        # Determine collections to search
        if collection:
            collections = [self.collections.get(collection, self.collections["general"])]
        else:
            # Agent-specific collection prioritization
            if agent_context == 'architect':
                collections = [self.collections.get(c) for c in ['architecture', 'documentation', 'code']]
            elif agent_context == 'teacher':
                collections = [self.collections.get(c) for c in ['zettelkasten', 'documentation']]
            elif agent_context == 'developer':
                collections = [self.collections.get(c) for c in ['code', 'tests', 'api']]
            else:
                collections = list(self.collections.values())
        
        all_results = []
        
        for coll in collections:
            if coll is None:
                continue
                
            try:
                # Vector search
                results = coll.query(
                    query_texts=[search_query],
                    n_results=min(limit * 2, coll.count()) if coll.count() > 0 else 1
                )
                
                if results["documents"][0]:
                    for i, doc in enumerate(results["documents"][0]):
                        all_results.append({
                            "content": doc,
                            "metadata": results["metadatas"][0][i] if results["metadatas"] else {},
                            "distance": results["distances"][0][i] if "distances" in results else 0,
                            "collection": coll.name,
                            "method": "vector"
                        })
            except Exception as e:
                print(f"Warning: Search error in collection {coll.name}: {e}")
                continue
        
        # Apply hybrid search if enabled
        if use_hybrid and all_results:
            all_results = self.hybrid_search_engine.hybrid_search(
                search_query, 
                all_results, 
                collection or "general",
                alpha=0.4  # Favor BM25 for technical queries
            )
        
        # Apply reranking if enabled
        if use_reranking and all_results:
            all_results = self.hybrid_search_engine.rerank(
                search_query,
                all_results,
                top_k=limit
            )
        else:
            # Sort by distance/score
            all_results.sort(key=lambda x: x.get('hybrid_score', x.get('distance', 0)), reverse=False)
            all_results = all_results[:limit]
        
        # Track query time
        query_time_ms = (time.time() - start_time) * 1000
        self.status_tracker.status['statistics']['last_query_time'] = datetime.now().isoformat()
        self.status_tracker.status['statistics']['avg_query_time_ms'] = query_time_ms
        self.status_tracker.save()
        
        return all_results
    
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
                        files.append(item)
        else:
            files = [f for f in path.glob("*") if f.is_file()]
        
        indexed_count = 0
        total_files = len(files)
        
        print(f"🔍 Indexing {total_files} files...")
        
        for i, file_path in enumerate(files, 1):
            if i % 10 == 0:
                print(f"  Progress: {i}/{total_files} files...")
            
            if self.index_file(file_path, force):
                indexed_count += 1
        
        # Update statistics
        self.status_tracker.status['last_full_index'] = datetime.now().isoformat()
        self.status_tracker.status['statistics']['total_documents'] = len(self.document_hashes)
        self.status_tracker.save()
        
        print(f"✅ Indexed {indexed_count} files (examined {total_files} total)")
        return indexed_count
    
    def get_status(self) -> Dict:
        """Get comprehensive KB status."""
        return {
            'health': self.status_tracker.get_health_report(),
            'collections': {
                name: {
                    'count': coll.count() if coll else 0,
                    'name': name
                }
                for name, coll in self.collections.items()
            },
            'total_documents': len(self.document_hashes),
            'last_index': self.status_tracker.status['last_full_index'],
            'pending_files': len(self.status_tracker.status['pending_files']),
            'search_performance': {
                'avg_query_time_ms': self.status_tracker.status['statistics']['avg_query_time_ms'],
                'last_query': self.status_tracker.status['statistics']['last_query_time']
            }
        }
    
    def clear(self, collection: str = None) -> bool:
        """Clear the knowledge base."""
        try:
            if collection:
                if collection in self.collections:
                    self.client.delete_collection(collection)
                    self.collections[collection] = self.client.create_collection(
                        name=collection,
                        embedding_function=self.embedding_function
                    )
                    print(f"Cleared collection: {collection}")
            else:
                for name in list(self.collections.keys()):
                    self.client.delete_collection(name)
                self.collections = self._init_collections()
                self.document_hashes = {}
                self._save_hashes()
                print("Cleared entire knowledge base")
            
            return True
        except Exception as e:
            print(f"Error clearing: {e}")
            return False


def main():
    """Enhanced CLI interface."""
    parser = argparse.ArgumentParser(description="Enhanced Knowledge Base Manager")
    subparsers = parser.add_subparsers(dest="command", help="Commands")
    
    # Index command
    index_parser = subparsers.add_parser("index", help="Index documents")
    index_parser.add_argument("--path", help="Path to index")
    index_parser.add_argument("--force", action="store_true", help="Force re-index")
    index_parser.add_argument("--mode", choices=["all", "arch", "code", "docs"], 
                            default="all", help="Indexing mode")
    
    # Search command
    search_parser = subparsers.add_parser("search", help="Search knowledge base")
    search_parser.add_argument("query", help="Search query")
    search_parser.add_argument("--collection", help="Collection to search")
    search_parser.add_argument("--limit", type=int, default=10, help="Result limit")
    search_parser.add_argument("--agent", help="Agent context (architect|teacher|developer|product)")
    search_parser.add_argument("--no-hybrid", action="store_true", help="Disable hybrid search")
    search_parser.add_argument("--no-rerank", action="store_true", help="Disable reranking")
    
    # Status command
    subparsers.add_parser("status", help="Show KB status")
    
    # Stats command
    subparsers.add_parser("stats", help="Show statistics")
    
    # Clear command
    clear_parser = subparsers.add_parser("clear", help="Clear knowledge base")
    clear_parser.add_argument("--collection", help="Collection to clear")
    
    # Build graph command
    subparsers.add_parser("build-graph", help="Build GraphRAG knowledge graph")
    
    args = parser.parse_args()
    
    # Initialize KB
    kb = EnhancedKnowledgeBase()
    
    if args.command == "index":
        path = Path(args.path) if args.path else None
        kb.index_directory(path, force=args.force)
    
    elif args.command == "search":
        results = kb.search(
            args.query,
            collection=args.collection,
            limit=args.limit,
            agent_context=args.agent,
            use_hybrid=not args.no_hybrid if hasattr(args, 'no_hybrid') else True,
            use_reranking=not args.no_rerank if hasattr(args, 'no_rerank') else True
        )
        
        print(f"\n📊 Found {len(results)} results:\n")
        for i, result in enumerate(results, 1):
            print(f"{i}. Collection: {result['collection']}")
            print(f"   Score: {result.get('rerank_score', result.get('hybrid_score', result.get('distance', 0))):.4f}")
            print(f"   Method: {result.get('method', 'unknown')}")
            
            metadata = result.get('metadata', {})
            if metadata:
                print(f"   Source: {metadata.get('source_type', 'unknown')}")
                print(f"   Category: {metadata.get('document_category', 'unknown')}")
                if metadata.get('agent_creator'):
                    print(f"   Creator: {metadata['agent_creator']}")
            
            # Show preview
            content = result['content'][:200] + "..." if len(result['content']) > 200 else result['content']
            print(f"   Preview: {content}")
            print()
    
    elif args.command == "status":
        status = kb.get_status()
        print("\n📈 Knowledge Base Status:")
        print(f"Health: {status['health']['health']}")
        print(f"Total Documents: {status['total_documents']}")
        print(f"Pending Files: {status['pending_files']}")
        print(f"Last Index: {status['last_index']}")
        print(f"\nCollections:")
        for name, info in status['collections'].items():
            print(f"  {name}: {info['count']} chunks")
        print(f"\nSearch Performance:")
        print(f"  Avg Query Time: {status['search_performance']['avg_query_time_ms']:.2f}ms")
        print(f"  Last Query: {status['search_performance']['last_query']}")
    
    elif args.command == "stats":
        status = kb.get_status()
        print(json.dumps(status, indent=2))
    
    elif args.command == "clear":
        kb.clear(collection=args.collection)
    
    elif args.command == "build-graph":
        print("Building enhanced GraphRAG knowledge graph...")
        # This would trigger the GraphRAG builder
        from enhanced_graphrag import EnhancedGraphRAG
        graph = EnhancedGraphRAG()
        import asyncio
        asyncio.run(graph.build_knowledge_graph())
    
    else:
        parser.print_help()


if __name__ == "__main__":
    main()