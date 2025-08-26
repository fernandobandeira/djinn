package resilience

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
)

// State represents the circuit breaker state
type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// CircuitBreakerConfig holds configuration for circuit breaker
type CircuitBreakerConfig struct {
	Name            string
	MaxFailures     uint32
	ResetTimeout    time.Duration
	HalfOpenMax     uint32
	OnStateChange   func(name string, from State, to State)
	Logger          *slog.Logger
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	name          string
	maxFailures   uint32
	resetTimeout  time.Duration
	halfOpenMax   uint32
	
	mu            sync.RWMutex
	state         State
	failures      uint32
	lastFailTime  time.Time
	halfOpenCount uint32
	
	onStateChange func(name string, from State, to State)
	logger        *slog.Logger
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
	if config.MaxFailures == 0 {
		config.MaxFailures = 5
	}
	if config.ResetTimeout == 0 {
		config.ResetTimeout = 60 * time.Second
	}
	if config.HalfOpenMax == 0 {
		config.HalfOpenMax = 2
	}
	if config.Logger == nil {
		config.Logger = slog.Default()
	}

	cb := &CircuitBreaker{
		name:          config.Name,
		maxFailures:   config.MaxFailures,
		resetTimeout:  config.ResetTimeout,
		halfOpenMax:   config.HalfOpenMax,
		state:         StateClosed,
		onStateChange: config.OnStateChange,
		logger:        config.Logger,
	}

	return cb
}

// Execute runs the function with circuit breaker protection
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() error) error {
	correlationID := getCorrelationID(ctx)
	
	if err := cb.canExecute(); err != nil {
		cb.logger.WarnContext(ctx, "circuit breaker preventing execution",
			"circuit", cb.name,
			"state", cb.state.String(),
			"correlation_id", correlationID)
		return err
	}

	err := fn()
	cb.recordResult(err)

	if err != nil {
		cb.logger.DebugContext(ctx, "circuit breaker execution failed",
			"circuit", cb.name,
			"state", cb.state.String(),
			"failures", cb.failures,
			"correlation_id", correlationID,
			"error", err.Error())
	}

	return err
}

// canExecute checks if execution is allowed
func (cb *CircuitBreaker) canExecute() error {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	switch cb.state {
	case StateClosed:
		return nil
	case StateOpen:
		if time.Since(cb.lastFailTime) > cb.resetTimeout {
			cb.mu.RUnlock()
			cb.mu.Lock()
			if cb.state == StateOpen {
				cb.transitionTo(StateHalfOpen)
			}
			cb.mu.Unlock()
			cb.mu.RLock()
			return nil
		}
		return fmt.Errorf("%w: %s", errors.ErrCircuitOpen, cb.name)
	case StateHalfOpen:
		if cb.halfOpenCount >= cb.halfOpenMax {
			return fmt.Errorf("%w: %s (half-open limit reached)", errors.ErrCircuitOpen, cb.name)
		}
		return nil
	default:
		return fmt.Errorf("unknown circuit breaker state: %v", cb.state)
	}
}

// recordResult records the execution result
func (cb *CircuitBreaker) recordResult(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case StateClosed:
		if err != nil {
			cb.failures++
			cb.lastFailTime = time.Now()
			if cb.failures >= cb.maxFailures {
				cb.transitionTo(StateOpen)
				cb.logger.Error("circuit breaker opened",
					"circuit", cb.name,
					"failures", cb.failures,
					"max_failures", cb.maxFailures)
			}
		}

	case StateHalfOpen:
		cb.halfOpenCount++
		if err != nil {
			cb.failures++
			cb.lastFailTime = time.Now()
			cb.transitionTo(StateOpen)
			cb.logger.Warn("circuit breaker re-opened from half-open",
				"circuit", cb.name)
		} else if cb.halfOpenCount >= cb.halfOpenMax {
			cb.failures = 0
			cb.halfOpenCount = 0
			cb.transitionTo(StateClosed)
			cb.logger.Info("circuit breaker closed after successful half-open tests",
				"circuit", cb.name)
		}

	case StateOpen:
		// Should not reach here due to canExecute check
		if err != nil {
			cb.lastFailTime = time.Now()
		}
	}
}

// transitionTo changes the circuit breaker state
func (cb *CircuitBreaker) transitionTo(newState State) {
	if cb.state == newState {
		return
	}

	oldState := cb.state
	cb.state = newState
	
	if newState == StateHalfOpen {
		cb.halfOpenCount = 0
	}

	cb.logger.Info("circuit breaker state changed",
		"circuit", cb.name,
		"from", oldState.String(),
		"to", newState.String())

	if cb.onStateChange != nil {
		cb.onStateChange(cb.name, oldState, newState)
	}
}

// GetState returns the current state
func (cb *CircuitBreaker) GetState() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// GetStats returns circuit breaker statistics
func (cb *CircuitBreaker) GetStats() CircuitBreakerStats {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	
	return CircuitBreakerStats{
		Name:         cb.name,
		State:        cb.state.String(),
		Failures:     cb.failures,
		LastFailTime: cb.lastFailTime,
		HalfOpenCount: cb.halfOpenCount,
	}
}

// Reset manually resets the circuit breaker
func (cb *CircuitBreaker) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	cb.failures = 0
	cb.halfOpenCount = 0
	if cb.state != StateClosed {
		cb.transitionTo(StateClosed)
		cb.logger.Info("circuit breaker manually reset",
			"circuit", cb.name)
	}
}

// CircuitBreakerStats holds circuit breaker statistics
type CircuitBreakerStats struct {
	Name          string    `json:"name"`
	State         string    `json:"state"`
	Failures      uint32    `json:"failures"`
	LastFailTime  time.Time `json:"last_fail_time,omitempty"`
	HalfOpenCount uint32    `json:"half_open_count,omitempty"`
}

// getCorrelationID extracts correlation ID from context
func getCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value("correlation_id").(string); ok {
		return id
	}
	if id, ok := ctx.Value("request_id").(string); ok {
		return id
	}
	return "unknown"
}