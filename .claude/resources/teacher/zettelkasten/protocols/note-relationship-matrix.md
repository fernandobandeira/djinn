# Note Relationship Matrix Protocol

## Purpose
Build and maintain a comprehensive matrix of relationships between notes to enable knowledge graph navigation and insight discovery.

## Matrix Structure

### Relationship Types
Define the types of connections between notes.

```yaml
relationship_types:
  builds_on:
    description: "Sequential knowledge where B extends A"
    bidirectional: false
    strength_indicator: "dependency_level"
    visualization: "A → B"
    
  contradicts:
    description: "Opposing views or conflicting information"
    bidirectional: true
    strength_indicator: "conflict_severity"
    visualization: "A ↔ B"
    
  synthesizes:
    description: "Combined insight from multiple notes"
    bidirectional: false
    strength_indicator: "integration_depth"
    visualization: "A + B → C"
    
  analogous:
    description: "Similar pattern in different domain"
    bidirectional: true
    strength_indicator: "similarity_score"
    visualization: "A ≈ B"
    
  generalizes:
    description: "Abstract principle from specific examples"
    bidirectional: false
    strength_indicator: "abstraction_level"
    visualization: "A ⊃ B"
    
  exemplifies:
    description: "Specific instance of general principle"
    bidirectional: false
    strength_indicator: "specificity"
    visualization: "A ⊂ B"
    
  causes:
    description: "Causal relationship"
    bidirectional: false
    strength_indicator: "causality_strength"
    visualization: "A ⇒ B"
    
  refines:
    description: "More precise or nuanced version"
    bidirectional: false
    strength_indicator: "refinement_degree"
    visualization: "A ⟶ B+"
    
  questions:
    description: "Raises questions answered elsewhere"
    bidirectional: false
    strength_indicator: "relevance"
    visualization: "A ? B"
    
  supports:
    description: "Provides evidence or support"
    bidirectional: false
    strength_indicator: "support_strength"
    visualization: "A ⊢ B"
```

### Matrix Implementation

```python
class NoteRelationshipMatrix:
    def __init__(self):
        self.matrix = {}  # Adjacency list representation
        self.nodes = {}   # Note metadata
        self.edges = []   # List of all relationships
        self.clusters = []  # Identified note clusters
        
    def add_note(self, note_id, metadata):
        """Add a note to the matrix"""
        
        self.nodes[note_id] = {
            "title": metadata.get("title"),
            "type": metadata.get("type"),  # permanent, literature, hub
            "created": metadata.get("created"),
            "tags": metadata.get("tags", []),
            "content_hash": self.hash_content(metadata.get("content")),
            "connections": 0,
            "centrality": 0.0
        }
        
        if note_id not in self.matrix:
            self.matrix[note_id] = []
    
    def add_relationship(self, note_a, note_b, rel_type, strength=0.5, metadata=None):
        """Add a relationship between two notes"""
        
        relationship = {
            "from": note_a,
            "to": note_b,
            "type": rel_type,
            "strength": strength,
            "created": datetime.now(),
            "metadata": metadata or {},
            "validated": False
        }
        
        # Add to adjacency list
        if note_a not in self.matrix:
            self.matrix[note_a] = []
        
        self.matrix[note_a].append({
            "target": note_b,
            "relationship": relationship
        })
        
        # Handle bidirectional relationships
        if self.is_bidirectional(rel_type):
            if note_b not in self.matrix:
                self.matrix[note_b] = []
            
            self.matrix[note_b].append({
                "target": note_a,
                "relationship": self.reverse_relationship(relationship)
            })
        
        # Update node connection counts
        self.nodes[note_a]["connections"] += 1
        self.nodes[note_b]["connections"] += 1
        
        # Add to edge list
        self.edges.append(relationship)
        
        # Recalculate centrality
        self.update_centrality()
    
    def find_paths(self, start_note, end_note, max_depth=5):
        """Find all paths between two notes"""
        
        paths = []
        visited = set()
        
        def dfs(current, target, path, depth):
            if depth > max_depth:
                return
            
            if current == target:
                paths.append(list(path))
                return
            
            visited.add(current)
            
            if current in self.matrix:
                for connection in self.matrix[current]:
                    next_note = connection["target"]
                    if next_note not in visited:
                        path.append({
                            "note": next_note,
                            "relationship": connection["relationship"]
                        })
                        dfs(next_note, target, path, depth + 1)
                        path.pop()
            
            visited.remove(current)
        
        dfs(start_note, end_note, [{"note": start_note, "relationship": None}], 0)
        
        return paths
    
    def identify_clusters(self, min_cluster_size=3):
        """Identify strongly connected clusters of notes"""
        
        clusters = []
        visited = set()
        
        for note_id in self.nodes:
            if note_id not in visited:
                cluster = self.explore_cluster(note_id, visited)
                
                if len(cluster) >= min_cluster_size:
                    clusters.append({
                        "notes": cluster,
                        "size": len(cluster),
                        "density": self.calculate_cluster_density(cluster),
                        "central_note": self.find_cluster_center(cluster),
                        "theme": self.identify_cluster_theme(cluster)
                    })
        
        self.clusters = clusters
        return clusters
    
    def calculate_cluster_density(self, cluster_notes):
        """Calculate connection density within cluster"""
        
        if len(cluster_notes) < 2:
            return 0.0
        
        internal_connections = 0
        possible_connections = len(cluster_notes) * (len(cluster_notes) - 1)
        
        for note in cluster_notes:
            if note in self.matrix:
                for connection in self.matrix[note]:
                    if connection["target"] in cluster_notes:
                        internal_connections += 1
        
        return internal_connections / possible_connections if possible_connections > 0 else 0.0
    
    def find_contradictions(self):
        """Find all contradiction relationships for review"""
        
        contradictions = []
        
        for edge in self.edges:
            if edge["type"] == "contradicts":
                contradictions.append({
                    "note_a": edge["from"],
                    "note_b": edge["to"],
                    "strength": edge["strength"],
                    "context": self.get_contradiction_context(edge["from"], edge["to"])
                })
        
        return sorted(contradictions, key=lambda x: x["strength"], reverse=True)
    
    def find_synthesis_opportunities(self):
        """Identify notes that could be synthesized"""
        
        opportunities = []
        
        # Look for notes with high interconnection but no synthesis
        for note_id in self.nodes:
            connections = self.get_connections(note_id)
            
            if len(connections) >= 3:
                # Check if synthesis already exists
                has_synthesis = any(c["relationship"]["type"] == "synthesizes" 
                                   for c in connections)
                
                if not has_synthesis:
                    connected_notes = [c["target"] for c in connections]
                    common_theme = self.find_common_theme(connected_notes)
                    
                    if common_theme:
                        opportunities.append({
                            "base_notes": connected_notes,
                            "theme": common_theme,
                            "potential_insight": self.predict_synthesis(connected_notes),
                            "priority": len(connected_notes) * 0.8
                        })
        
        return sorted(opportunities, key=lambda x: x["priority"], reverse=True)
```

### Relationship Analysis

```python
class RelationshipAnalyzer:
    def __init__(self, matrix):
        self.matrix = matrix
    
    def analyze_relationship_patterns(self):
        """Analyze patterns in relationships"""
        
        patterns = {
            "chains": self.find_knowledge_chains(),
            "cycles": self.find_cycles(),
            "hubs": self.find_hub_notes(),
            "bridges": self.find_bridge_notes(),
            "islands": self.find_isolated_notes()
        }
        
        return patterns
    
    def find_knowledge_chains(self):
        """Find sequential knowledge building chains"""
        
        chains = []
        visited = set()
        
        for note_id in self.matrix.nodes:
            if note_id not in visited:
                chain = self.build_chain(note_id, visited)
                
                if len(chain) >= 3:
                    chains.append({
                        "notes": chain,
                        "length": len(chain),
                        "topic_flow": self.analyze_topic_flow(chain),
                        "complexity_progression": self.analyze_complexity(chain)
                    })
        
        return chains
    
    def build_chain(self, start_note, visited):
        """Build a knowledge chain from a starting note"""
        
        chain = [start_note]
        visited.add(start_note)
        current = start_note
        
        while True:
            # Find 'builds_on' relationships
            next_notes = []
            
            if current in self.matrix.matrix:
                for connection in self.matrix.matrix[current]:
                    if (connection["relationship"]["type"] == "builds_on" and
                        connection["target"] not in visited):
                        next_notes.append(connection["target"])
            
            if not next_notes:
                break
            
            # Choose strongest connection
            next_note = max(next_notes, 
                          key=lambda n: self.get_connection_strength(current, n))
            
            chain.append(next_note)
            visited.add(next_note)
            current = next_note
        
        return chain
    
    def find_hub_notes(self, min_connections=5):
        """Find notes that serve as knowledge hubs"""
        
        hubs = []
        
        for note_id, node_data in self.matrix.nodes.items():
            if node_data["connections"] >= min_connections:
                hubs.append({
                    "note_id": note_id,
                    "title": node_data["title"],
                    "connections": node_data["connections"],
                    "centrality": node_data["centrality"],
                    "connection_types": self.get_connection_type_distribution(note_id),
                    "role": self.determine_hub_role(note_id)
                })
        
        return sorted(hubs, key=lambda x: x["centrality"], reverse=True)
    
    def determine_hub_role(self, note_id):
        """Determine what role a hub note plays"""
        
        connections = self.matrix.get_connections(note_id)
        
        if not connections:
            return "isolated"
        
        type_counts = {}
        for conn in connections:
            rel_type = conn["relationship"]["type"]
            type_counts[rel_type] = type_counts.get(rel_type, 0) + 1
        
        # Determine role based on connection types
        if type_counts.get("generalizes", 0) > len(connections) * 0.4:
            return "abstractor"  # Generalizes many specific concepts
        elif type_counts.get("synthesizes", 0) > len(connections) * 0.3:
            return "synthesizer"  # Combines multiple ideas
        elif type_counts.get("questions", 0) > len(connections) * 0.3:
            return "explorer"  # Raises many questions
        elif type_counts.get("supports", 0) > len(connections) * 0.4:
            return "foundation"  # Supports many other ideas
        else:
            return "connector"  # General connection hub
```

### Matrix Visualization Builder

```python
def build_matrix_visualization(matrix):
    """Build visualization data for the relationship matrix"""
    
    viz_data = {
        "nodes": [],
        "edges": [],
        "clusters": [],
        "statistics": {}
    }
    
    # Node data
    for note_id, node_data in matrix.nodes.items():
        viz_data["nodes"].append({
            "id": note_id,
            "label": node_data["title"],
            "type": node_data["type"],
            "size": max(5, node_data["connections"] * 2),  # Size by connections
            "color": get_node_color(node_data["type"]),
            "x": None,  # Will be calculated by layout algorithm
            "y": None
        })
    
    # Edge data
    for edge in matrix.edges:
        viz_data["edges"].append({
            "source": edge["from"],
            "target": edge["to"],
            "type": edge["type"],
            "weight": edge["strength"],
            "color": get_edge_color(edge["type"]),
            "style": get_edge_style(edge["type"])
        })
    
    # Cluster data
    for cluster in matrix.clusters:
        viz_data["clusters"].append({
            "nodes": cluster["notes"],
            "label": cluster["theme"],
            "color": generate_cluster_color(cluster)
        })
    
    # Statistics
    viz_data["statistics"] = {
        "total_notes": len(matrix.nodes),
        "total_relationships": len(matrix.edges),
        "avg_connections": sum(n["connections"] for n in matrix.nodes.values()) / len(matrix.nodes),
        "cluster_count": len(matrix.clusters),
        "relationship_distribution": get_relationship_distribution(matrix)
    }
    
    return viz_data
```

### Matrix Report Template

```markdown
# Note Relationship Matrix Report

## Overview
- Total Notes: [X]
- Total Relationships: [Y]
- Average Connections per Note: [Z]
- Matrix Density: [D]%

## Relationship Distribution
| Type | Count | Percentage |
|------|-------|------------|
| builds_on | X | Y% |
| contradicts | X | Y% |
| synthesizes | X | Y% |
| analogous | X | Y% |
| generalizes | X | Y% |

## Key Insights

### Knowledge Hubs (Top 5)
1. **[Note Title]** - [X connections]
   - Role: [Hub Role]
   - Key connections: [List]

### Knowledge Chains
1. [Start] → [Middle] → [...] → [End]
   - Length: [X notes]
   - Topic progression: [Description]

### Synthesis Opportunities
1. Combine: [Note A], [Note B], [Note C]
   - Common theme: [Theme]
   - Potential insight: [Description]

### Contradictions to Resolve
1. [Note A] ↔ [Note B]
   - Conflict: [Description]
   - Resolution path: [Suggestion]

### Isolated Notes
- [List of notes with < 2 connections]

## Cluster Analysis

### Major Clusters
1. **[Cluster Theme]** ([X] notes)
   - Central note: [Title]
   - Density: [Y]%
   - Key concepts: [List]

## Recommendations
1. Connect isolated notes to relevant clusters
2. Resolve identified contradictions
3. Create synthesis notes for high-opportunity areas
4. Strengthen weak connections with additional links
```

## Benefits

1. **Knowledge Navigation**: Easy traversal of knowledge graph
2. **Pattern Discovery**: Identify knowledge building patterns
3. **Gap Detection**: Find missing connections
4. **Synthesis Opportunities**: Identify combinable insights
5. **Contradiction Management**: Track and resolve conflicts
6. **Cluster Analysis**: Understand knowledge domains