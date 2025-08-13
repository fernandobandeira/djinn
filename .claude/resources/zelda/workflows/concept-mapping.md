# Concept Mapping Workflow

## Overview
Transform your Zettelkasten notes into visual knowledge graphs that reveal patterns, connections, and learning paths.

## Types of Concept Maps

### 1. Domain Overview Map
Shows all major concepts in a knowledge area:
```mermaid
graph TB
    Domain[Machine Learning]
    
    subgraph "Fundamentals"
        Math[Mathematics]
        Stats[Statistics]
        Prog[Programming]
    end
    
    subgraph "Core Concepts"
        SL[Supervised Learning]
        UL[Unsupervised Learning]
        RL[Reinforcement Learning]
    end
    
    subgraph "Applications"
        CV[Computer Vision]
        NLP[Natural Language]
        Rob[Robotics]
    end
    
    Domain --> Math
    Domain --> Stats
    Domain --> Prog
    Math --> SL
    Stats --> SL
    SL --> CV
    SL --> NLP
    UL --> CV
    RL --> Rob
```

### 2. Learning Path Map
Shows progression through concepts:
```mermaid
graph LR
    Start[Basic Python] --> DS[Data Structures]
    DS --> Algo[Algorithms]
    Algo --> ML1[ML Basics]
    ML1 --> ML2[ML Intermediate]
    ML2 --> DL[Deep Learning]
    
    DS -.-> Practice1[Project 1]
    ML1 -.-> Practice2[Project 2]
    DL -.-> Practice3[Capstone]
    
    style Start fill:#90EE90
    style Practice1 fill:#FFB6C1
    style Practice2 fill:#FFB6C1
    style Practice3 fill:#FFB6C1
```

### 3. Connection Network Map
Shows how a single concept connects:
```mermaid
graph TD
    Central[Backpropagation]
    
    Central --> Gradient[Gradient Descent]
    Central --> Chain[Chain Rule]
    Central --> Neural[Neural Networks]
    Central --> Optimize[Optimization]
    
    Gradient --> SGD[SGD]
    Gradient --> Adam[Adam]
    Chain --> Calculus[Calculus]
    Neural --> CNN[CNN]
    Neural --> RNN[RNN]
    Optimize --> Loss[Loss Functions]
    
    style Central fill:#FFD700
```

### 4. Synthesis Map
Shows how ideas combine:
```mermaid
graph BT
    Idea1[Attention Mechanism]
    Idea2[Sequence Models]
    Idea3[Self-Supervision]
    
    Synthesis[Transformers]
    
    Idea1 --> Synthesis
    Idea2 --> Synthesis
    Idea3 --> Synthesis
    
    Synthesis --> App1[BERT]
    Synthesis --> App2[GPT]
    Synthesis --> App3[Vision Transformers]
    
    style Synthesis fill:#9370DB
```

## Creating Concept Maps

### Step 1: Gather Related Notes
```bash
# Find all notes with specific tag
grep -l "#machine-learning" zettelkasten/**/*.md

# Find notes with specific connections
grep -l "\[\[central-note-id\]\]" zettelkasten/**/*.md

# Find notes in date range
find zettelkasten -name "*.md" -mtime -30
```

### Step 2: Identify Structure Type

#### Hierarchical Structure
Best for:
- Learning prerequisites
- Taxonomies
- Organizational charts
- Decision trees

#### Network Structure
Best for:
- Interconnected concepts
- Influence relationships
- System interactions
- Knowledge webs

#### Flow Structure
Best for:
- Processes
- Learning paths
- Cause and effect
- Temporal sequences

#### Radial Structure
Best for:
- Central concept exploration
- Mind mapping
- Feature analysis
- Characteristic breakdown

### Step 3: Extract Key Elements

#### Nodes (Concepts)
- Note titles
- Core insights
- Key terms
- Important questions

#### Edges (Relationships)
- "Builds on"
- "Contrasts with"
- "Enables"
- "Is example of"
- "Requires"
- "Leads to"

#### Clusters (Groups)
- By domain
- By complexity level
- By time period
- By application area

### Step 4: Generate Mermaid Diagram

#### Basic Template
```mermaid
graph TD
    %% Define nodes
    Node1[Concept 1]
    Node2[Concept 2]
    Node3[Concept 3]
    
    %% Define relationships
    Node1 --> Node2
    Node1 -.-> Node3
    Node2 ==> Node3
    
    %% Style important nodes
    style Node1 fill:#f9f,stroke:#333,stroke-width:4px
    
    %% Add subgraphs for grouping
    subgraph "Group Name"
        Node4[Concept 4]
        Node5[Concept 5]
    end
```

#### Relationship Types
- `-->` Solid arrow (strong connection)
- `-.->` Dotted arrow (weak connection)
- `==>` Thick arrow (critical path)
- `---` Line without arrow (association)

### Step 5: Enhance with Metadata

#### Node Styling by Maturity
```mermaid
graph TD
    Seedling[New Concept]
    Growing[Developing Concept]
    Mature[Established Concept]
    
    style Seedling fill:#90EE90
    style Growing fill:#FFD700
    style Mature fill:#4169E1
```

#### Edge Labels
```mermaid
graph TD
    A[Concept A] -->|enables| B[Concept B]
    B -->|requires| C[Concept C]
    C -->|contradicts| D[Concept D]
```

## Automated Map Generation

### Python Script for Note Network
```python
#!/usr/bin/env python3
import os
import re
from collections import defaultdict

def generate_concept_map(directory):
    """Generate Mermaid diagram from note connections"""
    
    connections = defaultdict(list)
    notes = {}
    
    # Parse all notes
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith('.md'):
                note_id = file.replace('.md', '')
                filepath = os.path.join(root, file)
                
                with open(filepath, 'r') as f:
                    content = f.read()
                    
                    # Extract title
                    title_match = re.search(r'^# (.+)$', content, re.M)
                    if title_match:
                        notes[note_id] = title_match.group(1)
                    
                    # Extract connections
                    links = re.findall(r'\[\[([^\]]+)\]\]', content)
                    for link in links:
                        connections[note_id].append(link)
    
    # Generate Mermaid
    mermaid = ["graph TD"]
    
    # Add nodes
    for note_id, title in notes.items():
        safe_id = note_id.replace('-', '_')
        mermaid.append(f"    {safe_id}[{title}]")
    
    # Add connections
    for source, targets in connections.items():
        safe_source = source.replace('-', '_')
        for target in targets:
            if target in notes:
                safe_target = target.replace('-', '_')
                mermaid.append(f"    {safe_source} --> {safe_target}")
    
    return '\n'.join(mermaid)

# Generate map
map_code = generate_concept_map('zettelkasten/permanent')
print(map_code)
```

### Interactive D3.js Visualization
```html
<!DOCTYPE html>
<html>
<head>
    <script src="https://d3js.org/d3.v7.min.js"></script>
    <style>
        .node {
            fill: #69b3a2;
            stroke: #000;
            stroke-width: 1.5px;
        }
        .link {
            stroke: #999;
            stroke-opacity: 0.6;
        }
        .label {
            font: 12px sans-serif;
            pointer-events: none;
        }
    </style>
</head>
<body>
    <svg width="960" height="600"></svg>
    <script>
        // Load note data
        d3.json("notes.json").then(function(data) {
            // Create force simulation
            const simulation = d3.forceSimulation(data.nodes)
                .force("link", d3.forceLink(data.links).id(d => d.id))
                .force("charge", d3.forceManyBody().strength(-300))
                .force("center", d3.forceCenter(480, 300));
            
            // Create SVG elements
            const svg = d3.select("svg");
            
            const link = svg.selectAll(".link")
                .data(data.links)
                .enter().append("line")
                .attr("class", "link");
            
            const node = svg.selectAll(".node")
                .data(data.nodes)
                .enter().append("circle")
                .attr("class", "node")
                .attr("r", 5)
                .call(d3.drag()
                    .on("start", dragstarted)
                    .on("drag", dragged)
                    .on("end", dragended));
            
            // Add labels
            const label = svg.selectAll(".label")
                .data(data.nodes)
                .enter().append("text")
                .attr("class", "label")
                .text(d => d.title);
            
            // Update positions
            simulation.on("tick", () => {
                link
                    .attr("x1", d => d.source.x)
                    .attr("y1", d => d.source.y)
                    .attr("x2", d => d.target.x)
                    .attr("y2", d => d.target.y);
                
                node
                    .attr("cx", d => d.x)
                    .attr("cy", d => d.y);
                
                label
                    .attr("x", d => d.x + 8)
                    .attr("y", d => d.y + 3);
            });
        });
    </script>
</body>
</html>
```

## Map Analysis Techniques

### Pattern Recognition
Look for:
- Dense clusters (key topic areas)
- Bridge nodes (connecting different domains)
- Isolated nodes (needs integration)
- Long paths (complex dependencies)
- Cycles (reinforcing concepts)

### Gap Analysis
Identify:
- Missing connections
- Underdeveloped areas
- Potential synthesis points
- Learning prerequisites not documented

### Path Finding
Discover:
- Shortest learning paths
- Alternative routes to understanding
- Critical concepts (many paths pass through)
- Optional vs required knowledge

## Export Formats

### For Documentation
```markdown
## Knowledge Map

![Concept Map](concept-map.png)

### Key Nodes
- **Central Concepts**: [List]
- **Prerequisites**: [List]
- **Applications**: [List]

### Learning Paths
1. Beginner: A → B → C
2. Intermediate: D → E → F
3. Advanced: G → H → I
```

### For Presentation
```mermaid
%%{init: {'theme':'base', 'themeVariables': { 'primaryColor':'#ff0000'}}}%%
graph LR
    Start([Start Here])
    End([Mastery])
    
    Start ==> Foundation
    Foundation ==> Intermediate
    Intermediate ==> Advanced
    Advanced ==> End
    
    style Start fill:#90EE90
    style End fill:#FFD700
```

### For Analysis
```csv
source,target,relationship,strength
note1,note2,builds_on,strong
note2,note3,enables,medium
note3,note4,contrasts,weak
```

## Integration with Learning

### Before Learning Session
1. Generate map of target area
2. Identify prerequisites
3. Plan learning path
4. Note current position

### During Learning
1. Add new nodes as discovered
2. Mark traversed paths
3. Note unexpected connections
4. Identify missing links

### After Learning
1. Update map with new knowledge
2. Reorganize if patterns emerge
3. Highlight insights
4. Plan next exploration

## Remember
- Maps are thinking tools, not just displays
- Keep maps focused (max 20-30 nodes)
- Use consistent relationship types
- Update maps as understanding evolves
- Different views reveal different insights