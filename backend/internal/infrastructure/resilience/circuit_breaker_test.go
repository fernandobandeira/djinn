package resilience

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	
	apperrors "github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
)

func TestCircuitBreakerStates(t *testing.T) {
	stateChanges := make([]string, 0)
	var mu sync.Mutex
	
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name:         "test",
		MaxFailures:  2,
		ResetTimeout: 100 * time.Millisecond,
		HalfOpenMax:  1,
		OnStateChange: func(name string, from State, to State) {
			mu.Lock()
			defer mu.Unlock()
			stateChanges = append(stateChanges, fmt.Sprintf("%s->%s", from, to))
		},
	})
	
	ctx := context.Background()
	successFn := func() error { return nil }
	failureFn := func() error { return errors.New("failure") }
	
	// Initial state should be closed
	assert.Equal(t, StateClosed, cb.GetState())
	
	// First execution should succeed
	err := cb.Execute(ctx, successFn)
	assert.NoError(t, err)
	assert.Equal(t, StateClosed, cb.GetState())
	
	// First failure
	err = cb.Execute(ctx, failureFn)
	assert.Error(t, err)
	assert.Equal(t, StateClosed, cb.GetState())
	assert.Equal(t, uint32(1), cb.failures)
	
	// Second failure should open the circuit
	err = cb.Execute(ctx, failureFn)
	assert.Error(t, err)
	assert.Equal(t, StateOpen, cb.GetState())
	assert.Equal(t, uint32(2), cb.failures)
	
	// Should reject calls when open
	err = cb.Execute(ctx, successFn)
	assert.True(t, errors.Is(err, apperrors.ErrCircuitOpen))
	
	// Wait for reset timeout
	time.Sleep(150 * time.Millisecond)
	
	// Should transition to half-open and then immediately to closed after success
	err = cb.Execute(ctx, successFn)
	assert.NoError(t, err)
	// State transitions immediately to closed after successful half-open test
	assert.Equal(t, StateClosed, cb.GetState())
	
	// Verify state changes
	mu.Lock()
	defer mu.Unlock()
	assert.Contains(t, stateChanges, "closed->open")
	assert.Contains(t, stateChanges, "open->half-open")
	assert.Contains(t, stateChanges, "half-open->closed")
}

func TestCircuitBreakerHalfOpenFailure(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name:         "test",
		MaxFailures:  1,
		ResetTimeout: 50 * time.Millisecond,
		HalfOpenMax:  1,
	})
	
	ctx := context.Background()
	failureFn := func() error { return errors.New("failure") }
	
	// Open the circuit
	cb.Execute(ctx, failureFn)
	assert.Equal(t, StateOpen, cb.GetState())
	
	// Wait for reset timeout
	time.Sleep(60 * time.Millisecond)
	
	// Failure in half-open should re-open the circuit
	err := cb.Execute(ctx, failureFn)
	assert.Error(t, err)
	assert.Equal(t, StateOpen, cb.GetState())
}

func TestCircuitBreakerReset(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name:        "test",
		MaxFailures: 1,
	})
	
	ctx := context.Background()
	failureFn := func() error { return errors.New("failure") }
	
	// Open the circuit
	cb.Execute(ctx, failureFn)
	assert.Equal(t, StateOpen, cb.GetState())
	assert.Equal(t, uint32(1), cb.failures)
	
	// Manual reset
	cb.Reset()
	assert.Equal(t, StateClosed, cb.GetState())
	assert.Equal(t, uint32(0), cb.failures)
}

func TestCircuitBreakerConcurrency(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name:         "test",
		MaxFailures:  5,
		ResetTimeout: 100 * time.Millisecond,
	})
	
	ctx := context.Background()
	failureFn := func() error { return errors.New("failure") }
	
	var wg sync.WaitGroup
	
	// Run multiple concurrent failures
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.Execute(ctx, failureFn)
		}()
	}
	
	wg.Wait()
	
	// Circuit should be open after max failures
	assert.Equal(t, StateOpen, cb.GetState())
}

func TestCircuitBreakerStats(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name:        "test-circuit",
		MaxFailures: 2,
	})
	
	ctx := context.Background()
	failureFn := func() error { return errors.New("failure") }
	
	// Execute a failure
	cb.Execute(ctx, failureFn)
	
	stats := cb.GetStats()
	assert.Equal(t, "test-circuit", stats.Name)
	assert.Equal(t, "closed", stats.State)
	assert.Equal(t, uint32(1), stats.Failures)
	assert.NotZero(t, stats.LastFailTime)
}

func TestCircuitBreakerWithContext(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		Name: "test",
	})
	
	// Test with correlation ID in context
	ctx := context.WithValue(context.Background(), "correlation_id", "test-123")
	
	successFn := func() error { return nil }
	err := cb.Execute(ctx, successFn)
	assert.NoError(t, err)
}