# Multi-Phase Verification Framework

## Purpose
Systematic verification approach shared across all agents to ensure quality, correctness, and learning.

## Core Philosophy
Every significant task should be verified through multiple phases to catch issues early and learn from both successes and failures.

## Framework Structure

### Phase 1: Parse & Validate Input
**Objective**: Ensure we understand what's being asked and have everything needed.

```yaml
parse_validation:
  checks:
    - input_completeness: "All required information provided"
    - input_clarity: "No ambiguities in request"
    - prerequisites_met: "Dependencies and context available"
    - scope_defined: "Clear boundaries of task"
    
  actions:
    - identify_ambiguities: "Flag unclear requirements"
    - request_clarification: "Ask for missing information"
    - validate_prerequisites: "Check all dependencies exist"
    - define_success_criteria: "Establish measurable outcomes"
    
  outputs:
    - structured_requirements: "Parsed and validated input"
    - validation_report: "What passed/failed"
    - clarification_requests: "What needs user input"
```

### Phase 2: Process & Monitor
**Objective**: Verify each step during execution.

```yaml
process_monitoring:
  checks:
    - step_completion: "Each step produces expected output"
    - intermediate_validity: "Intermediate results are correct"
    - constraint_satisfaction: "All constraints being met"
    - resource_availability: "Required resources accessible"
    
  actions:
    - validate_step_output: "Check each step's results"
    - track_state_changes: "Monitor system state evolution"
    - detect_anomalies: "Flag unexpected behaviors"
    - checkpoint_progress: "Save state at key points"
    
  outputs:
    - step_validations: "Per-step verification results"
    - state_timeline: "How system evolved"
    - anomaly_log: "Unexpected occurrences"
```

### Phase 3: Synthesize & Verify Output
**Objective**: Ensure final output meets all requirements.

```yaml
output_synthesis:
  checks:
    - completeness: "All requirements addressed"
    - correctness: "Output is accurate and valid"
    - coherence: "All parts work together"
    - optimization: "Output is efficient and clean"
    
  actions:
    - verify_against_requirements: "Match output to original ask"
    - check_internal_consistency: "Ensure coherent whole"
    - validate_integration_points: "Test connections"
    - optimize_if_needed: "Improve where possible"
    
  outputs:
    - final_output: "Verified and optimized result"
    - verification_report: "What was checked and results"
    - optimization_suggestions: "Potential improvements"
```

### Phase 4: Learn & Improve
**Objective**: Extract patterns and learnings for future use.

```yaml
learning_extraction:
  checks:
    - pattern_identification: "What patterns emerged"
    - success_factors: "What worked well"
    - failure_points: "What could be better"
    - reusability: "What can be generalized"
    
  actions:
    - extract_patterns: "Document successful patterns"
    - analyze_failures: "Understand what went wrong"
    - update_knowledge_base: "Store learnings"
    - improve_processes: "Refine verification itself"
    
  outputs:
    - learned_patterns: "New patterns discovered"
    - improvement_recommendations: "How to do better"
    - kb_updates: "Knowledge base additions"
```

## Implementation by Agent Type

### For Rita (Recruiter)
```python
class RitaVerification(MultiPhaseVerification):
    def phase1_parse(self, agent_request):
        """Parse agent creation request"""
        return {
            "agent_type": self.determine_type(agent_request),
            "requirements": self.extract_requirements(agent_request),
            "existing_agents": self.check_duplicates(agent_request),
            "validation": self.validate_feasibility(agent_request)
        }
    
    def phase2_process(self, creation_state):
        """Monitor agent creation"""
        return {
            "component_validation": self.validate_each_component(),
            "constraint_tracking": self.monitor_constraints(),
            "coherence_checks": self.verify_coherence()
        }
    
    def phase3_synthesize(self, agent_config):
        """Verify complete agent"""
        return {
            "syntax_validation": self.check_syntax(),
            "integration_testing": self.test_delegation(),
            "constraint_score": self.calculate_constraint_score()
        }
    
    def phase4_learn(self, creation_result):
        """Learn from creation"""
        return {
            "patterns": self.extract_successful_patterns(),
            "improvements": self.identify_optimizations(),
            "kb_update": self.update_pattern_library()
        }
```

### For Tina (Teacher)
```python
class TinaVerification(MultiPhaseVerification):
    def phase1_parse(self, learning_request):
        """Parse learning objective"""
        return {
            "topic": self.identify_topic(learning_request),
            "prerequisites": self.check_prior_knowledge(),
            "learning_style": self.detect_learner_profile(),
            "success_metrics": self.define_understanding_criteria()
        }
    
    def phase2_process(self, teaching_session):
        """Monitor learning progress"""
        return {
            "comprehension_checks": self.assess_understanding(),
            "engagement_tracking": self.monitor_engagement(),
            "pace_adjustment": self.adapt_teaching_pace()
        }
    
    def phase3_synthesize(self, session_content):
        """Verify learning outcomes"""
        return {
            "objective_completion": self.check_learning_goals(),
            "knowledge_synthesis": self.create_summary(),
            "retention_plan": self.schedule_reviews()
        }
    
    def phase4_learn(self, session_result):
        """Learn from teaching"""
        return {
            "effective_methods": self.identify_what_worked(),
            "struggle_points": self.note_difficulties(),
            "teaching_improvements": self.refine_approach()
        }
```

### For Zettelkasten Guide
```python
class ZettelkastenVerification(MultiPhaseVerification):
    def phase1_parse(self, note_input):
        """Parse note content"""
        return {
            "concepts": self.extract_concepts(note_input),
            "connections": self.identify_links(),
            "uniqueness": self.check_duplicates(),
            "atomicity": self.verify_single_idea()
        }
    
    def phase2_process(self, note_creation):
        """Monitor note integration"""
        return {
            "link_validation": self.verify_connections(),
            "hierarchy_check": self.validate_structure(),
            "metadata_tracking": self.ensure_proper_tagging()
        }
    
    def phase3_synthesize(self, note_network):
        """Verify knowledge graph"""
        return {
            "graph_coherence": self.check_network_integrity(),
            "cluster_identification": self.find_knowledge_clusters(),
            "synthesis_opportunities": self.identify_meta_insights()
        }
    
    def phase4_learn(self, note_result):
        """Learn from knowledge capture"""
        return {
            "emerging_patterns": self.detect_knowledge_patterns(),
            "connection_strength": self.analyze_link_quality(),
            "graph_evolution": self.track_knowledge_growth()
        }
```

## Verification Strategies

### Early Termination
Stop processing if critical validation fails:

```python
def execute_with_verification(task):
    # Phase 1
    parsed = phase1_parse(task)
    if not parsed["validation"]["passed"]:
        return {"error": "Failed initial validation", "details": parsed}
    
    # Phase 2
    for step in process_steps(task):
        validation = validate_step(step)
        if not validation["passed"]:
            return {"error": f"Failed at step {step.name}", "details": validation}
    
    # Continue only if all steps pass
```

### Progressive Verification
Build confidence through incremental checks:

```python
def progressive_verification(components):
    confidence_score = 0.0
    verified_components = []
    
    for component in components:
        # Verify component
        result = verify_component(component)
        
        # Update confidence
        confidence_score = (confidence_score * len(verified_components) + 
                          result["score"]) / (len(verified_components) + 1)
        
        # Stop if confidence drops too low
        if confidence_score < 0.6:
            return {
                "status": "failed",
                "confidence": confidence_score,
                "failed_at": component.name
            }
        
        verified_components.append(component)
    
    return {
        "status": "passed",
        "confidence": confidence_score,
        "components": verified_components
    }
```

### Parallel Verification
Run independent checks simultaneously:

```python
async def parallel_verification(agent_config):
    # Run all independent checks in parallel
    results = await asyncio.gather(
        verify_syntax(agent_config),
        verify_constraints(agent_config),
        verify_resources(agent_config),
        verify_integration(agent_config)
    )
    
    # Aggregate results
    return {
        "syntax": results[0],
        "constraints": results[1],
        "resources": results[2],
        "integration": results[3],
        "overall": all(r["passed"] for r in results)
    }
```

## Verification Metrics

### Success Metrics
```yaml
metrics:
  accuracy:
    description: "Correctness of output"
    measurement: "Percentage of requirements met"
    target: "> 95%"
    
  completeness:
    description: "All aspects addressed"
    measurement: "Coverage of requirements"
    target: "100%"
    
  efficiency:
    description: "Resource usage"
    measurement: "Time and token usage vs baseline"
    target: "< 120% of baseline"
    
  learning_rate:
    description: "Pattern extraction success"
    measurement: "New patterns per task"
    target: "> 0.5 patterns/task"
    
  error_prevention:
    description: "Issues caught early"
    measurement: "Phase 1+2 catches vs Phase 3"
    target: "> 80% early detection"
```

### Failure Analysis
```yaml
failure_tracking:
  categories:
    - input_issues: "Problems with requirements"
    - process_failures: "Execution problems"
    - integration_errors: "Component mismatch"
    - constraint_violations: "Rules broken"
    
  recovery_strategies:
    input_issues: "Request clarification"
    process_failures: "Rollback and retry"
    integration_errors: "Fix interfaces"
    constraint_violations: "Negotiate or adjust"
    
  learning_focus:
    - "Why did this fail?"
    - "Could we have detected earlier?"
    - "What pattern prevents this?"
    - "How do we recover gracefully?"
```

## Integration Guidelines

### For New Agents
1. Inherit from `MultiPhaseVerification` base
2. Implement all four phases
3. Define agent-specific checks
4. Connect to knowledge base
5. Enable learning extraction

### For Existing Agents
1. Wrap existing logic in phases
2. Add verification checkpoints
3. Extract patterns from successes
4. Document failures for learning
5. Gradually increase verification depth

## Usage Examples

### Basic Usage
```python
# Any agent can use the framework
verifier = MultiPhaseVerification()

# Phase 1
input_valid = verifier.phase1_parse(user_request)
if not input_valid["passed"]:
    return handle_invalid_input(input_valid)

# Phase 2
process_result = verifier.phase2_process(execution_state)
if process_result["anomalies"]:
    handle_anomalies(process_result["anomalies"])

# Phase 3
output = verifier.phase3_synthesize(raw_result)
if not output["meets_requirements"]:
    attempt_correction(output)

# Phase 4
learning = verifier.phase4_learn(entire_process)
update_knowledge_base(learning["patterns"])
```

### Advanced Usage
```python
# Custom verification pipeline
class CustomVerification(MultiPhaseVerification):
    def __init__(self):
        super().__init__()
        self.add_custom_phase("security_check", after="phase2")
        self.add_custom_phase("performance_test", after="phase3")
    
    def security_check(self, state):
        """Custom security validation"""
        return validate_security(state)
    
    def performance_test(self, output):
        """Custom performance validation"""
        return measure_performance(output)
```

## Benefits

1. **Systematic Quality**: Every task verified thoroughly
2. **Early Error Detection**: Catch issues before they compound
3. **Continuous Learning**: Every task improves the system
4. **Cross-Agent Consistency**: All agents follow same quality bar
5. **Traceable Validation**: Clear record of what was checked