#!/usr/bin/env -S uv run python
"""
Migration script to upgrade KB to enhanced version
Reindexes all documents with new embedding model
"""

import os
import sys
import json
import shutil
from pathlib import Path
from datetime import datetime

# Add parent directory to path
sys.path.insert(0, str(Path(__file__).parent.parent))

from kb_enhanced import EnhancedKnowledgeBase


def migrate_kb():
    """Migrate existing KB to enhanced version."""
    print("🔄 Starting KB Migration to Enhanced Version...")
    print("")
    
    kb_path = Path(__file__).parent
    chromadb_path = kb_path / "chromadb"
    backup_path = kb_path / f"chromadb_backup_{datetime.now().strftime('%Y%m%d_%H%M%S')}"
    
    # Check if we have an existing DB
    if chromadb_path.exists():
        print(f"📦 Backing up existing ChromaDB to {backup_path.name}...")
        shutil.copytree(chromadb_path, backup_path)
        print("✅ Backup complete")
        print("")
        
        # Clear the existing DB
        print("🗑️  Clearing old embeddings...")
        shutil.rmtree(chromadb_path)
        print("✅ Old embeddings cleared")
        print("")
    
    # Initialize new enhanced KB
    print("🚀 Initializing Enhanced KB with new embedding model...")
    kb = EnhancedKnowledgeBase()
    print("")
    
    # Get list of files to reindex from document_hashes.json
    hash_file = kb_path / "document_hashes.json"
    if hash_file.exists():
        with open(hash_file, 'r') as f:
            document_hashes = json.load(f)
        
        files_to_index = list(document_hashes.keys())
        print(f"📚 Found {len(files_to_index)} documents to reindex")
        print("")
        
        # Clear the hashes to force reindexing
        kb.document_hashes = {}
        kb._save_hashes()
        
        # Reindex all files
        print("🔍 Reindexing with enhanced features...")
        print("  • Better embeddings (BAAI/bge-large-en-v1.5)")
        print("  • Enhanced metadata extraction")
        print("  • Document categorization")
        print("")
        
        indexed_count = 0
        failed_files = []
        
        for i, file_path in enumerate(files_to_index, 1):
            if i % 10 == 0:
                print(f"  Progress: {i}/{len(files_to_index)} files...")
            
            try:
                path = Path(file_path)
                if path.exists():
                    if kb.index_file(path, force=True):
                        indexed_count += 1
                else:
                    print(f"  ⚠️  File no longer exists: {file_path}")
            except Exception as e:
                failed_files.append((file_path, str(e)))
                print(f"  ❌ Error indexing {file_path}: {e}")
        
        print("")
        print(f"✅ Migration complete!")
        print(f"  • Successfully indexed: {indexed_count} files")
        if failed_files:
            print(f"  • Failed: {len(failed_files)} files")
            for file_path, error in failed_files[:5]:
                print(f"    - {Path(file_path).name}: {error}")
        
        # Show new status
        print("")
        print("📊 New KB Status:")
        status = kb.get_status()
        print(f"  • Total documents: {status['total_documents']}")
        print(f"  • Health: {status['health']['health']}")
        print(f"  • Collections:")
        for name, info in status['collections'].items():
            if info['count'] > 0:
                print(f"    - {name}: {info['count']} chunks")
        
        # Test search
        print("")
        print("🧪 Testing enhanced search...")
        test_results = kb.search("architecture patterns", limit=3)
        if test_results:
            print(f"  ✅ Search working! Found {len(test_results)} results")
        else:
            print("  ⚠️  No results found (you may need to index more files)")
        
        print("")
        print("🎉 Migration successful!")
        print("")
        print("💡 Next steps:")
        print("  1. Test search with agent context:")
        print("     ./kb search 'your query' --agent architect")
        print("")
        print("  2. Start auto-indexer for real-time updates:")
        print("     python kb_auto_indexer.py")
        print("")
        print("  3. If you encounter issues, restore backup with:")
        print(f"     mv {backup_path.name} chromadb")
        
    else:
        print("❌ No existing document_hashes.json found")
        print("   Running full index instead...")
        
        # Run full index
        indexed = kb.index_directory(force=True)
        print(f"✅ Indexed {indexed} files")
    
    return True


if __name__ == "__main__":
    try:
        success = migrate_kb()
        sys.exit(0 if success else 1)
    except Exception as e:
        print(f"❌ Migration failed: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)