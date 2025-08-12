#!/usr/bin/env python3
"""
Unified Knowledge Base Manager
A single tool for all vector database operations in the project.
"""

import os
import sys
import hashlib
import json
import re
import argparse
from pathlib import Path
from typing import List, Dict, Any, Optional
from datetime import datetime
from dataclasses import dataclass
from enum import Enum

# Add parent directory to path for imports
sys.path.insert(0, str(Path(__file__).parent.parent))

try:
    import chromadb
    from chromadb.config import Settings
    from chromadb.utils import embedding_functions
except ImportError:
    print("Installing required packages...")
    os.system("pip install chromadb sentence-transformers")
    import chromadb
    from chromadb.config import Settings
    from chromadb.utils import embedding_functions


class IndexMode(Enum):
    """Different indexing modes for various use cases."""
    ALL = "all"              # Index everything
    ARCHITECTURE = "arch"    # Only architecture docs
    CODE = "code"           # Only source code
    DOCS = "docs"           # Only documentation
    CONFIG = "config"       # Only configuration files


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


class KnowledgeBase:
    """Unified knowledge base manager for the entire project."""
    
    # Define document types and their configurations
    DOCUMENT_TYPES = [
        # Architecture documentation
        DocumentType("adr", [".md"], ["docs/architecture/adrs"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        DocumentType("design", [".md"], ["docs/architecture/designs"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        DocumentType("pattern", [".md"], ["docs/architecture/patterns"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        DocumentType("standard", [".md"], ["docs/architecture/standards"], 
                    collection="architecture", mode=IndexMode.ARCHITECTURE),
        
        # General documentation
        DocumentType("documentation", [".md", ".rst", ".txt"], 
                    ["docs", "README.md", "CLAUDE.md", "*.md"], 
                    collection="documentation", mode=IndexMode.DOCS),
        
        # Source code
        DocumentType("python", [".py"], ["*.py", "src", "lib", "scripts"], 
                    chunk_size=500, collection="code", mode=IndexMode.CODE),
        DocumentType("javascript", [".js", ".jsx", ".ts", ".tsx"], 
                    ["src", "lib", "app", "components"], 
                    chunk_size=500, collection="code", mode=IndexMode.CODE),
        DocumentType("web", [".html", ".css", ".scss"], 
                    ["src", "public", "static"], 
                    chunk_size=500, collection="code", mode=IndexMode.CODE),
        
        # Configuration
        DocumentType("config", [".json", ".yaml", ".yml", ".toml", ".ini", ".env"], 
                    ["config", ".claude", ".", "settings"], 
                    collection="config", mode=IndexMode.CONFIG),
        
        # API specifications
        DocumentType("api", [".yaml", ".yml", ".json"], 
                    ["api", "openapi", "swagger", "spec"], 
                    collection="api", mode=IndexMode.ALL),
        
        # Tests
        DocumentType("test", [".py", ".js", ".ts"], 
                    ["test", "tests", "__tests__", "spec", "*test*", "*spec*"], 
                    collection="tests", mode=IndexMode.CODE),
        
        # Scripts and tools
        DocumentType("script", [".sh", ".bash", ".ps1", ".bat", ".cmd"], 
                    ["scripts", "bin", "tools"], 
                    collection="scripts", mode=IndexMode.ALL),
    ]
    
    def __init__(self, project_root: str = None):
        """Initialize the knowledge base."""
        if project_root is None:
            # Auto-detect project root (parent of .vector_db)
            project_root = Path(__file__).parent.parent
        
        self.project_root = Path(project_root)
        self.kb_path = self.project_root / ".vector_db"
        self.kb_path.mkdir(parents=True, exist_ok=True)
        
        # Initialize ChromaDB with persistent storage
        self.client = chromadb.PersistentClient(
            path=str(self.kb_path / "chromadb"),
            settings=Settings(
                anonymized_telemetry=False,
                allow_reset=True
            )
        )
        
        # Use a balanced embedding model
        self.embedding_function = embedding_functions.SentenceTransformerEmbeddingFunction(
            model_name="all-MiniLM-L6-v2"
        )
        
        # Initialize collections
        self.collections = self._init_collections()
        
        # Track document hashes for change detection
        self.hash_file = self.kb_path / "document_hashes.json"
        self.document_hashes = self._load_hashes()
        
        # Statistics file
        self.stats_file = self.kb_path / "statistics.json"
    
    def _init_collections(self) -> Dict[str, Any]:
        """Initialize all collections for different document types."""
        collections = {}
        
        # Get unique collection names
        collection_names = set()
        for doc_type in self.DOCUMENT_TYPES:
            collection_names.add(doc_type.collection)
        collection_names.add("general")
        
        for name in collection_names:
            try:
                collection = self.client.get_collection(
                    name=name,
                    embedding_function=self.embedding_function
                )
            except:
                collection = self.client.create_collection(
                    name=name,
                    embedding_function=self.embedding_function,
                    metadata={"description": f"Collection for {name} documents"}
                )
            collections[name] = collection
        
        return collections
    
    def _load_hashes(self) -> Dict[str, str]:
        """Load document hashes from disk."""
        if self.hash_file.exists():
            with open(self.hash_file, 'r') as f:
                return json.load(f)
        return {}
    
    def _save_hashes(self):
        """Save document hashes to disk."""
        with open(self.hash_file, 'w') as f:
            json.dump(self.document_hashes, f, indent=2)
    
    def _get_file_hash(self, file_path: Path) -> str:
        """Calculate hash of file content."""
        with open(file_path, 'rb') as f:
            return hashlib.md5(f.read()).hexdigest()
    
    def _extract_code_documentation(self, content: str, file_ext: str) -> List[Dict[str, str]]:
        """Extract docstrings and comments from code files."""
        docs = []
        
        if file_ext == '.py':
            # Python docstrings
            docstring_pattern = r'"""(.*?)"""'
            for match in re.finditer(docstring_pattern, content, re.DOTALL):
                doc = match.group(1).strip()
                if len(doc) > 50:
                    docs.append({"type": "docstring", "content": doc})
            
            # Python comments
            lines = content.split('\n')
            comment_block = []
            for line in lines:
                if line.strip().startswith('#'):
                    comment_block.append(line.strip()[1:].strip())
                elif comment_block and len(' '.join(comment_block)) > 50:
                    docs.append({"type": "comment", "content": ' '.join(comment_block)})
                    comment_block = []
                else:
                    comment_block = []
        
        elif file_ext in ['.js', '.ts', '.jsx', '.tsx']:
            # JavaScript/TypeScript JSDoc
            jsdoc_pattern = r'/\*\*(.*?)\*/'
            for match in re.finditer(jsdoc_pattern, content, re.DOTALL):
                doc = match.group(1).strip()
                if len(doc) > 50:
                    docs.append({"type": "jsdoc", "content": doc})
        
        return docs
    
    def _chunk_content(self, content: str, chunk_size: int = 1000, overlap: int = 200) -> List[str]:
        """Split content into overlapping chunks for better retrieval."""
        if len(content) <= chunk_size:
            return [content]
        
        chunks = []
        start = 0
        
        while start < len(content):
            end = start + chunk_size
            
            # Find natural break point
            if end < len(content):
                for delimiter in ['\n\n', '. ', '\n', ' ']:
                    break_point = content.rfind(delimiter, start + overlap, end)
                    if break_point != -1:
                        end = break_point + len(delimiter)
                        break
            
            chunks.append(content[start:end])
            start = end - overlap if end < len(content) else end
        
        return chunks
    
    def _parse_architecture_doc(self, content: str, file_path: Path) -> Dict[str, Any]:
        """Parse architecture documentation for enhanced metadata."""
        metadata = {
            "title": "Untitled",
            "status": "active",
            "doc_type": "general"
        }
        
        # Extract title
        for line in content.split('\n'):
            if line.startswith('# '):
                metadata["title"] = line[2:].strip()
                break
        
        # For ADRs, extract status
        if "adrs" in str(file_path).lower():
            metadata["doc_type"] = "adr"
            if "## Status" in content:
                status_section = content.split("## Status")[1].split("##")[0]
                status_line = status_section.strip().split('\n')[0].lower()
                if any(s in status_line for s in ['deprecated', 'superseded']):
                    metadata["status"] = "deprecated"
                elif 'proposed' in status_line:
                    metadata["status"] = "proposed"
                elif 'accepted' in status_line:
                    metadata["status"] = "accepted"
        
        # Identify document type from path
        if "designs" in str(file_path):
            metadata["doc_type"] = "design"
        elif "patterns" in str(file_path):
            metadata["doc_type"] = "pattern"
        elif "standards" in str(file_path):
            metadata["doc_type"] = "standard"
        elif "runbook" in str(file_path).lower():
            metadata["doc_type"] = "runbook"
        elif "rfc" in str(file_path).lower():
            metadata["doc_type"] = "rfc"
        
        return metadata
    
    def _should_index_file(self, file_path: Path, mode: IndexMode) -> Optional[DocumentType]:
        """Determine if a file should be indexed based on mode."""
        # Skip excluded directories
        exclude_dirs = {'.git', 'node_modules', '__pycache__', '.venv', 'venv',
                       'build', 'dist', '.cache', 'coverage', '.vector_db'}
        
        if any(excluded in file_path.parts for excluded in exclude_dirs):
            return None
        
        # Find matching document type
        for doc_type in self.DOCUMENT_TYPES:
            # Check mode compatibility
            if mode != IndexMode.ALL and doc_type.mode != mode:
                continue
            
            # Check extension
            if file_path.suffix not in doc_type.extensions:
                continue
            
            # Check path patterns
            rel_path = str(file_path.relative_to(self.project_root))
            for pattern in doc_type.paths:
                if '*' in pattern:
                    # Handle wildcards
                    import fnmatch
                    if fnmatch.fnmatch(rel_path, pattern):
                        return doc_type
                    if fnmatch.fnmatch(file_path.name, pattern):
                        return doc_type
                else:
                    # Check if pattern is in path
                    if pattern in rel_path:
                        return doc_type
                    # Check exact file match
                    if pattern == file_path.name:
                        return doc_type
        
        return None
    
    def index_file(self, file_path: Path, force: bool = False, mode: IndexMode = IndexMode.ALL) -> bool:
        """Index a single file into the knowledge base."""
        try:
            # Check if file should be indexed
            doc_type = self._should_index_file(file_path, mode)
            if not doc_type:
                return False
            
            # Skip large files
            if file_path.stat().st_size > 10_000_000:  # 10MB
                return False
            
            # Check if file has changed
            current_hash = self._get_file_hash(file_path)
            if not force and str(file_path) in self.document_hashes:
                if self.document_hashes[str(file_path)] == current_hash:
                    return False
            
            # Read file content
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
            except UnicodeDecodeError:
                return False
            
            # Get collection
            collection = self.collections[doc_type.collection]
            
            # Prepare documents to index
            documents_to_index = []
            metadata_base = {
                "file_path": str(file_path.relative_to(self.project_root)),
                "file_name": file_path.name,
                "file_type": file_path.suffix,
                "modified": datetime.fromtimestamp(file_path.stat().st_mtime).isoformat()
            }
            
            # Special handling for architecture docs
            if doc_type.collection == "architecture":
                arch_metadata = self._parse_architecture_doc(content, file_path)
                metadata_base.update(arch_metadata)
            
            # Extract code documentation if applicable
            if file_path.suffix in ['.py', '.js', '.ts', '.jsx', '.tsx']:
                code_docs = self._extract_code_documentation(content, file_path.suffix)
                for doc in code_docs:
                    documents_to_index.append({
                        "content": doc["content"],
                        "metadata": {**metadata_base, "content_type": f"code_{doc['type']}"}
                    })
            
            # Chunk the full content
            chunks = self._chunk_content(content, doc_type.chunk_size, doc_type.chunk_overlap)
            for i, chunk in enumerate(chunks):
                documents_to_index.append({
                    "content": chunk,
                    "metadata": {
                        **metadata_base,
                        "content_type": doc_type.name,
                        "chunk_index": i,
                        "total_chunks": len(chunks)
                    }
                })
            
            # Delete old chunks for this file
            try:
                old_results = collection.get(
                    where={"file_path": metadata_base["file_path"]}
                )
                if old_results["ids"]:
                    collection.delete(ids=old_results["ids"])
            except:
                pass
            
            # Add new chunks
            if documents_to_index:
                ids = []
                documents = []
                metadatas = []
                
                for i, doc in enumerate(documents_to_index):
                    chunk_id = f"{file_path.stem}_{i}_{current_hash[:8]}"
                    ids.append(chunk_id)
                    documents.append(doc["content"])
                    metadatas.append(doc["metadata"])
                
                collection.add(ids=ids, documents=documents, metadatas=metadatas)
            
            # Update hash
            self.document_hashes[str(file_path)] = current_hash
            self._save_hashes()
            
            print(f"✓ Indexed: {file_path.relative_to(self.project_root)} ({len(documents_to_index)} chunks)")
            return True
            
        except Exception as e:
            print(f"✗ Error indexing {file_path}: {e}")
            return False
    
    def index(self, path: Path = None, mode: IndexMode = IndexMode.ALL, 
              force: bool = False, recursive: bool = True) -> int:
        """Index files based on mode."""
        if path is None:
            # Default to docs folder for documentation-focused indexing
            if mode in [IndexMode.ARCHITECTURE, IndexMode.DOCS]:
                path = self.project_root / "docs"
                if not path.exists():
                    path = self.project_root
            else:
                path = self.project_root
        else:
            # Safety check: ensure path is within docs folder
            path = Path(path).resolve()
            docs_path = (self.project_root / "docs").resolve()
            
            # Only allow indexing within docs folder or specific safe locations
            safe_paths = [
                docs_path,
                self.project_root / "CLAUDE.md",
                self.project_root / "README.md"
            ]
            
            is_safe = False
            for safe_path in safe_paths:
                try:
                    if safe_path.is_file():
                        if path == safe_path:
                            is_safe = True
                            break
                    else:
                        # Check if path is within safe directory
                        path.relative_to(safe_path)
                        is_safe = True
                        break
                except ValueError:
                    continue
            
            if not is_safe:
                print(f"⚠️  Error: Path '{path}' is outside the docs folder.")
                print("For safety, indexing is restricted to the /docs directory.")
                print("If you need to index other locations, please update the code.")
                return 0
        
        mode_desc = {
            IndexMode.ALL: "all files",
            IndexMode.ARCHITECTURE: "architecture documentation",
            IndexMode.CODE: "source code",
            IndexMode.DOCS: "documentation",
            IndexMode.CONFIG: "configuration files"
        }
        
        print(f"\n🔍 Indexing {mode_desc[mode]} in {path}...")
        
        # Find all files, excluding hidden directories
        files = []
        exclude_dirs = {'.git', 'node_modules', '__pycache__', '.venv', 'venv',
                       'build', 'dist', '.cache', 'coverage', '.vector_db'}
        
        def should_skip_dir(dir_path: Path) -> bool:
            """Check if directory should be skipped."""
            # Skip if any part of the path is in exclude_dirs
            for part in dir_path.parts:
                if part.startswith('.') or part in exclude_dirs:
                    return True
            return False
        
        if recursive:
            # Walk directory tree, skipping excluded dirs
            for item in path.rglob("*"):
                if item.is_file() and not should_skip_dir(item.parent):
                    files.append(item)
        else:
            files = [f for f in path.glob("*") if f.is_file()]
        
        indexed_count = 0
        total_files = len(files)
        
        for i, file_path in enumerate(files, 1):
            if i % 100 == 0:
                print(f"  Progress: {i}/{total_files} files processed...")
            
            if self.index_file(file_path, force, mode):
                indexed_count += 1
        
        # Save statistics
        self._save_statistics()
        
        print(f"\n✅ Indexed {indexed_count} files (examined {total_files} total)")
        return indexed_count
    
    def search(self, query: str, collection: str = None, limit: int = 5,
               file_type: str = None, content_type: str = None) -> List[Dict]:
        """Search the knowledge base."""
        results_list = []
        
        # Determine collections to search
        if collection:
            collections = [self.collections.get(collection, self.collections["general"])]
        else:
            collections = self.collections.values()
        
        for coll in collections:
            try:
                # Build filter
                where = {}
                if file_type:
                    where["file_type"] = file_type
                if content_type:
                    where["content_type"] = content_type
                
                # Query
                results = coll.query(
                    query_texts=[query],
                    n_results=min(limit, coll.count()),
                    where=where if where else None
                )
                
                if results["documents"][0]:
                    for i, doc in enumerate(results["documents"][0]):
                        results_list.append({
                            "content": doc,
                            "metadata": results["metadatas"][0][i],
                            "distance": results["distances"][0][i] if "distances" in results else 0,
                            "collection": coll.name
                        })
            except:
                continue
        
        # Sort by relevance
        results_list.sort(key=lambda x: x["distance"])
        
        return results_list[:limit]
    
    def clear(self, collection: str = None) -> bool:
        """Clear the knowledge base."""
        try:
            if collection:
                # Clear specific collection
                if collection in self.collections:
                    self.client.delete_collection(collection)
                    self.collections[collection] = self.client.create_collection(
                        name=collection,
                        embedding_function=self.embedding_function
                    )
                    print(f"Cleared collection: {collection}")
            else:
                # Clear all collections
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
    
    def _save_statistics(self):
        """Save knowledge base statistics."""
        stats = {
            "last_updated": datetime.now().isoformat(),
            "total_files": len(self.document_hashes),
            "collections": {}
        }
        
        for name, collection in self.collections.items():
            stats["collections"][name] = {
                "chunks": collection.count(),
                "name": name
            }
        
        with open(self.stats_file, 'w') as f:
            json.dump(stats, f, indent=2)
    
    def stats(self) -> Dict:
        """Get knowledge base statistics."""
        stats = {
            "total_files": len(self.document_hashes),
            "collections": {}
        }
        
        total_chunks = 0
        for name, collection in self.collections.items():
            count = collection.count()
            stats["collections"][name] = count
            total_chunks += count
        
        stats["total_chunks"] = total_chunks
        
        return stats
    
    def print_stats(self):
        """Print formatted statistics."""
        stats = self.stats()
        
        print("\n📊 Knowledge Base Statistics")
        print("=" * 50)
        print(f"📁 Total indexed files: {stats['total_files']}")
        print(f"📄 Total chunks: {stats['total_chunks']}")
        
        if stats['collections']:
            print("\n📚 Collections:")
            for name, count in sorted(stats['collections'].items()):
                print(f"   {name:15} {count:6} chunks")
        
        # Show file type distribution
        type_counts = {}
        for file_path in self.document_hashes:
            ext = Path(file_path).suffix
            type_counts[ext] = type_counts.get(ext, 0) + 1
        
        if type_counts:
            print("\n📝 File Types:")
            for ext, count in sorted(type_counts.items(), key=lambda x: x[1], reverse=True)[:10]:
                print(f"   {ext:10} {count:6} files")


def main():
    """CLI interface for the knowledge base."""
    parser = argparse.ArgumentParser(
        description="Unified Knowledge Base Manager",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Index everything
  python kb.py index
  
  # Index only architecture docs
  python kb.py index --mode arch
  
  # Search for a term
  python kb.py search "authentication flow"
  
  # Search in specific collection
  python kb.py search "database" --collection code
  
  # View statistics
  python kb.py stats
  
  # Clear everything
  python kb.py clear --confirm
        """
    )
    
    subparsers = parser.add_subparsers(dest="command", help="Command to execute")
    
    # Index command
    index_parser = subparsers.add_parser("index", help="Index files into knowledge base")
    index_parser.add_argument("--path", help="Path to index (default: project root)")
    index_parser.add_argument("--mode", choices=["all", "arch", "code", "docs", "config"],
                             default="all", help="What to index")
    index_parser.add_argument("--force", action="store_true", help="Force re-index all files")
    
    # Search command
    search_parser = subparsers.add_parser("search", help="Search the knowledge base")
    search_parser.add_argument("query", help="Search query")
    search_parser.add_argument("--collection", help="Specific collection to search")
    search_parser.add_argument("--limit", type=int, default=5, help="Number of results")
    search_parser.add_argument("--type", help="File type filter (e.g., .py)")
    
    # Stats command
    stats_parser = subparsers.add_parser("stats", help="Show knowledge base statistics")
    
    # Clear command
    clear_parser = subparsers.add_parser("clear", help="Clear the knowledge base")
    clear_parser.add_argument("--collection", help="Clear specific collection only")
    clear_parser.add_argument("--confirm", action="store_true", 
                             help="Confirm clearing (required)")
    
    args = parser.parse_args()
    
    # Initialize knowledge base
    kb = KnowledgeBase()
    
    if args.command == "index":
        path = Path(args.path) if args.path else None
        mode = IndexMode(args.mode)
        kb.index(path, mode, args.force)
    
    elif args.command == "search":
        results = kb.search(args.query, args.collection, args.limit, args.type)
        
        if results:
            print(f"\n🔍 Search results for '{args.query}':\n")
            for i, result in enumerate(results, 1):
                print(f"Result {i}:")
                print(f"  📁 File: {result['metadata']['file_path']}")
                print(f"  📚 Collection: {result['collection']}")
                relevance = (1 - result['distance']) * 100
                print(f"  🎯 Relevance: {relevance:.1f}%")
                
                # Show additional metadata for architecture docs
                if result['collection'] == 'architecture':
                    if 'title' in result['metadata']:
                        print(f"  📝 Title: {result['metadata']['title']}")
                    if 'status' in result['metadata']:
                        print(f"  📊 Status: {result['metadata']['status']}")
                
                preview = result['content'][:200].replace('\n', ' ')
                print(f"  📄 Preview: {preview}...")
                print()
        else:
            print(f"No results found for '{args.query}'")
    
    elif args.command == "stats":
        kb.print_stats()
    
    elif args.command == "clear":
        if not args.confirm:
            print("⚠️  Add --confirm to clear the knowledge base")
            sys.exit(1)
        
        if args.collection:
            kb.clear(args.collection)
        else:
            response = input("Clear entire knowledge base? (yes/no): ")
            if response.lower() == "yes":
                kb.clear()
            else:
                print("Cancelled")
    
    else:
        parser.print_help()


if __name__ == "__main__":
    main()