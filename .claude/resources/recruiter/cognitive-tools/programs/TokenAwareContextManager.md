# Token-Aware Context Manager - Cognitive Tool

## Purpose
Intelligently manage context to stay within token limits while maximizing relevant information inclusion.

## Core Concept
Build context progressively with prioritization, summarization, and intelligent truncation.

## Implementation

### Token Budget Management

```python
class TokenAwareContextManager:
    def __init__(self, max_tokens=8000):
        self.max_tokens = max_tokens
        self.safety_margin = 0.9  # Use only 90% to leave room
        self.effective_limit = int(max_tokens * self.safety_margin)
        self.current_usage = 0
        self.context_buffer = []
        self.priority_queue = []
    
    def estimate_tokens(self, text):
        """Rough token estimation (4 chars ≈ 1 token)"""
        return len(text) // 4
    
    def add_context(self, content, priority=5, category="general"):
        """Add content with priority (1=highest, 10=lowest)"""
        tokens = self.estimate_tokens(content)
        self.priority_queue.append({
            "content": content,
            "tokens": tokens,
            "priority": priority,
            "category": category,
            "added_at": time.time()
        })
    
    def build_context(self):
        """Build optimized context within token limits"""
        # Sort by priority (ascending = highest first)
        sorted_items = sorted(
            self.priority_queue, 
            key=lambda x: (x["priority"], -x["added_at"])
        )
        
        final_context = []
        token_count = 0
        
        for item in sorted_items:
            if token_count + item["tokens"] > self.effective_limit:
                # Try to summarize if possible
                summary = self.summarize_if_needed(
                    item["content"], 
                    self.effective_limit - token_count
                )
                if summary:
                    final_context.append(f"[Summarized {item['category']}]: {summary}")
                    token_count += self.estimate_tokens(summary)
                break
            else:
                final_context.append(f"[{item['category']}]: {item['content']}")
                token_count += item["tokens"]
        
        self.current_usage = token_count
        return "\n\n".join(final_context)
    
    def summarize_if_needed(self, content, available_tokens):
        """Intelligent summarization when content exceeds limits"""
        if available_tokens < 100:
            return None  # Too small to be useful
        
        # Different summarization strategies by content type
        if "```" in content:  # Code block
            return self.summarize_code(content, available_tokens)
        elif len(content.split("\n")) > 10:  # Multi-line content
            return self.summarize_structured(content, available_tokens)
        else:
            return self.summarize_prose(content, available_tokens)
    
    def summarize_code(self, code, max_tokens):
        """Summarize code intelligently"""
        lines = code.split("\n")
        
        # Keep signatures and key logic
        important_lines = []
        for line in lines:
            if any(keyword in line for keyword in 
                   ["def ", "class ", "return", "import", "# TODO", "# IMPORTANT"]):
                important_lines.append(line)
        
        summary = "\n".join(important_lines[:max_tokens // 10])
        return f"Code outline:\n{summary}\n[{len(lines)} total lines]"
    
    def summarize_structured(self, content, max_tokens):
        """Summarize structured content (YAML, JSON, etc.)"""
        lines = content.split("\n")
        
        # Keep headers and key fields
        summary_lines = []
        indent_level = 0
        
        for line in lines:
            current_indent = len(line) - len(line.lstrip())
            
            # Include top-level and important nested items
            if current_indent <= 4 or any(key in line for key in 
                                         ["name:", "type:", "required:", "description:"]):
                summary_lines.append(line)
                
                if self.estimate_tokens("\n".join(summary_lines)) > max_tokens * 0.8:
                    break
        
        return "\n".join(summary_lines) + f"\n[... {len(lines) - len(summary_lines)} lines omitted]"
    
    def summarize_prose(self, content, max_tokens):
        """Summarize prose text"""
        sentences = content.split(". ")
        
        # Take first and last sentences, key points in middle
        if len(sentences) > 3:
            summary = f"{sentences[0]}. [...] {sentences[-1]}"
        else:
            summary = content[:max_tokens * 4]  # Direct truncation
        
        return summary
```

### Context Building Strategies

```python
def build_pattern_context(pattern_type, max_tokens=2000):
    """Build context for pattern search with smart inclusion"""
    
    manager = TokenAwareContextManager(max_tokens)
    
    # Search for patterns
    patterns = search_kb(pattern_type)
    
    # Categorize and prioritize patterns
    for pattern in patterns:
        # Calculate relevance score
        relevance = calculate_relevance(pattern, pattern_type)
        
        # Determine priority (1-10 scale)
        if relevance > 0.9:
            priority = 1  # Must include
        elif relevance > 0.7:
            priority = 3  # Should include
        elif relevance > 0.5:
            priority = 5  # Nice to include
        else:
            priority = 8  # Include if space
        
        # Categorize pattern
        if "constraint" in pattern.lower():
            category = "constraint"
        elif "protocol" in pattern.lower():
            category = "protocol"
        elif "example" in pattern.lower():
            category = "example"
        else:
            category = "general"
        
        manager.add_context(pattern, priority, category)
    
    # Build optimized context
    return manager.build_context()

def calculate_relevance(pattern, search_term):
    """Calculate relevance score between pattern and search term"""
    
    # Simple relevance scoring
    score = 0.0
    
    # Direct match
    if search_term.lower() in pattern.lower():
        score += 0.5
    
    # Keyword overlap
    search_words = set(search_term.lower().split())
    pattern_words = set(pattern.lower().split())
    overlap = len(search_words & pattern_words) / len(search_words)
    score += overlap * 0.3
    
    # Recency (if pattern has timestamp)
    # Assume newer patterns are more relevant
    if "2024" in pattern or "2025" in pattern:
        score += 0.2
    
    return min(score, 1.0)
```

### Progressive Context Building

```python
class ProgressiveContextBuilder:
    """Build context progressively as agent creation proceeds"""
    
    def __init__(self, total_budget=30000):
        self.total_budget = total_budget
        self.phase_budgets = {
            "requirements": 3000,
            "planning": 5000,
            "implementation": 15000,
            "validation": 5000,
            "documentation": 2000
        }
        self.phase_contexts = {}
        self.current_phase = "requirements"
    
    def enter_phase(self, phase_name):
        """Switch to new phase with its budget"""
        self.current_phase = phase_name
        budget = self.phase_budgets.get(phase_name, 5000)
        return TokenAwareContextManager(budget)
    
    def transfer_context(self, from_phase, to_phase, summary_ratio=0.3):
        """Transfer summarized context between phases"""
        
        if from_phase not in self.phase_contexts:
            return ""
        
        old_context = self.phase_contexts[from_phase]
        
        # Summarize old context to use only 30% of new phase budget
        new_budget = self.phase_budgets.get(to_phase, 5000)
        summary_budget = int(new_budget * summary_ratio)
        
        manager = TokenAwareContextManager(summary_budget)
        manager.add_context(old_context, priority=7, category=f"from_{from_phase}")
        
        return manager.build_context()
```

### Integration Examples

```python
# Example 1: Agent Creation with Progressive Context
def create_agent_with_context(agent_name):
    builder = ProgressiveContextBuilder()
    
    # Phase 1: Requirements
    req_manager = builder.enter_phase("requirements")
    req_manager.add_context(user_requirements, priority=1)
    req_manager.add_context(similar_agents, priority=3)
    requirements_context = req_manager.build_context()
    
    # Phase 2: Planning (with summarized requirements)
    plan_manager = builder.enter_phase("planning")
    plan_manager.add_context(
        builder.transfer_context("requirements", "planning"),
        priority=5
    )
    plan_manager.add_context(design_patterns, priority=2)
    planning_context = plan_manager.build_context()
    
    # Continue through phases...
    return final_agent

# Example 2: Pattern Search with Token Awareness
def search_patterns_smart(query, max_results=10):
    # Allocate token budget
    budget = 4000
    
    # Search and rank patterns
    patterns = search_kb(query)
    ranked = rank_by_relevance(patterns, query)
    
    # Build context within budget
    manager = TokenAwareContextManager(budget)
    
    for i, pattern in enumerate(ranked[:max_results]):
        priority = min(i + 1, 10)  # First results get higher priority
        manager.add_context(pattern["content"], priority, pattern["type"])
    
    return manager.build_context()
```

## Usage in Rita's Workflow

### When Creating Agents
```python
# At start of agent creation
context_builder = ProgressiveContextBuilder()

# During requirements gathering
req_manager = context_builder.enter_phase("requirements")
req_manager.add_context(user_input, priority=1, category="user_requirements")
req_manager.add_context(discovered_patterns, priority=3, category="patterns")

# Get optimized context for processing
working_context = req_manager.build_context()
```

### When Searching Patterns
```python
# Search with token awareness
def get_relevant_patterns(search_term):
    manager = TokenAwareContextManager(max_tokens=3000)
    
    # Add patterns by relevance
    patterns = kb_search(search_term)
    for pattern in patterns:
        score = calculate_relevance(pattern, search_term)
        priority = int((1 - score) * 9) + 1  # Convert score to priority
        manager.add_context(pattern, priority, "pattern")
    
    return manager.build_context()
```

### When Validating
```python
# Validation with focused context
def validate_with_context(agent_config):
    manager = TokenAwareContextManager(max_tokens=2000)
    
    # High priority: Current agent config
    manager.add_context(agent_config, priority=1, category="target")
    
    # Medium priority: Validation rules
    manager.add_context(validation_rules, priority=3, category="rules")
    
    # Low priority: Examples
    manager.add_context(validation_examples, priority=7, category="examples")
    
    context = manager.build_context()
    return perform_validation(context)
```

## Benefits

1. **Token Efficiency**: Never exceed limits, always include most relevant info
2. **Progressive Building**: Context evolves with agent creation phases
3. **Intelligent Summarization**: Preserves key information when truncating
4. **Priority-Based Inclusion**: Most important information always included
5. **Category Awareness**: Different content types handled appropriately

## Integration Points

- **ProgressiveAgentPlanner**: Use for phase-based context management
- **Component Coherence Verification**: Focus context on specific components
- **Pattern Learning**: Efficiently include learned patterns
- **User Expertise Profiles**: Adjust context detail based on user level