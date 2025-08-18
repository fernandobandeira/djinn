# Concept Extractor - Cognitive Tool for Zettelkasten

## Purpose
Systematically extract, structure, and score concepts from various sources for atomic note creation.

## Extraction Pipeline

### Phase 1: Identify Concepts
Scan sources for key terms, ideas, and insights.

```python
class ConceptExtractor:
    def __init__(self):
        self.extraction_patterns = {
            "definitions": r"(\w+) is defined as|(\w+) means|(\w+) refers to",
            "properties": r"characteristics of|properties include|features are",
            "relationships": r"relates to|connected with|depends on|causes|leads to",
            "examples": r"for example|such as|including|like",
            "distinctions": r"unlike|different from|in contrast to|whereas",
            "principles": r"principle of|law of|rule of|theorem",
            "processes": r"steps include|process involves|procedure for",
            "insights": r"key insight|importantly|notably|significantly"
        }
        
        self.concept_registry = {}
        self.relevance_threshold = 0.5
    
    def extract_concepts(self, source_text, metadata=None):
        """Main extraction method"""
        
        # Preprocessing
        cleaned_text = self.preprocess(source_text)
        
        # Multi-strategy extraction
        concepts = []
        concepts.extend(self.pattern_based_extraction(cleaned_text))
        concepts.extend(self.entity_extraction(cleaned_text))
        concepts.extend(self.key_phrase_extraction(cleaned_text))
        concepts.extend(self.semantic_extraction(cleaned_text))
        
        # Deduplicate and merge
        unified_concepts = self.unify_concepts(concepts)
        
        # Score and rank
        scored_concepts = self.score_concepts(unified_concepts, metadata)
        
        # Filter by relevance
        relevant_concepts = [c for c in scored_concepts if c["relevance"] >= self.relevance_threshold]
        
        return self.structure_concepts(relevant_concepts)
    
    def pattern_based_extraction(self, text):
        """Extract using predefined patterns"""
        
        concepts = []
        
        for pattern_type, pattern in self.extraction_patterns.items():
            matches = re.finditer(pattern, text, re.IGNORECASE)
            
            for match in matches:
                concept = {
                    "text": match.group(0),
                    "type": pattern_type,
                    "context": text[max(0, match.start()-50):min(len(text), match.end()+50)],
                    "position": match.start(),
                    "confidence": 0.8  # Pattern matches are usually reliable
                }
                concepts.append(concept)
        
        return concepts
    
    def entity_extraction(self, text):
        """Extract named entities and important terms"""
        
        concepts = []
        
        # Identify capitalized terms (potential important concepts)
        capitalized = re.findall(r'\b[A-Z][a-z]+(?:\s+[A-Z][a-z]+)*\b', text)
        
        for term in capitalized:
            if len(term) > 3 and term not in self.common_words:
                concepts.append({
                    "text": term,
                    "type": "entity",
                    "context": self.get_context(text, term),
                    "confidence": 0.7
                })
        
        # Identify technical terms (containing numbers, symbols, or specific patterns)
        technical = re.findall(r'\b\w+[-_]\w+\b|\b\w+\d+\w*\b', text)
        
        for term in technical:
            concepts.append({
                "text": term,
                "type": "technical_term",
                "context": self.get_context(text, term),
                "confidence": 0.75
            })
        
        return concepts
    
    def key_phrase_extraction(self, text):
        """Extract important phrases"""
        
        concepts = []
        
        # Look for emphasized phrases
        emphasized = re.findall(r'\"([^\"]+)\"|\'([^\']+)\'|\*([^\*]+)\*|_([^_]+)_', text)
        
        for groups in emphasized:
            phrase = next(g for g in groups if g)  # Get non-empty group
            if len(phrase.split()) <= 5:  # Reasonable phrase length
                concepts.append({
                    "text": phrase,
                    "type": "key_phrase",
                    "context": self.get_context(text, phrase),
                    "confidence": 0.85  # Emphasized text is usually important
                })
        
        # Look for phrases with signal words
        signal_phrases = [
            r'the concept of ([\w\s]+)',
            r'the idea that ([\w\s]+)',
            r'the principle of ([\w\s]+)',
            r'the theory of ([\w\s]+)'
        ]
        
        for pattern in signal_phrases:
            matches = re.finditer(pattern, text, re.IGNORECASE)
            for match in matches:
                concepts.append({
                    "text": match.group(1).strip(),
                    "type": "signaled_concept",
                    "context": match.group(0),
                    "confidence": 0.9  # Explicitly signaled concepts are important
                })
        
        return concepts
    
    def semantic_extraction(self, text):
        """Extract concepts based on semantic patterns"""
        
        concepts = []
        sentences = text.split('.')
        
        for sentence in sentences:
            # Look for definition patterns
            if any(marker in sentence.lower() for marker in ['is a', 'are a', 'refers to', 'defined as']):
                # Extract subject and object
                parts = re.split(r'\bis\s+a\b|\bare\s+a\b|\brefers\s+to\b|\bdefined\s+as\b', sentence, flags=re.IGNORECASE)
                if len(parts) >= 2:
                    subject = parts[0].strip().split()[-3:]  # Last few words before marker
                    subject = ' '.join(subject).strip(' ,')
                    
                    concepts.append({
                        "text": subject,
                        "type": "defined_concept",
                        "context": sentence,
                        "confidence": 0.85,
                        "definition": parts[1].strip()
                    })
            
            # Look for causal relationships
            if any(marker in sentence.lower() for marker in ['causes', 'leads to', 'results in', 'produces']):
                concepts.append({
                    "text": sentence,
                    "type": "causal_relationship",
                    "context": sentence,
                    "confidence": 0.75
                })
        
        return concepts
```

### Phase 2: Structure Concepts
Organize extracted concepts with properties and relationships.

```python
def structure_concepts(self, concepts):
    """Structure concepts with full metadata"""
    
    structured = []
    
    for concept in concepts:
        structured_concept = {
            "concept_name": self.normalize_name(concept["text"]),
            "raw_text": concept["text"],
            "type": concept["type"],
            "definition": self.extract_definition(concept),
            "properties": self.extract_properties(concept),
            "variations": self.find_variations(concept, concepts),
            "relevance": concept.get("relevance", 0),
            "confidence": concept.get("confidence", 0),
            "context": concept.get("context", ""),
            "relationships": self.identify_relationships(concept, concepts),
            "examples": self.find_examples(concept),
            "metadata": {
                "extraction_method": concept["type"],
                "source_position": concept.get("position", -1),
                "timestamp": time.time()
            }
        }
        
        structured.append(structured_concept)
    
    return structured

def extract_definition(self, concept):
    """Extract or generate definition"""
    
    if "definition" in concept:
        return concept["definition"]
    
    # Try to find definition in context
    context = concept.get("context", "")
    
    # Look for definition patterns around the concept
    patterns = [
        f"{concept['text']} is (.*?)\\.",
        f"{concept['text']}, which is (.*?)\\.",
        f"{concept['text']} \\((.*?)\\)",
    ]
    
    for pattern in patterns:
        match = re.search(pattern, context, re.IGNORECASE)
        if match:
            return match.group(1).strip()
    
    # Generate from context if no explicit definition
    return self.generate_definition_from_context(concept)

def extract_properties(self, concept):
    """Extract concept properties"""
    
    properties = []
    context = concept.get("context", "")
    
    # Property extraction patterns
    property_patterns = [
        r'has (\w+)',
        r'with (\w+)',
        r'includes (\w+)',
        r'characterized by (\w+)',
        r'(\w+) property',
        r'(\w+) characteristic'
    ]
    
    for pattern in property_patterns:
        matches = re.findall(pattern, context, re.IGNORECASE)
        properties.extend(matches)
    
    # Deduplicate and clean
    properties = list(set(p.lower() for p in properties if len(p) > 2))
    
    return properties

def find_variations(self, concept, all_concepts):
    """Find variations of the concept"""
    
    variations = {}
    concept_text = concept["text"].lower()
    
    for other in all_concepts:
        other_text = other["text"].lower()
        
        # Skip identical concepts
        if other_text == concept_text:
            continue
        
        # Check for variations
        similarity = self.calculate_similarity(concept_text, other_text)
        
        if similarity > 0.7:  # Similar but not identical
            source = other.get("metadata", {}).get("source", "unknown")
            variations[source] = other_text
    
    return variations

def identify_relationships(self, concept, all_concepts):
    """Identify relationships with other concepts"""
    
    relationships = []
    context = concept.get("context", "")
    
    for other in all_concepts:
        if other["text"] == concept["text"]:
            continue
        
        # Check if both appear in same context
        if other["text"] in context or concept["text"] in other.get("context", ""):
            # Determine relationship type
            rel_type = self.determine_relationship_type(concept, other, context)
            
            if rel_type:
                relationships.append({
                    "target": other["text"],
                    "type": rel_type,
                    "strength": self.calculate_relationship_strength(concept, other)
                })
    
    return relationships

def determine_relationship_type(self, concept1, concept2, context):
    """Determine the type of relationship between concepts"""
    
    relationship_markers = {
        "causes": ["causes", "leads to", "results in", "produces"],
        "part_of": ["part of", "component of", "includes", "contains"],
        "example_of": ["example of", "instance of", "such as", "like"],
        "opposite_of": ["opposite of", "contrary to", "unlike", "versus"],
        "related_to": ["related to", "associated with", "connected to"],
        "depends_on": ["depends on", "requires", "needs", "based on"]
    }
    
    for rel_type, markers in relationship_markers.items():
        for marker in markers:
            pattern = f"{concept1['text']}.*{marker}.*{concept2['text']}"
            if re.search(pattern, context, re.IGNORECASE):
                return rel_type
    
    return None
```

### Phase 3: Relevance Scoring
Assign relevance scores to prioritize important concepts.

```python
def score_concepts(self, concepts, metadata=None):
    """Score concepts by relevance (1-10 scale)"""
    
    scored = []
    
    for concept in concepts:
        score = 0.0
        
        # Base score from confidence
        score += concept.get("confidence", 0.5) * 3
        
        # Frequency bonus
        frequency = self.count_occurrences(concept["text"])
        score += min(2, frequency / 5)  # Max 2 points for frequency
        
        # Type-based scoring
        type_scores = {
            "defined_concept": 2.0,
            "signaled_concept": 1.8,
            "key_phrase": 1.5,
            "causal_relationship": 1.7,
            "technical_term": 1.3,
            "entity": 1.0,
            "pattern": 1.2
        }
        score += type_scores.get(concept["type"], 0.5)
        
        # Relationship bonus
        num_relationships = len(concept.get("relationships", []))
        score += min(1.5, num_relationships * 0.3)
        
        # Properties bonus
        num_properties = len(concept.get("properties", []))
        score += min(1, num_properties * 0.2)
        
        # Context quality
        if concept.get("definition"):
            score += 0.5
        if concept.get("examples"):
            score += 0.3
        
        # Metadata bonuses
        if metadata:
            if metadata.get("source_type") == "academic":
                score += 0.5
            if metadata.get("peer_reviewed"):
                score += 0.3
        
        # Normalize to 1-10 scale
        concept["relevance"] = min(10, max(1, score))
        scored.append(concept)
    
    return sorted(scored, key=lambda x: x["relevance"], reverse=True)
```

## Integration with Zettelkasten

```python
def create_atomic_notes(extracted_concepts):
    """Convert extracted concepts to atomic notes"""
    
    notes = []
    
    for concept in extracted_concepts:
        if concept["relevance"] < 5:
            continue  # Skip low-relevance concepts
        
        note = {
            "id": generate_note_id(),
            "title": concept["concept_name"],
            "content": format_concept_as_note(concept),
            "type": "permanent",
            "tags": extract_tags(concept),
            "links": create_links(concept["relationships"]),
            "metadata": {
                "source": concept["metadata"],
                "relevance": concept["relevance"],
                "extraction_date": datetime.now(),
                "properties": concept["properties"]
            }
        }
        
        notes.append(note)
    
    return notes

def format_concept_as_note(concept):
    """Format concept as atomic note content"""
    
    content = []
    
    # Main definition
    content.append(f"# {concept['concept_name']}")
    content.append("")
    
    if concept.get("definition"):
        content.append(concept["definition"])
        content.append("")
    
    # Properties
    if concept.get("properties"):
        content.append("## Properties")
        for prop in concept["properties"]:
            content.append(f"- {prop}")
        content.append("")
    
    # Examples
    if concept.get("examples"):
        content.append("## Examples")
        for example in concept["examples"]:
            content.append(f"- {example}")
        content.append("")
    
    # Relationships
    if concept.get("relationships"):
        content.append("## Related Concepts")
        for rel in concept["relationships"]:
            content.append(f"- {rel['type']}: [[{rel['target']}]]")
        content.append("")
    
    # Variations
    if concept.get("variations"):
        content.append("## Variations")
        for source, variation in concept["variations"].items():
            content.append(f"- {source}: \"{variation}\"")
        content.append("")
    
    # Context
    if concept.get("context"):
        content.append("## Context")
        content.append(f"> {concept['context']}")
    
    return "\n".join(content)
```

## Example Extraction

```yaml
input_text: |
  "Recursion is a programming technique where a function calls itself 
  to solve a problem. The key components of recursion are the base case, 
  which provides a stopping condition, and the recursive case, which 
  breaks down the problem into smaller subproblems. Unlike iteration, 
  which uses loops, recursion uses the call stack to manage state."

extracted_concepts:
  - concept_name: "Recursion"
    type: "defined_concept"
    definition: "A programming technique where a function calls itself to solve a problem"
    properties: ["self-calling", "problem-solving", "stack-based"]
    relevance: 9.2
    relationships:
      - target: "Base Case"
        type: "part_of"
        strength: 0.9
      - target: "Recursive Case"
        type: "part_of"
        strength: 0.9
      - target: "Iteration"
        type: "opposite_of"
        strength: 0.7
  
  - concept_name: "Base Case"
    type: "signaled_concept"
    definition: "Provides a stopping condition"
    properties: ["stopping", "terminating"]
    relevance: 8.5
    relationships:
      - target: "Recursion"
        type: "component_of"
        strength: 0.9
  
  - concept_name: "Recursive Case"
    type: "signaled_concept"
    definition: "Breaks down the problem into smaller subproblems"
    properties: ["decomposing", "self-referential"]
    relevance: 8.3
    relationships:
      - target: "Recursion"
        type: "component_of"
        strength: 0.9
  
  - concept_name: "Call Stack"
    type: "technical_term"
    definition: "Manages state in recursion"
    properties: ["state-management", "memory-structure"]
    relevance: 7.8
    relationships:
      - target: "Recursion"
        type: "used_by"
        strength: 0.8
```

## Benefits

1. **Comprehensive Extraction**: Multiple strategies ensure nothing important is missed
2. **Structured Knowledge**: Concepts are fully structured with properties and relationships
3. **Relevance Prioritization**: Focus on most important concepts first
4. **Atomic Note Ready**: Output directly convertible to Zettelkasten notes
5. **Relationship Discovery**: Automatically identifies concept connections
6. **Variation Tracking**: Captures different interpretations of same concept