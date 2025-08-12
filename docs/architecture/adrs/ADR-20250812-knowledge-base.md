# ADR-20250812: Implement Vector Knowledge Base System

## Status
Accepted

## Context
The project needs a way to intelligently search and understand the codebase, documentation, and architectural decisions. Traditional text search is limited and doesn't understand semantic meaning or relationships between concepts.

## Decision
We will implement a unified vector knowledge base system using ChromaDB and sentence-transformers that:
- Indexes all project files including code, documentation, and configs
- Uses semantic search to find relevant information
- Stores embeddings in a hidden `.vector_db` directory
- Provides a simple CLI interface through a `kb` wrapper script

## Consequences

### Positive
- Semantic search understands meaning, not just keywords
- Architecture assistant can reference past decisions automatically
- Code discovery becomes much more intelligent
- Single tool manages all knowledge base operations
- Local-first approach with no cloud dependencies

### Negative
- Requires ~3GB of dependencies (PyTorch, transformers)
- Initial indexing takes time for large codebases
- Need to re-index when files change significantly

### Risks
- Vector database could grow large for very large codebases
- Embedding model might not understand domain-specific terminology initially

## Alternatives Considered

### Option A: Elasticsearch with Traditional Search
- Full-text search engine with powerful query capabilities
- Pros: Mature, well-documented, powerful search features
- Cons: Requires Java, complex setup, doesn't understand semantic meaning
- Reason for not choosing: Lacks semantic understanding crucial for code comprehension

### Option B: OpenAI Embeddings API
- Cloud-based embedding generation using GPT models
- Pros: High quality embeddings, no local compute requirements
- Cons: Requires API keys, costs per query, data leaves local environment
- Reason for not choosing: Project requires local-first approach with no cloud dependencies

### Option C: Manual Documentation and Tagging
- Manually maintain documentation index and tags
- Pros: Full control, no dependencies, human-curated quality
- Cons: Time-intensive, doesn't scale, prone to becoming outdated
- Reason for not choosing: Not sustainable for active development

## Implementation Notes

### Key Implementation Details
- Uses all-MiniLM-L6-v2 model for balanced speed and quality
- ChromaDB provides persistent local storage
- Virtual environment managed by uv for isolation
- Simple wrapper script for easy CLI access

### Migration Strategy
1. Install dependencies in isolated virtual environment
2. Run initial indexing of entire codebase
3. Test search functionality with common queries
4. Integrate with architect assistant
5. Document usage patterns for team

### Rollback Plan
- Remove `.vector_db` directory to clean all indexed data
- Uninstall virtual environment with `uv`
- Revert to traditional grep/find commands

### Testing Approach
- Unit tests for chunking and parsing logic
- Integration tests for search accuracy
- Performance benchmarks for indexing speed
- User acceptance testing with development team

### Success Metrics
- Search query response time < 2 seconds
- Relevant results in top 5 for 80% of queries
- Indexing speed > 100 files per minute
- Developer adoption rate > 75% within first month

## References
- [ChromaDB Documentation](https://docs.trychroma.com/)
- [Sentence Transformers Documentation](https://www.sbert.net/)
- [Embedding Models Comparison](https://www.sbert.net/docs/pretrained_models.html)
- Related ADRs: None yet (this is the first architecture decision)

## Decision Makers
- Author: [To be filled]
- Reviewers: [To be filled]
- Approver: [To be filled]
- Date: 2025-08-12

## Revision History
- 2025-08-12: Initial draft created
- 2025-08-12: Updated to follow ADR template properly