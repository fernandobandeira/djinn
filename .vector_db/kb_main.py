#!/usr/bin/env -S uv run python
"""
Primary KB Command - GraphRAG-Enhanced by Default

This makes GraphRAG the primary search method while keeping vector search as fallback.
"""

import sys
import subprocess
import json
from pathlib import Path

def load_microsoft_graphrag():
    """Load Microsoft GraphRAG data."""
    graphrag_file = Path(__file__).parent / "microsoft_graphrag.json"
    if graphrag_file.exists():
        try:
            with open(graphrag_file, 'r') as f:
                return json.load(f)
        except:
            pass
    return None

def find_related_files(query_terms, relationships):
    """Find files related to query through relationships."""
    if not relationships:
        return []
    
    related_files = set()
    query_lower = [term.lower() for term in query_terms]
    
    for rel in relationships.get('relationships', []):
        target_lower = rel['target'].lower()
        source = rel['source']
        
        if any(term in target_lower for term in query_lower):
            related_files.add(source)
        
        source_lower = source.lower()
        if any(term in source_lower for term in query_lower):
            related_files.add(source)
    
    return list(related_files)

def search(query, collection=None, limit=5, force_vector=False):
    """Primary search - Microsoft GraphRAG enhanced by default."""
    
    # If not forcing vector-only, try Microsoft GraphRAG first
    if not force_vector:
        graphrag_script = Path(__file__).parent / "enhanced_graphrag.py"
        
        if graphrag_script.exists():
            try:
                result = subprocess.run(["uv", "run", "python", str(graphrag_script), "search", query], 
                                      capture_output=True, text=True)
                if result.returncode == 0:
                    print(result.stdout)
                    return
                else:
                    print("⚠️ Enhanced GraphRAG search failed, falling back to vector search")
            except Exception as e:
                print(f"⚠️ Enhanced GraphRAG error: {e}, falling back to vector search")
    
    # Fallback to vector search
    kb_script = ["uv", "run", "python", str(Path(__file__).parent / "kb.py")]
    
    # Build KB command
    cmd = kb_script + ["search", query]
    if collection:
        cmd.extend(["--collection", collection])
    cmd.extend(["--limit", str(limit)])
    
    try:
        # Get vector search results
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            print(f"Error: {result.stderr}")
            return
        
        search_type = "Vector" if force_vector else "Fallback Vector"
        print(f"🔍 {search_type} Search results for '{query}':")
        if not force_vector:
            print("💡 Build Microsoft GraphRAG with: kb build-graph")
        print()
        print(result.stdout)
        
    except Exception as e:
        print(f"Search error: {e}")

def build_graph():
    """Build the Microsoft GraphRAG knowledge graph."""
    graphrag_script = Path(__file__).parent / "enhanced_graphrag.py"
    
    if not graphrag_script.exists():
        print("❌ Microsoft GraphRAG script not found")
        return
    
    try:
        result = subprocess.run(["uv", "run", "python", str(graphrag_script), "build"], 
                              capture_output=True, text=True)
        if result.returncode == 0:
            print("✅ Enhanced GraphRAG knowledge graph built successfully!")
            print(result.stdout)
        else:
            print(f"❌ Error building Enhanced GraphRAG: {result.stderr}")
    except Exception as e:
        print(f"Build error: {e}")

def show_stats():
    """Show comprehensive statistics."""
    kb_script = ["uv", "run", "python", str(Path(__file__).parent / "kb.py")]
    
    # Vector search stats
    try:
        result = subprocess.run(kb_script + ["stats"], capture_output=True, text=True)
        if result.returncode == 0:
            print("📊 Vector Search Statistics:")
            for line in result.stdout.split('\n'):
                if line.strip():
                    print(f"  {line}")
            print()
    except:
        print("Vector search statistics unavailable")
    
    # Microsoft GraphRAG stats
    graphrag_data = load_microsoft_graphrag()
    if graphrag_data:
        print("🧠 Microsoft GraphRAG Enhancement:")
        print(f"  🔗 Entities: {len(graphrag_data.get('entities', {}))}")
        print(f"  🌐 Relationships: {len(graphrag_data.get('relationships', []))}")
        print(f"  📁 Files: {graphrag_data.get('file_count', 0)}")
        print(f"  🧠 Method: {graphrag_data.get('method', 'unknown')}")
        
        # Entity breakdown
        entity_types = {}
        for entity_data in graphrag_data.get('entities', {}).values():
            entity_type = entity_data.get('type', 'unknown')
            entity_types[entity_type] = entity_types.get(entity_type, 0) + 1
        
        if entity_types:
            print("  📊 Entity types:")
            for entity_type, count in entity_types.items():
                print(f"    {entity_type}: {count}")
        
        print("\n✨ Microsoft GraphRAG is active - all searches are AI-enhanced!")
    else:
        print("🔗 Microsoft GraphRAG: Not available")
        print("💡 Build with: kb build-graph")

def main():
    """Main interface - GraphRAG by default."""
    if len(sys.argv) < 2:
        print("KB - Microsoft GraphRAG Enhanced Knowledge Base")
        print()
        print("Usage:")
        print("  kb search <query> [--collection <collection>] [--vector-only]")
        print("  kb build-graph          # Build Microsoft GraphRAG knowledge graph")
        print("  kb stats                # Show enhanced statistics")
        print("  kb index               # Pass through to vector indexing")
        print("  kb clear               # Pass through to vector clearing")
        print()
        print("Examples:")
        print("  kb search 'constraint architecture'    # Microsoft GraphRAG enhanced (default)")
        print("  kb search 'text' --vector-only        # Force vector-only search")
        print("  kb build-graph                         # Enable Microsoft GraphRAG")
        print()
        print("🧠 Microsoft GraphRAG provides AI-powered entity extraction and relationship mapping!")
        sys.exit(1)
    
    command = sys.argv[1]
    
    if command == "search":
        if len(sys.argv) < 3:
            print("Usage: kb search <query> [--collection <collection>] [--vector-only]")
            sys.exit(1)
        
        query = sys.argv[2]
        collection = None
        force_vector = False
        
        # Parse arguments
        i = 3
        while i < len(sys.argv):
            if sys.argv[i] == "--collection" and i + 1 < len(sys.argv):
                collection = sys.argv[i + 1]
                i += 2
            elif sys.argv[i] == "--vector-only":
                force_vector = True
                i += 1
            else:
                i += 1
        
        search(query, collection, 5, force_vector)
    
    elif command == "build-graph":
        build_graph()
    
    elif command == "stats":
        show_stats()
    
    elif command in ["index", "clear"]:
        # Pass through to original KB
        kb_script = ["uv", "run", "python", str(Path(__file__).parent / "kb.py")]
        try:
            subprocess.run(kb_script + sys.argv[1:])
        except Exception as e:
            print(f"Error: {e}")
    
    else:
        print(f"Unknown command: {command}")
        print("Use: search, build-graph, stats, index, or clear")

if __name__ == "__main__":
    main()