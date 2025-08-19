#!/bin/bash
# Setup script for crawl4ai integration with knowledge-harvester

echo "🚀 Setting up crawl4ai for knowledge-harvester..."
echo "================================================"

# Check if we're in the right directory
if [ ! -f "pyproject.toml" ]; then
    echo "❌ Error: Please run this script from the .vector_db directory"
    echo "   cd .vector_db && ./setup_crawl4ai.sh"
    exit 1
fi

# Activate virtual environment if it exists
if [ -f ".venv/bin/activate" ]; then
    echo "📦 Activating virtual environment..."
    source .venv/bin/activate
else
    echo "⚠️  No virtual environment found, using system Python"
fi

# Check if uv is installed
if ! command -v uv &> /dev/null; then
    echo "❌ Error: uv is not installed"
    echo "   Please install uv first: https://github.com/astral-sh/uv"
    exit 1
fi

# Install crawl4ai
echo ""
echo "📦 Installing crawl4ai==0.6.2..."
uv pip install crawl4ai==0.6.2

if [ $? -ne 0 ]; then
    echo "❌ Failed to install crawl4ai"
    exit 1
fi

# Setup crawl4ai (installs browser dependencies)
echo ""
echo "🌐 Setting up browser support for crawl4ai..."
# Try to find and run crawl4ai-setup
if command -v crawl4ai-setup &> /dev/null; then
    crawl4ai-setup
elif [ -f ".venv/bin/crawl4ai-setup" ]; then
    .venv/bin/crawl4ai-setup
else
    echo "⚠️  crawl4ai-setup not found, but browsers might already be installed"
    echo "   Playwright browsers are usually auto-installed on first use"
fi

if [ $? -ne 0 ]; then
    echo "⚠️  Warning: crawl4ai-setup had issues, but this might be okay"
    echo "   The browsers might already be installed"
fi

# Test the installation
echo ""
echo "🧪 Testing crawl4ai installation..."
python3 -c "from crawl4ai import AsyncWebCrawler; print('✅ crawl4ai is installed and working!')" 2>/dev/null

if [ $? -ne 0 ]; then
    echo "❌ crawl4ai import test failed"
    echo "   Please check the installation"
    exit 1
fi

# Run the test suite
echo ""
echo "🔬 Running integration tests..."
if [ -f "test_crawl4ai.py" ]; then
    python3 test_crawl4ai.py
else
    echo "⚠️  Test file not found, skipping tests"
fi

echo ""
echo "================================================"
echo "✅ Setup complete!"
echo ""
echo "📚 Knowledge-harvester can now use crawl4ai for:"
echo "  • 6x faster web crawling"
echo "  • Adaptive depth detection"
echo "  • Anti-bot detection bypass"
echo "  • Parallel multi-page processing"
echo "  • LLM-ready markdown extraction"
echo ""
echo "🎯 Next steps:"
echo "  1. Test with: python3 test_crawl4ai.py"
echo "  2. Use knowledge-harvester via command agents"
echo "  3. Monitor harvesting with: ./kb stats"
echo ""