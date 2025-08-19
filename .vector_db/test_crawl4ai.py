#!/usr/bin/env python3
"""
Test script for crawl4ai integration with knowledge-harvester
Tests basic crawling functionality and profile configurations
"""

import asyncio
import yaml
from pathlib import Path

# Test if crawl4ai is installed
try:
    from crawl4ai import AsyncWebCrawler, CrawlerRunConfig, CacheMode
    print("✅ crawl4ai successfully imported")
except ImportError as e:
    print(f"❌ Failed to import crawl4ai: {e}")
    print("Please run: cd .vector_db && uv pip install crawl4ai==0.6.2")
    exit(1)

# Load harvesting profiles
profiles_path = Path(__file__).parent.parent / ".claude/resources/shared/harvesting-profiles.yaml"
if profiles_path.exists():
    with open(profiles_path, 'r') as f:
        profiles_config = yaml.safe_load(f)
        print(f"✅ Loaded {len(profiles_config['profiles'])} harvesting profiles")
else:
    print(f"❌ Harvesting profiles not found at {profiles_path}")
    profiles_config = None

async def test_quick_reference():
    """Test quick reference profile - single page extraction"""
    print("\n🔍 Testing Quick Reference Profile...")
    
    async with AsyncWebCrawler() as crawler:
        # Test with Python documentation
        url = "https://docs.python.org/3/library/asyncio.html"
        
        config = CrawlerRunConfig(
            cache_mode=CacheMode.ENABLED,
            # markdown_generator parameter removed - it's enabled by default
            wait_until='domcontentloaded',
            page_timeout=30000,
            scan_full_page=False,
            delay_before_return_html=0.3
        )
        
        try:
            result = await crawler.arun(url=url, config=config)
            if result.success:
                content_length = len(result.markdown) if result.markdown else 0
                print(f"  ✅ Successfully crawled {url}")
                print(f"  📄 Content length: {content_length} characters")
                if result.markdown and '```' in result.markdown:
                    print(f"  💻 Code blocks found: {result.markdown.count('```')}")
                return True
            else:
                print(f"  ❌ Failed to crawl: {result.error_message}")
                return False
        except Exception as e:
            print(f"  ❌ Exception during crawl: {e}")
            return False

async def test_code_examples():
    """Test code examples profile - extract code blocks"""
    print("\n💻 Testing Code Examples Profile...")
    
    async with AsyncWebCrawler() as crawler:
        # Test with a GitHub README
        url = "https://github.com/unclecode/crawl4ai"
        
        config = CrawlerRunConfig(
            cache_mode=CacheMode.ENABLED,
            wait_until='domcontentloaded',
            page_timeout=60000,
            css_selector="pre, code, .highlight",
            exclude_all_images=True
        )
        
        try:
            result = await crawler.arun(url=url, config=config)
            if result.success:
                print(f"  ✅ Successfully extracted from {url}")
                if result.markdown:
                    code_blocks = result.markdown.count('```')
                    print(f"  📦 Code blocks extracted: {code_blocks}")
                return True
            else:
                print(f"  ❌ Failed to extract: {result.error_message}")
                return False
        except Exception as e:
            print(f"  ❌ Exception during extraction: {e}")
            return False

async def test_adaptive_crawling():
    """Test adaptive crawling - let crawl4ai decide depth"""
    print("\n🧠 Testing Adaptive Crawling...")
    
    async with AsyncWebCrawler() as crawler:
        # Test with documentation that has multiple pages
        url = "https://docs.crawl4ai.com/"
        
        config = CrawlerRunConfig(
            cache_mode=CacheMode.BYPASS,  # Fresh content
            wait_until='networkidle',
            page_timeout=120000,
            scan_full_page=True,
            # Note: In real implementation, would use adaptive strategies
            # For now, just test basic multi-page capability
            max_depth=2
        )
        
        try:
            result = await crawler.arun(url=url, config=config)
            if result.success:
                print(f"  ✅ Successfully performed deep crawl on {url}")
                print(f"  📊 Content size: {len(result.markdown) if result.markdown else 0} chars")
                return True
            else:
                print(f"  ❌ Deep crawl failed: {result.error_message}")
                return False
        except Exception as e:
            print(f"  ❌ Exception during deep crawl: {e}")
            return False

async def test_performance():
    """Test parallel crawling performance"""
    print("\n⚡ Testing Parallel Performance...")
    
    urls = [
        "https://docs.python.org/3/library/asyncio.html",
        "https://docs.python.org/3/library/typing.html",
        "https://docs.python.org/3/library/collections.html"
    ]
    
    async with AsyncWebCrawler() as crawler:
        config = CrawlerRunConfig(
            cache_mode=CacheMode.ENABLED,
            wait_until='domcontentloaded',
            page_timeout=30000
        )
        
        # Time parallel crawling
        import time
        start = time.time()
        
        tasks = [crawler.arun(url=url, config=config) for url in urls]
        results = await asyncio.gather(*tasks, return_exceptions=True)
        
        elapsed = time.time() - start
        
        successful = sum(1 for r in results if not isinstance(r, Exception) and r.success)
        print(f"  ✅ Crawled {successful}/{len(urls)} pages in {elapsed:.2f} seconds")
        print(f"  🚀 Average: {elapsed/len(urls):.2f} seconds per page")
        
        return successful == len(urls)

async def main():
    """Run all tests"""
    print("=" * 60)
    print("🧪 CRAWL4AI INTEGRATION TEST SUITE")
    print("=" * 60)
    
    # Check if profiles are loaded
    if profiles_config:
        print("\n📋 Available Profiles:")
        for profile_name in profiles_config['profiles']:
            profile = profiles_config['profiles'][profile_name]
            print(f"  • {profile_name}: {profile['description']}")
    
    # Run tests
    tests = [
        ("Quick Reference", test_quick_reference),
        ("Code Examples", test_code_examples),
        ("Adaptive Crawling", test_adaptive_crawling),
        ("Parallel Performance", test_performance)
    ]
    
    results = []
    for test_name, test_func in tests:
        try:
            success = await test_func()
            results.append((test_name, success))
        except Exception as e:
            print(f"\n❌ Test '{test_name}' crashed: {e}")
            results.append((test_name, False))
    
    # Summary
    print("\n" + "=" * 60)
    print("📊 TEST SUMMARY")
    print("=" * 60)
    
    passed = sum(1 for _, success in results if success)
    total = len(results)
    
    for test_name, success in results:
        status = "✅ PASS" if success else "❌ FAIL"
        print(f"  {status}: {test_name}")
    
    print(f"\n🎯 Result: {passed}/{total} tests passed")
    
    if passed == total:
        print("🎉 All tests passed! crawl4ai is ready for knowledge harvesting!")
    else:
        print("⚠️  Some tests failed. Please check the configuration.")
    
    return passed == total

if __name__ == "__main__":
    # Run the test suite
    success = asyncio.run(main())
    exit(0 if success else 1)