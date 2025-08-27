package lifecycle

import (
	"context"
	"fmt"
	"log/slog"
	"syscall"
	"time"

	"github.com/oklog/run"
)

// Manager orchestrates component lifecycle with proper dependency management
type Manager struct {
	group  *run.Group
	logger *slog.Logger
	
	// Component tracking
	components     map[string]Component
	shutdownOrder  []string
	healthChecks   map[string]HealthChecker
}

// Component represents a lifecycle-managed system component
type Component interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Dependencies() []string
}

// HealthChecker provides health check capabilities for components
type HealthChecker interface {
	HealthCheck(ctx context.Context) error
}

// ComponentRunner represents components that need continuous execution
type ComponentRunner interface {
	Run(ctx context.Context) error
}

// NewManager creates a new lifecycle manager
func NewManager(logger *slog.Logger) *Manager {
	return &Manager{
		group:         &run.Group{},
		logger:        logger,
		components:    make(map[string]Component),
		healthChecks:  make(map[string]HealthChecker),
		shutdownOrder: make([]string, 0),
	}
}

// RegisterComponent adds a component to the lifecycle manager
func (m *Manager) RegisterComponent(component Component) error {
	name := component.Name()
	
	if _, exists := m.components[name]; exists {
		return fmt.Errorf("component %s already registered", name)
	}
	
	m.components[name] = component
	
	// Add to health checks if component implements HealthChecker
	if hc, ok := component.(HealthChecker); ok {
		m.healthChecks[name] = hc
	}
	
	// Calculate shutdown order based on dependencies
	if err := m.calculateShutdownOrder(); err != nil {
		return fmt.Errorf("failed to calculate shutdown order: %w", err)
	}
	
	// Create component actor with proper oklog/run pattern - FIXED RACE CONDITION
	var runCancel context.CancelFunc
	
	m.group.Add(
		// Execute function: starts component and runs until interrupted
		func() error {
			// Create context INSIDE execute function to avoid race conditions
			runCtx, cancel := context.WithCancel(context.Background())
			runCancel = cancel // Store for interrupt handler
			defer cancel()     // Ensure cleanup
			
			m.logger.Info("starting component", slog.String("component", name))
			
			// Start the component and verify it's ready
			if err := component.Start(runCtx); err != nil {
				m.logger.Error("component failed to start",
					slog.String("component", name),
					slog.Any("error", err))
				return fmt.Errorf("component %s failed to start: %w", name, err)
			}
			
			m.logger.Info("component started successfully", slog.String("component", name))
			
			// Run the component's main loop (if it has one)
			if runner, ok := component.(ComponentRunner); ok {
				return runner.Run(runCtx)
			}
			
			// For components without a run loop, just wait for shutdown
			<-runCtx.Done()
			return runCtx.Err()
		},
		// Interrupt function: gracefully stops the component
		func(error) {
			// Signal shutdown to component if context is available
			if runCancel != nil {
				runCancel()
			}
			
			// Give component time to stop gracefully
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			
			m.logger.Info("stopping component", slog.String("component", name))
			
			if err := component.Stop(shutdownCtx); err != nil {
				m.logger.Error("component failed to stop gracefully",
					slog.String("component", name),
					slog.Any("error", err))
			} else {
				m.logger.Info("component stopped successfully", slog.String("component", name))
			}
		},
	)
	
	return nil
}

// Run starts the lifecycle manager and blocks until shutdown signal
func (m *Manager) Run() error {
	// Add signal handler for graceful shutdown
	m.group.Add(run.SignalHandler(context.Background(), 
		syscall.SIGINT, 
		syscall.SIGTERM))
	
	m.logger.Info("starting lifecycle manager", 
		slog.Int("components", len(m.components)),
		slog.Any("shutdown_order", m.shutdownOrder))
	
	err := m.group.Run()
	
	if err != nil {
		m.logger.Error("lifecycle manager stopped with error", slog.Any("error", err))
	} else {
		m.logger.Info("lifecycle manager stopped gracefully")
	}
	
	return err
}

// calculateShutdownOrder computes reverse dependency order for shutdown
func (m *Manager) calculateShutdownOrder() error {
	// Topological sort implementation for dependency resolution
	visited := make(map[string]bool)
	recursionStack := make(map[string]bool)
	order := make([]string, 0)
	
	var visit func(string) error
	visit = func(name string) error {
		if recursionStack[name] {
			return fmt.Errorf("circular dependency detected involving %s", name)
		}
		if visited[name] {
			return nil
		}
		
		visited[name] = true
		recursionStack[name] = true
		
		component := m.components[name]
		for _, dep := range component.Dependencies() {
			if _, exists := m.components[dep]; !exists {
				return fmt.Errorf("dependency %s not found for component %s", dep, name)
			}
			if err := visit(dep); err != nil {
				return err
			}
		}
		
		recursionStack[name] = false
		order = append(order, name)
		return nil
	}
	
	for name := range m.components {
		if err := visit(name); err != nil {
			return err
		}
	}
	
	// Reverse order for shutdown (dependencies shut down after dependents)
	m.shutdownOrder = make([]string, len(order))
	for i := range order {
		m.shutdownOrder[i] = order[len(order)-1-i]
	}
	
	m.logger.Debug("calculated shutdown order", 
		slog.Any("order", m.shutdownOrder))
	
	return nil
}

// HealthCheck runs health checks on all registered components
func (m *Manager) HealthCheck(ctx context.Context) error {
	for name, checker := range m.healthChecks {
		if err := checker.HealthCheck(ctx); err != nil {
			return fmt.Errorf("health check failed for component %s: %w", name, err)
		}
	}
	return nil
}

// GetComponent returns a registered component by name
func (m *Manager) GetComponent(name string) (Component, bool) {
	component, exists := m.components[name]
	return component, exists
}