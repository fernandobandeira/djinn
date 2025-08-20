#!/usr/bin/env -S uv run python
"""
Auto-Indexer for Knowledge Base
Watches for file changes and automatically updates the index
"""

import os
import sys
import time
import json
import threading
from pathlib import Path
from typing import Set, Dict, List
from datetime import datetime, timedelta
from queue import Queue, Empty

try:
    from watchdog.observers import Observer
    from watchdog.events import FileSystemEventHandler, FileModifiedEvent, FileCreatedEvent
except ImportError:
    print("Installing watchdog...")
    os.system("pip install watchdog")
    from watchdog.observers import Observer
    from watchdog.events import FileSystemEventHandler, FileModifiedEvent, FileCreatedEvent

# Import our enhanced KB
sys.path.insert(0, str(Path(__file__).parent))
from kb_enhanced import EnhancedKnowledgeBase


class KBFileWatcher(FileSystemEventHandler):
    """Watches for file changes and triggers indexing."""
    
    def __init__(self, kb: EnhancedKnowledgeBase, auto_index: bool = True):
        self.kb = kb
        self.auto_index = auto_index
        self.pending_queue = Queue()
        self.processing_thread = None
        self.running = False
        
        # Patterns to watch
        self.watch_extensions = {'.md', '.py', '.js', '.ts', '.jsx', '.tsx', '.yaml', '.yml', '.json'}
        self.watch_dirs = {
            'docs', 'zettelkasten', '.claude', 'src', 'lib', 'harvested'
        }
        
        # Patterns to ignore
        self.ignore_patterns = {
            '.git', 'node_modules', '__pycache__', '.venv', 'venv',
            'build', 'dist', '.cache', '.vector_db', '.DS_Store'
        }
        
        # Debounce settings
        self.debounce_seconds = 2
        self.last_event_time = {}
        
        # Statistics
        self.stats = {
            'files_detected': 0,
            'files_indexed': 0,
            'files_skipped': 0,
            'errors': 0,
            'start_time': datetime.now().isoformat()
        }
        
        # Start processing thread
        if auto_index:
            self.start_processing()
    
    def should_process(self, path: str) -> bool:
        """Check if file should be processed."""
        path_obj = Path(path)
        
        # Check if it's a file
        if not path_obj.is_file():
            return False
        
        # Check extension
        if path_obj.suffix not in self.watch_extensions:
            return False
        
        # Check for ignore patterns
        path_str = str(path_obj)
        for pattern in self.ignore_patterns:
            if pattern in path_str:
                return False
        
        # Check if in watch directories
        path_parts = path_obj.parts
        if not any(watch_dir in path_parts for watch_dir in self.watch_dirs):
            # Also check if it's a root-level markdown file
            if path_obj.suffix == '.md' and path_obj.parent == Path.cwd():
                return True
            return False
        
        return True
    
    def debounce_event(self, path: str) -> bool:
        """Debounce file events to avoid multiple triggers."""
        current_time = time.time()
        
        if path in self.last_event_time:
            if current_time - self.last_event_time[path] < self.debounce_seconds:
                return False
        
        self.last_event_time[path] = current_time
        return True
    
    def on_modified(self, event):
        """Handle file modification events."""
        if event.is_directory:
            return
        
        if not self.should_process(event.src_path):
            return
        
        if not self.debounce_event(event.src_path):
            return
        
        self.stats['files_detected'] += 1
        print(f"📝 Detected change: {Path(event.src_path).name}")
        
        if self.auto_index:
            self.pending_queue.put({
                'path': event.src_path,
                'event_type': 'modified',
                'timestamp': datetime.now().isoformat()
            })
    
    def on_created(self, event):
        """Handle file creation events."""
        if event.is_directory:
            return
        
        if not self.should_process(event.src_path):
            return
        
        self.stats['files_detected'] += 1
        print(f"✨ Detected new file: {Path(event.src_path).name}")
        
        if self.auto_index:
            # Give a bit of time for file to be fully written
            time.sleep(0.5)
            self.pending_queue.put({
                'path': event.src_path,
                'event_type': 'created',
                'timestamp': datetime.now().isoformat()
            })
    
    def start_processing(self):
        """Start the background processing thread."""
        if not self.running:
            self.running = True
            self.processing_thread = threading.Thread(target=self._process_queue, daemon=True)
            self.processing_thread.start()
            print("✅ Auto-indexing enabled")
    
    def stop_processing(self):
        """Stop the background processing thread."""
        self.running = False
        if self.processing_thread:
            self.processing_thread.join(timeout=5)
            print("⏹️  Auto-indexing stopped")
    
    def _process_queue(self):
        """Process pending files in the background."""
        batch_size = 5
        batch_wait = 3  # seconds
        batch = []
        last_batch_time = time.time()
        
        while self.running:
            try:
                # Try to get an item with timeout
                item = self.pending_queue.get(timeout=1)
                batch.append(item)
                
                # Process batch if it's full or enough time has passed
                if len(batch) >= batch_size or (time.time() - last_batch_time) > batch_wait:
                    self._process_batch(batch)
                    batch = []
                    last_batch_time = time.time()
                    
            except Empty:
                # Process any remaining items in batch
                if batch:
                    self._process_batch(batch)
                    batch = []
                    last_batch_time = time.time()
                continue
    
    def _process_batch(self, batch: List[Dict]):
        """Process a batch of files."""
        if not batch:
            return
        
        print(f"\n🔄 Processing batch of {len(batch)} files...")
        
        for item in batch:
            try:
                path = Path(item['path'])
                if path.exists():
                    success = self.kb.index_file(path, force=True)
                    if success:
                        self.stats['files_indexed'] += 1
                        print(f"  ✅ Indexed: {path.name}")
                    else:
                        self.stats['files_skipped'] += 1
                        print(f"  ⏭️  Skipped: {path.name}")
                else:
                    print(f"  ❌ File no longer exists: {path.name}")
                    
            except Exception as e:
                self.stats['errors'] += 1
                print(f"  ❌ Error processing {item['path']}: {e}")
        
        # Save status
        self.kb.status_tracker.save()
        print(f"📊 Stats: Indexed={self.stats['files_indexed']}, Skipped={self.stats['files_skipped']}, Errors={self.stats['errors']}")
    
    def get_stats(self) -> Dict:
        """Get watcher statistics."""
        runtime = datetime.now() - datetime.fromisoformat(self.stats['start_time'])
        return {
            **self.stats,
            'runtime_seconds': runtime.total_seconds(),
            'pending_count': self.pending_queue.qsize()
        }


class KBAutoIndexer:
    """Main auto-indexer that manages file watching and periodic indexing."""
    
    def __init__(self, project_root: str = None):
        self.project_root = Path(project_root) if project_root else Path.cwd()
        self.kb = EnhancedKnowledgeBase(self.project_root)
        self.observer = Observer()
        self.watcher = KBFileWatcher(self.kb)
        
        # Paths to watch
        self.watch_paths = [
            self.project_root / "docs",
            self.project_root / "zettelkasten",
            self.project_root / ".claude",
            self.project_root / "src",
            self.project_root / "harvested",
            self.project_root  # For root-level files like README.md
        ]
        
        # Schedule periodic full reindex
        self.last_full_index = None
        self.full_index_interval = timedelta(hours=24)  # Daily full reindex
    
    def start(self):
        """Start the auto-indexer."""
        print(f"🚀 Starting KB Auto-Indexer for {self.project_root}")
        
        # Schedule watchers for each path
        scheduled_paths = []
        for path in self.watch_paths:
            if path.exists():
                self.observer.schedule(self.watcher, str(path), recursive=True)
                scheduled_paths.append(str(path))
                print(f"  👁️  Watching: {path}")
        
        if not scheduled_paths:
            print("❌ No valid paths to watch!")
            return
        
        # Start observer
        self.observer.start()
        print(f"✅ Auto-indexer started, watching {len(scheduled_paths)} paths")
        
        # Initial status
        status = self.kb.get_status()
        print(f"\n📊 Initial KB Status:")
        print(f"  Total documents: {status['total_documents']}")
        print(f"  Health: {status['health']['health']}")
        
        # Run main loop
        try:
            while True:
                time.sleep(60)  # Check every minute
                
                # Check if we need a full reindex
                if self.should_full_reindex():
                    self.run_full_reindex()
                
                # Print periodic status
                if int(time.time()) % 300 == 0:  # Every 5 minutes
                    self.print_status()
                    
        except KeyboardInterrupt:
            print("\n⏹️  Shutting down auto-indexer...")
            self.stop()
    
    def stop(self):
        """Stop the auto-indexer."""
        self.watcher.stop_processing()
        self.observer.stop()
        self.observer.join()
        print("✅ Auto-indexer stopped")
        
        # Final stats
        print("\n📊 Final Statistics:")
        stats = self.watcher.get_stats()
        print(f"  Files detected: {stats['files_detected']}")
        print(f"  Files indexed: {stats['files_indexed']}")
        print(f"  Files skipped: {stats['files_skipped']}")
        print(f"  Errors: {stats['errors']}")
        print(f"  Runtime: {stats['runtime_seconds']:.1f} seconds")
    
    def should_full_reindex(self) -> bool:
        """Check if we should run a full reindex."""
        if self.last_full_index is None:
            return False  # Don't run on startup
        
        return datetime.now() - self.last_full_index > self.full_index_interval
    
    def run_full_reindex(self):
        """Run a full reindex of all documents."""
        print("\n🔄 Running scheduled full reindex...")
        self.watcher.auto_index = False  # Pause auto-indexing
        
        try:
            indexed = self.kb.index_directory(force=False)  # Only changed files
            self.last_full_index = datetime.now()
            print(f"✅ Full reindex complete: {indexed} files updated")
        finally:
            self.watcher.auto_index = True  # Resume auto-indexing
    
    def print_status(self):
        """Print current status."""
        status = self.kb.get_status()
        stats = self.watcher.get_stats()
        
        print(f"\n📊 Status Update:")
        print(f"  KB Health: {status['health']['health']}")
        print(f"  Pending files: {status['pending_files']}")
        print(f"  Watcher queue: {stats['pending_count']}")
        print(f"  Files indexed (session): {stats['files_indexed']}")


def main():
    """Main entry point for auto-indexer."""
    import argparse
    
    parser = argparse.ArgumentParser(description="KB Auto-Indexer")
    parser.add_argument("--path", help="Project root path", default=".")
    parser.add_argument("--no-watch", action="store_true", help="Disable file watching")
    parser.add_argument("--reindex", action="store_true", help="Run full reindex and exit")
    
    args = parser.parse_args()
    
    if args.reindex:
        # Just run a full reindex and exit
        kb = EnhancedKnowledgeBase(args.path)
        print("🔄 Running full reindex...")
        indexed = kb.index_directory(force=True)
        print(f"✅ Indexed {indexed} files")
        
        # Show status
        status = kb.get_status()
        print(f"\n📊 KB Status:")
        print(f"  Total documents: {status['total_documents']}")
        print(f"  Health: {status['health']['health']}")
        for name, info in status['collections'].items():
            if info['count'] > 0:
                print(f"  {name}: {info['count']} chunks")
    else:
        # Run auto-indexer
        indexer = KBAutoIndexer(args.path)
        
        if args.no_watch:
            print("File watching disabled, running in manual mode")
            # Just show status
            status = indexer.kb.get_status()
            print(json.dumps(status, indent=2))
        else:
            indexer.start()


if __name__ == "__main__":
    main()