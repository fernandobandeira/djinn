#!/usr/bin/env python3
"""
Harvest and Store Module for knowledge-harvester
Provides automatic saving and indexing of harvested content
"""

import asyncio
import hashlib
import subprocess
from datetime import datetime
from pathlib import Path
from typing import Dict, Optional
import yaml

from crawl4ai import AsyncWebCrawler, CrawlerRunConfig, CacheMode

class HarvestAndStore:
    """Automatically harvest, save, and index web content."""
    
    def __init__(self, base_path: str = None):
        # Use relative path from .vector_db directory
        if base_path is None:
            # We're in .vector_db, so parent is djinn root
            self.base_path = Path(__file__).parent.parent
        else:
            self.base_path = Path(base_path)
        
        self.harvested_dir = self.base_path / "harvested"
        self.kb_path = self.base_path / ".vector_db"
        
        # Load harvesting profiles
        profiles_file = self.base_path / ".claude/resources/shared/harvesting-profiles.yaml"
        with open(profiles_file, 'r') as f:
            self.profiles = yaml.safe_load(f)['profiles']
    
    def generate_filename(self, topic: str, profile: str, url: str) -> str:
        """Generate a unique filename for harvested content."""
        # Create slug from topic
        topic_slug = topic.lower().replace(' ', '_')[:50]
        
        # Create short hash from URL
        url_hash = hashlib.md5(url.encode()).hexdigest()[:6]
        
        # Current date
        date_str = datetime.now().strftime("%Y-%m-%d")
        
        # Combine
        filename = f"{date_str}_{profile}_{topic_slug}_{url_hash}.md"
        return filename
    
    def determine_subdirectory(self, profile: str) -> Path:
        """Determine which subdirectory to save to based on profile."""
        mapping = {
            'troubleshooting': 'troubleshooting',
            'api_documentation': 'documentation',
            'deep_research': 'research',
            'security_advisories': 'web',
            'code_examples': 'documentation',
            'quick_reference': 'web'
        }
        
        subdir = mapping.get(profile, 'web')
        return self.harvested_dir / subdir
    
    def check_duplicate(self, url: str) -> Optional[Dict]:
        """Check if URL has already been harvested."""
        # Check all subdirectories in harvested
        for subdir in self.harvested_dir.glob("*"):
            if subdir.is_dir():
                for file in subdir.glob("*.md"):
                    try:
                        with open(file, 'r') as f:
                            content = f.read()
                            # Extract metadata from frontmatter
                            if content.startswith('---'):
                                metadata_end = content.find('---', 3)
                                if metadata_end > 0:
                                    metadata_str = content[3:metadata_end]
                                    metadata = yaml.safe_load(metadata_str)
                                    if metadata.get('source') == url:
                                        return {
                                            'file': str(file),
                                            'harvested_at': metadata.get('harvested_at'),
                                            'topic': metadata.get('topic')
                                        }
                    except:
                        continue
        return None
    
    async def harvest_and_save(
        self, 
        url: str,
        topic: str,
        profile: str = 'quick_reference',
        agent_context: str = 'unknown',
        force: bool = False
    ) -> Dict:
        """Harvest content and automatically save/index it."""
        
        # Check for duplicates unless forcing
        if not force:
            duplicate = self.check_duplicate(url)
            if duplicate:
                return {
                    'status': 'duplicate',
                    'message': f"Already harvested on {duplicate['harvested_at']}",
                    'filepath': duplicate['file'],
                    'original_topic': duplicate['topic']
                }
        
        # Get profile configuration
        profile_config = self.profiles.get(profile, self.profiles['quick_reference'])
        
        # Configure crawler
        config = CrawlerRunConfig(
            cache_mode=CacheMode.ENABLED if profile != 'security_advisories' else CacheMode.BYPASS,
            wait_until=profile_config['config'].get('wait_until', 'domcontentloaded'),
            page_timeout=profile_config['config'].get('timeout', 30000),
            scan_full_page=profile_config['config'].get('scan_full_page', False),
            delay_before_return_html=profile_config['config'].get('delay_before_return', 0.3)
        )
        
        # Harvest content
        async with AsyncWebCrawler() as crawler:
            result = await crawler.arun(url=url, config=config)
            
            if not result.success:
                return {
                    'status': 'failed',
                    'error': result.error_message
                }
            
            # Prepare content with metadata
            metadata = {
                'source': url,
                'harvested_at': datetime.now().isoformat(),
                'profile': profile,
                'agent_context': agent_context,
                'topic': topic,
                'confidence': 0.95 if result.markdown else 0.0
            }
            
            # Format content
            content = f"""---
{yaml.dump(metadata, default_flow_style=False)}---

# {topic}

{result.markdown if result.markdown else "No content extracted"}

## Source Information
- URL: {url}
- Harvested: {metadata['harvested_at']}
- Profile: {profile}
- Agent: {agent_context}
"""
            
            # Generate filename and save
            filename = self.generate_filename(topic, profile, url)
            subdir = self.determine_subdirectory(profile)
            subdir.mkdir(parents=True, exist_ok=True)
            
            filepath = subdir / filename
            filepath.write_text(content)
            
            # Always index immediately (no background process needed)
            self.index_file(filepath)
            
            return {
                'status': 'success',
                'filepath': str(filepath),
                'content_length': len(result.markdown) if result.markdown else 0,
                'metadata': metadata,
                'auto_indexed': not self.is_auto_indexer_running()
            }
    
    def is_auto_indexer_running(self) -> bool:
        """Check if auto_indexer.py is running."""
        try:
            result = subprocess.run(
                ["pgrep", "-f", "auto_indexer.py"],
                capture_output=True
            )
            return result.returncode == 0
        except:
            return False
    
    def index_file(self, filepath: Path):
        """Manually index a file in the KB."""
        try:
            subprocess.run(
                ["./kb", "index", str(filepath)],
                cwd=self.kb_path,
                check=True,
                capture_output=True
            )
            print(f"  ✅ Indexed: {filepath.name}")
        except subprocess.CalledProcessError:
            print(f"  ⚠️  Failed to index: {filepath.name}")


# CLI interface for command-line usage
async def main():
    """CLI interface for harvest_and_store."""
    import argparse
    import json
    
    parser = argparse.ArgumentParser(description='Harvest and store web content')
    parser.add_argument('--url', required=True, help='URL to harvest')
    parser.add_argument('--topic', required=True, help='Topic description')
    parser.add_argument('--profile', default='quick_reference', 
                       choices=['quick_reference', 'deep_research', 'code_examples', 
                               'api_documentation', 'troubleshooting', 'security_advisories'],
                       help='Harvesting profile to use')
    parser.add_argument('--agent', default='unknown', help='Calling agent context')
    parser.add_argument('--json', action='store_true', help='Output as JSON')
    parser.add_argument('--force', action='store_true', help='Force re-harvest even if duplicate exists')
    
    args = parser.parse_args()
    
    harvester = HarvestAndStore()
    
    result = await harvester.harvest_and_save(
        url=args.url,
        topic=args.topic,
        profile=args.profile,
        agent_context=args.agent,
        force=args.force
    )
    
    if args.json:
        print(json.dumps(result))
    else:
        if result['status'] == 'success':
            print(f"✅ Harvested: {args.topic}")
            print(f"📄 Saved to: {result['filepath']}")
            print(f"📏 Size: {result['content_length']} chars")
            print(f"🔍 Indexed: Yes")
        elif result['status'] == 'duplicate':
            print(f"⚠️  Duplicate: {args.url}")
            print(f"📄 Already harvested: {result['filepath']}")
            print(f"📅 Original harvest: {result['message']}")
            print(f"📌 Original topic: {result['original_topic']}")
            print(f"💡 Use --force to re-harvest")
        else:
            print(f"❌ Failed: {result.get('error', 'Unknown error')}")

if __name__ == "__main__":
    asyncio.run(main())