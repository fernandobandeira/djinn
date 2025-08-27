package lifecycle

import (
	"context"
	"log/slog"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

// TestMain ensures no goroutine leaks across the entire test package
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

// MockComponent for testing lifecycle management
type MockComponent struct {
	name         string
	dependencies []string
	startCalled  bool
	stopCalled   bool
	startError   error
	stopError    error
	startDelay   time.Duration
	stopDelay    time.Duration
}

func (m *MockComponent) Name() string { return m.name }
func (m *MockComponent) Dependencies() []string { return m.dependencies }

func (m *MockComponent) Start(ctx context.Context) error {
	if m.startDelay > 0 {
		time.Sleep(m.startDelay)
	}
	m.startCalled = true
	return m.startError
}

func (m *MockComponent) Stop(ctx context.Context) error {
	if m.stopDelay > 0 {
		select {
		case <-time.After(m.stopDelay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	m.stopCalled = true
	return m.stopError
}

func TestLifecycleManager_NoGoroutineLeaks(t *testing.T) {
	defer goleak.VerifyNone(t) // ✅ Automatic goroutine leak detection

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	manager := NewManager(logger)

	// Register a simple component
	component := &MockComponent{name: "test-component"}
	err := manager.RegisterComponent(component)
	require.NoError(t, err)

	// Start manager in background
	done := make(chan error, 1)
	go func() {
		done <- manager.Run()
	}()

	// Let it start
	time.Sleep(50 * time.Millisecond)

	// Send shutdown signal
	err = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	require.NoError(t, err)

	// Wait for shutdown with reasonable timeout
	select {
	case err := <-done:
		assert.NoError(t, err)
	case <-time.After(5 * time.Second):
		t.Fatal("Manager failed to shutdown within timeout")
	}

	// Verify component was started and stopped
	assert.True(t, component.startCalled, "Component should have been started")
	assert.True(t, component.stopCalled, "Component should have been stopped")

	// goleak will automatically fail test if goroutines leaked
}

func TestLifecycleManager_ComponentStartupFailure(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	manager := NewManager(logger)

	// Component that fails to start
	component := &MockComponent{
		name:       "failing-component",
		startError: assert.AnError,
	}
	err := manager.RegisterComponent(component)
	require.NoError(t, err)

	// Run should fail immediately
	err = manager.Run()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start")

	// No goroutines should leak even on failure
}

func TestLifecycleManager_MultipleComponents(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	manager := NewManager(logger)

	// Register multiple components
	components := []*MockComponent{
		{name: "component-1"},
		{name: "component-2"},
		{name: "component-3"},
	}

	for _, comp := range components {
		err := manager.RegisterComponent(comp)
		require.NoError(t, err)
	}

	// Start and shutdown cycle
	done := make(chan error, 1)
	go func() {
		done <- manager.Run()
	}()

	time.Sleep(50 * time.Millisecond)

	err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	require.NoError(t, err)

	select {
	case err := <-done:
		assert.NoError(t, err)
	case <-time.After(5 * time.Second):
		t.Fatal("Manager failed to shutdown within timeout")
	}

	// All components should be started and stopped
	for i, comp := range components {
		assert.True(t, comp.startCalled, "Component %d should have been started", i)
		assert.True(t, comp.stopCalled, "Component %d should have been stopped", i)
	}
}