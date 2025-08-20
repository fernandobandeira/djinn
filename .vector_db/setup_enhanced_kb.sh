#!/bin/bash
# Setup script for Enhanced Knowledge Base

echo "🚀 Setting up Enhanced Knowledge Base..."
echo ""

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if we're in the right directory
if [ ! -f "kb_enhanced.py" ]; then
    echo -e "${RED}Error: Please run this script from the .vector_db directory${NC}"
    exit 1
fi

echo "📦 Installing Python dependencies..."
echo ""

# Try uv first, then pip
if command -v uv &> /dev/null; then
    echo "Using uv for installation..."
    uv pip install sentence-transformers rank-bm25 nltk watchdog transformers torch
else
    echo "Using pip for installation..."
    pip install -r requirements_enhanced.txt
fi

echo ""
echo -e "${GREEN}✅ Dependencies installed${NC}"
echo ""

# Download NLTK data
echo "📚 Downloading NLTK data..."
python -c "import nltk; nltk.download('punkt'); nltk.download('stopwords')" 2>/dev/null
echo -e "${GREEN}✅ NLTK data downloaded${NC}"
echo ""

# Check if we need to migrate
if [ -f "document_hashes.json" ] && [ ! -f "index_status.json" ]; then
    echo -e "${YELLOW}📋 Existing KB detected. Creating migration...${NC}"
    
    # Create index status file
    cat > index_status.json << 'EOF'
{
  "last_full_index": null,
  "last_incremental_index": null,
  "pending_files": [],
  "failed_files": [],
  "index_health": "unknown",
  "auto_index_enabled": true,
  "next_scheduled_index": null,
  "statistics": {
    "total_documents": 0,
    "total_chunks": 0,
    "avg_query_time_ms": 0,
    "last_query_time": null
  }
}
EOF
    echo -e "${GREEN}✅ Migration prepared${NC}"
fi

echo ""
echo "🔍 Testing enhanced KB..."
./kb status

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}✅ Enhanced KB is ready!${NC}"
    echo ""
    echo "📚 Usage Examples:"
    echo ""
    echo "  Search with AI agent context:"
    echo -e "    ${YELLOW}./kb search 'architecture patterns' --agent architect${NC}"
    echo ""
    echo "  Index with better metadata:"
    echo -e "    ${YELLOW}./kb index --path ../docs${NC}"
    echo ""
    echo "  Start auto-indexer:"
    echo -e "    ${YELLOW}python kb_auto_indexer.py${NC}"
    echo ""
    echo "  Check KB health:"
    echo -e "    ${YELLOW}./kb status${NC}"
    echo ""
    echo "💡 New Features:"
    echo "  • Better embedding model (BAAI/bge-large-en-v1.5)"
    echo "  • Hybrid search (BM25 + Vector)"
    echo "  • Query reranking for better precision"
    echo "  • Auto-indexing with file watching"
    echo "  • Enhanced metadata tracking"
    echo "  • Document source categorization"
    echo "  • Agent-specific search optimization"
    echo ""
else
    echo -e "${RED}❌ Setup completed but KB test failed${NC}"
    echo "Please check the error messages above"
fi