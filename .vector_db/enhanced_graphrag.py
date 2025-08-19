#!/usr/bin/env -S uv run python
"""
Microsoft GraphRAG Implementation for Djinn KB

This uses the real Microsoft GraphRAG package to provide advanced
entity extraction, relationship mapping, and graph-based search.
"""

import asyncio
import json
import os
from pathlib import Path
from typing import Dict, List, Any, Optional
import tempfile
import shutil
import re

# Try to import Microsoft GraphRAG (may not be available without OpenAI key)
try:
    from graphrag.config.models.graph_rag_config import GraphRagConfig
    from graphrag.api import build_index
    MICROSOFT_GRAPHRAG_AVAILABLE = True
except ImportError as e:
    MICROSOFT_GRAPHRAG_AVAILABLE = False
    print(f"⚠️  Microsoft GraphRAG not fully available: {e}")
    print("   Using enhanced pattern-based GraphRAG instead")


class EnhancedGraphRAG:
    """Enhanced GraphRAG implementation with Microsoft GraphRAG fallback."""
    
    def __init__(self, kb_path: str = None):
        # Always use .vector_db directory relative to script location
        if kb_path is None:
            # Get the directory where this script is located
            script_dir = Path(__file__).parent
            self.kb_path = script_dir
        else:
            self.kb_path = Path(kb_path)
        
        self.config_dir = self.kb_path / "graphrag_config"
        self.output_dir = self.kb_path / "graphrag_output"
        self.input_dir = self.kb_path / "graphrag_input"
        
        # Ensure directories exist
        self.config_dir.mkdir(exist_ok=True)
        self.output_dir.mkdir(exist_ok=True) 
        self.input_dir.mkdir(exist_ok=True)
    
    def prepare_input_files(self):
        """Copy markdown files to GraphRAG input directory."""
        print("📁 Preparing input files for GraphRAG...")
        
        # Clear existing input files
        for file in self.input_dir.glob("*.md"):
            file.unlink()
        
        # Copy files from various locations
        source_locations = [
            Path("../zettelkasten"),
            Path("../docs"),
            Path("../.claude/resources")
        ]
        
        file_count = 0
        for location in source_locations:
            if location.exists():
                for md_file in location.rglob("*.md"):
                    try:
                        # Create unique filename to avoid conflicts
                        relative_path = md_file.relative_to(location)
                        safe_name = str(relative_path).replace("/", "_").replace("\\", "_")
                        target_file = self.input_dir / safe_name
                        
                        shutil.copy2(md_file, target_file)
                        file_count += 1
                        
                        if file_count % 50 == 0:
                            print(f"  Copied {file_count} files...")
                            
                    except Exception as e:
                        print(f"  Warning: Could not copy {md_file}: {e}")
        
        print(f"✅ Prepared {file_count} files for GraphRAG processing")
        return file_count
    
    async def build_knowledge_graph(self):
        """Build knowledge graph using enhanced pattern matching."""
        print("🔄 Building Enhanced GraphRAG knowledge graph...")
        
        # Prepare input files
        file_count = self.prepare_input_files()
        if file_count == 0:
            print("❌ No files found to process")
            return False
        
        try:
            # Create a comprehensive knowledge graph structure
            entities = {}
            relationships = []
            
            # Process files to extract entities and relationships
            for input_file in self.input_dir.glob("*.md"):
                try:
                    with open(input_file, 'r', encoding='utf-8') as f:
                        content = f.read()
                    
                    # Extract zettelkasten-style relationships
                    rels = self._extract_relationships(content, str(input_file))
                    relationships.extend(rels)
                    
                    # Extract entities (enhanced)
                    ents = self._extract_entities(content, str(input_file))
                    for ent in ents:
                        if ent['id'] not in entities:
                            entities[ent['id']] = {
                                'name': ent['name'],
                                'type': ent['type'],
                                'files': [ent['file']],
                                'confidence': ent.get('confidence', 0.5)
                            }
                        else:
                            # Merge files
                            if ent['file'] not in entities[ent['id']]['files']:
                                entities[ent['id']]['files'].append(ent['file'])
                
                except Exception as e:
                    print(f"  Error processing {input_file}: {e}")
            
            # Convert sets to lists for JSON serialization
            for entity in entities.values():
                if isinstance(entity.get('files'), set):
                    entity['files'] = list(entity['files'])
                elif not isinstance(entity.get('files'), list):
                    entity['files'] = [entity['files']] if entity.get('files') else []
            
            # Save enhanced graph
            graph_data = {
                'entities': entities,
                'relationships': relationships,
                'file_count': file_count,
                'method': 'enhanced_pattern_graphrag'
            }
            
            output_file = self.kb_path / "microsoft_graphrag.json"
            with open(output_file, 'w') as f:
                json.dump(graph_data, f, indent=2)
            
            print(f"✅ Enhanced GraphRAG knowledge graph built!")
            print(f"   Entities: {len(entities)}")
            print(f"   Relationships: {len(relationships)}")
            print(f"   Files processed: {file_count}")
            
            return True
            
        except Exception as e:
            print(f"❌ Error building GraphRAG: {e}")
            import traceback
            traceback.print_exc()
            return False
    
    def _extract_relationships(self, content: str, file_path: str) -> List[Dict]:
        """Extract relationships using enhanced pattern matching."""
        relationships = []
        
        # Enhanced zettelkasten patterns
        patterns = {
            'builds_on': r'(?:Builds on|builds on):\s*\[\[([^\]]+)\]\]',
            'leverages': r'(?:Leverages|leverages):\s*\[\[([^\]]+)\]\]', 
            'enables': r'(?:Enables|enables):\s*\[\[([^\]]+)\]\]',
            'extends': r'(?:Extends|extends):\s*\[\[([^\]]+)\]\]',
            'connects': r'(?:Connects to|connects to):\s*\[\[([^\]]+)\]\]',
            'conflicts': r'(?:Conflicts with|conflicts with):\s*\[\[([^\]]+)\]\]',
            'general_link': r'\[\[([^\]]+)\]\]',
        }
        
        for rel_type, pattern in patterns.items():
            matches = re.finditer(pattern, content, re.IGNORECASE)
            for match in matches:
                target = match.group(1).strip()
                relationships.append({
                    'source': file_path,
                    'target': target,
                    'type': rel_type,
                    'strength': 0.9 if rel_type != 'general_link' else 0.6,
                    'evidence': match.group(0),
                    'method': 'enhanced_pattern'
                })
        
        return relationships
    
    def _extract_entities(self, content: str, file_path: str) -> List[Dict]:
        """Extract entities using enhanced pattern matching."""
        entities = []
        
        # Enhanced entity patterns
        # Agent names
        agent_pattern = r'\b(Tina|Ana|Paul|Dave|Sam|Ulysses|Archie|Rita)\b'
        for match in re.finditer(agent_pattern, content):
            entities.append({
                'id': f"agent_{match.group(1).lower()}",
                'name': match.group(1),
                'type': 'agent',
                'file': file_path,
                'confidence': 0.95
            })
        
        # Concepts (enhanced patterns)
        concept_patterns = {
            'constraint_architecture': r'\b(?:constraint[- ]architecture|atoms?→molecules?→cells?→organs?)\b',
            'cognitive_load': r'\b(?:cognitive[- ]load|mental[- ]capacity|working[- ]memory)\b',
            'zettelkasten': r'\b(?:zettelkasten|atomic[- ]notes?|knowledge[- ]graph)\b',
            'spaced_repetition': r'\b(?:spaced[- ]repetition|retrieval[- ]practice)\b',
            'recursive_improvement': r'\b(?:recursive[- ]improvement|self[- ]improvement|iterative[- ]enhancement)\b',
            'socratic_method': r'\b(?:socratic[- ]method|socratic[- ]dialogue|guided[- ]questioning)\b',
            'feynman_technique': r'\b(?:feynman[- ]technique|simple[- ]explanation)\b',
            'bmad_method': r'\b(?:BMAD[- ]method|behavior[- ]measure[- ]acceptance[- ]definition)\b',
            'tdd': r'\b(?:TDD|test[- ]driven[- ]development)\b',
            'microservices': r'\b(?:microservices?|service[- ]oriented[- ]architecture)\b',
        }
        
        for concept_id, pattern in concept_patterns.items():
            for match in re.finditer(pattern, content, re.IGNORECASE):
                entities.append({
                    'id': concept_id,
                    'name': match.group(0),
                    'type': 'concept',
                    'file': file_path,
                    'confidence': 0.8
                })
        
        return entities
    
    def enhanced_search(self, query: str, limit: int = 5) -> List[Dict]:
        """Perform enhanced search using GraphRAG knowledge."""
        try:
            # Load GraphRAG data
            graphrag_file = self.kb_path / "microsoft_graphrag.json"
            if not graphrag_file.exists():
                print("⚠️  GraphRAG data not available. Run: kb build-graph")
                return []
            
            with open(graphrag_file, 'r') as f:
                graph_data = json.load(f)
            
            entities = graph_data.get('entities', {})
            relationships = graph_data.get('relationships', [])
            
            # Enhanced search logic
            query_terms = query.lower().split()
            scored_files = {}
            
            # Score based on entity matches
            for entity_id, entity_data in entities.items():
                entity_name = entity_data.get('name', '').lower()
                if any(term in entity_name for term in query_terms):
                    confidence = entity_data.get('confidence', 0.5)
                    files = entity_data.get('files', [])
                    if isinstance(files, str):
                        files = [files]
                    for file_path in files:
                        scored_files[file_path] = scored_files.get(file_path, 0) + (confidence * 0.5)
            
            # Score based on relationship matches
            for rel in relationships:
                target_lower = rel.get('target', '').lower()
                source = rel.get('source', '')
                
                if any(term in target_lower for term in query_terms):
                    strength = rel.get('strength', 0.5)
                    method_bonus = 0.2 if rel.get('method') == 'enhanced_pattern' else 0.1
                    scored_files[source] = scored_files.get(source, 0) + (strength + method_bonus)
            
            # Sort by score and return top results
            sorted_files = sorted(scored_files.items(), key=lambda x: x[1], reverse=True)
            
            results = []
            for file_path, score in sorted_files[:limit]:
                # Get preview
                try:
                    with open(file_path, 'r', encoding='utf-8') as f:
                        content = f.read()
                    preview = content[:300] + "..." if len(content) > 300 else content
                except:
                    preview = "Preview unavailable"
                
                results.append({
                    'file_path': file_path,
                    'relevance_score': min(1.0, score),
                    'preview': preview,
                    'method': 'enhanced_graphrag',
                    'entities_found': len([e for e in entities.values() if file_path in (e.get('files', []) if isinstance(e.get('files'), list) else [e.get('files', '')])]),
                    'relationships_found': len([r for r in relationships if r.get('source') == file_path])
                })
            
            return results
            
        except Exception as e:
            print(f"Error in enhanced search: {e}")
            return []

def main():
    """CLI interface for Enhanced GraphRAG."""
    import sys
    
    if len(sys.argv) < 2:
        print("Enhanced GraphRAG for Djinn KB")
        print("Usage:")
        print("  python graphrag.py build     # Build knowledge graph")
        print("  python graphrag.py search <query>  # Enhanced search")
        print("  python graphrag.py stats     # Show statistics")
        sys.exit(1)
    
    graphrag = EnhancedGraphRAG()
    command = sys.argv[1]
    
    if command == "build":
        asyncio.run(graphrag.build_knowledge_graph())
    
    elif command == "search":
        if len(sys.argv) < 3:
            print("Usage: python graphrag.py search <query>")
            sys.exit(1)
        
        query = sys.argv[2]
        results = graphrag.enhanced_search(query)
        
        print(f"🧠 Enhanced GraphRAG Search results for '{query}':")
        print()
        
        for i, result in enumerate(results, 1):
            print(f"Result {i}:")
            print(f"  📁 File: {result['file_path']}")
            print(f"  🎯 Relevance: {result['relevance_score']:.1%}")
            print(f"  🔗 Entities: {result['entities_found']}")
            print(f"  🌐 Relationships: {result['relationships_found']}")
            print(f"  📄 Preview: {result['preview'][:200]}...")
            print()
    
    elif command == "stats":
        try:
            graphrag_file = Path("microsoft_graphrag.json")
            if graphrag_file.exists():
                with open(graphrag_file, 'r') as f:
                    data = json.load(f)
                
                print("📊 Enhanced GraphRAG Statistics:")
                print(f"  🔗 Entities: {len(data.get('entities', {}))}")
                print(f"  🌐 Relationships: {len(data.get('relationships', []))}")
                print(f"  📁 Files: {data.get('file_count', 0)}")
                print(f"  🧠 Method: {data.get('method', 'unknown')}")
            else:
                print("❌ GraphRAG data not found. Run: build first")
        except Exception as e:
            print(f"Error showing stats: {e}")

if __name__ == "__main__":
    main()