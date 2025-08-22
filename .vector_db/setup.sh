#!/bin/bash
# Setup script for Djinn Knowledge Base and Dependencies
# Run this after cloning the repository

set -e  # Exit on error

echo "🚀 Setting up Djinn Knowledge Base and Dependencies..."
echo "=================================================="

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo -e "${YELLOW}📦 Step 1: Installing Python dependencies...${NC}"
cd "$SCRIPT_DIR"

# Check if uv is installed
if ! command -v uv &> /dev/null; then
    echo "Installing uv (fast Python package manager)..."
    curl -LsSf https://astral.sh/uv/install.sh | sh
    source $HOME/.cargo/env
fi

# Create virtual environment and install dependencies
if [ ! -d ".venv" ]; then
    uv venv
fi

# Install KB dependencies
uv pip install qdrant-client sentence-transformers tiktoken

echo -e "${GREEN}✅ Python dependencies installed${NC}"

# Setup crawl4ai if not already installed
echo -e "${YELLOW}🕷️ Step 2: Setting up crawl4ai...${NC}"

# Check if crawl4ai is already set up
if ! uv run python -c "import crawl4ai" 2>/dev/null; then
    echo "Installing crawl4ai..."
    uv pip install crawl4ai
    
    # Install playwright for crawl4ai
    echo "Installing Playwright browsers..."
    uv run playwright install chromium
    
    echo -e "${GREEN}✅ crawl4ai setup complete${NC}"
else
    echo -e "${GREEN}✅ crawl4ai already installed${NC}"
fi

# Initialize Qdrant KB
echo -e "${YELLOW}🔍 Step 3: Initializing Qdrant Knowledge Base...${NC}"

# Note about Qdrant
echo "ℹ️  Qdrant runs in local mode - no server needed!"
echo "   All data is stored locally in .vector_db/qdrant_storage/"

# Check if KB is already initialized
if [ ! -d "$SCRIPT_DIR/qdrant_storage" ]; then
    echo "Creating Qdrant storage directory..."
    mkdir -p "$SCRIPT_DIR/qdrant_storage"
fi

# Index the documentation
echo -e "${YELLOW}📚 Step 4: Indexing project documentation...${NC}"

echo "Indexing /docs directory..."
cd "$PROJECT_ROOT"
./.vector_db/kb index --path docs

if [ -d "$PROJECT_ROOT/harvested" ]; then
    echo "Indexing /harvested directory (external content)..."
    ./.vector_db/kb index --path harvested
fi

if [ -d "$PROJECT_ROOT/zettelkasten" ]; then
    echo "Indexing /zettelkasten directory (learning notes)..."
    ./.vector_db/kb index --path zettelkasten
fi

# Show KB statistics
echo -e "${YELLOW}📊 Step 5: Verifying installation...${NC}"
./.vector_db/kb stats

echo ""
echo -e "${GREEN}🎉 Setup complete!${NC}"
echo ""
echo "📖 Quick Start Guide:"
echo "   Search:  ./.vector_db/kb search \"your query\""
echo "   Index:   ./.vector_db/kb index --path <directory>"
echo "   Stats:   ./.vector_db/kb stats"
echo ""
echo "💡 Tips:"
echo "   - KB uses semantic search - finds concepts, not just keywords"
echo "   - Search is context-aware: use --agent architect|developer|analyst"
echo "   - All data is stored locally - no external services needed"
echo ""
echo "📚 Full documentation: ./.vector_db/KB-INSTRUCTIONS.md"