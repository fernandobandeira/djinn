package observability

import (
	"context"
	"fmt"
	"net/http"
	"log/slog"
	"sync"
	"time"

	"github.com/posthog/posthog-go"
)

// AnalyticsTracker handles user behavior analytics with PostHog
type AnalyticsTracker struct {
	client       posthog.Client
	mu           sync.RWMutex
	breadcrumbs  map[string][]Breadcrumb
	enabled      bool
	logger       *slog.Logger
	failedEvents chan interface{} // Buffer for failed events
	stopCleanup  chan struct{}
}

// Breadcrumb represents a user action or system event for crash context
type Breadcrumb struct {
	Timestamp time.Time              `json:"timestamp"`
	Category  string                 `json:"category"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// AnalyticsConfig holds configuration for PostHog analytics
type AnalyticsConfig struct {
	APIKey      string
	Enabled     bool
	Environment string
}

// NewAnalyticsTracker creates a new analytics tracker
func NewAnalyticsTracker(config AnalyticsConfig, logger *slog.Logger) (*AnalyticsTracker, error) {
	tracker := &AnalyticsTracker{
		enabled:      config.Enabled,
		breadcrumbs:  make(map[string][]Breadcrumb),
		logger:       logger,
		failedEvents: make(chan interface{}, 1000), // Buffer up to 1000 failed events
		stopCleanup:  make(chan struct{}),
	}

	if !config.Enabled || config.APIKey == "" {
		// Return disabled tracker
		tracker.enabled = false
		return tracker, nil
	}

	client := posthog.New(config.APIKey)
	tracker.client = client

	// Start breadcrumb cleanup goroutine
	go tracker.cleanupBreadcrumbs()

	// Start failed events retry goroutine
	go tracker.retryFailedEvents()

	return tracker, nil
}

// TrackUserAction tracks custom user behavior events
func (a *AnalyticsTracker) TrackUserAction(userID, action string, properties map[string]interface{}) {
	if !a.enabled || a.client == nil {
		return
	}

	// Add correlation ID to breadcrumbs if present
	correlationID := ""
	if id, ok := properties["correlation_id"].(string); ok {
		correlationID = id
	}

	// Add to breadcrumbs
	a.AddBreadcrumb(correlationID, Breadcrumb{
		Timestamp: time.Now(),
		Category:  "user_action",
		Message:   action,
		Data:      properties,
	})

	// Create PostHog properties
	props := posthog.NewProperties()
	for k, v := range properties {
		props.Set(k, v)
	}

	// Track in PostHog with error handling
	capture := posthog.Capture{
		DistinctId: userID,
		Event:      action,
		Properties: props,
	}
	
	if err := a.client.Enqueue(capture); err != nil {
		if a.logger != nil {
			a.logger.Warn("Failed to enqueue PostHog event",
				slog.String("event", action),
				slog.String("error", err.Error()),
			)
		}
		// Try to buffer for retry
		select {
		case a.failedEvents <- capture:
		default:
			// Buffer full, drop event
			if a.logger != nil {
				a.logger.Error("Failed event buffer full, dropping event",
					slog.String("event", action),
				)
			}
		}
	}
}

// TrackRequest tracks HTTP request for analytics
func (a *AnalyticsTracker) TrackRequest(userID string, r *http.Request, statusCode int, duration float64) {
	if !a.enabled || a.client == nil || userID == "" {
		return
	}

	properties := posthog.NewProperties().
		Set("method", r.Method).
		Set("path", r.URL.Path).
		Set("status_code", statusCode).
		Set("duration_ms", duration*1000).
		Set("user_agent", r.UserAgent()).
		Set("referer", r.Referer())

	_ = a.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "api_request",
		Properties: properties,
	})
}

// TrackReceiptProcessing tracks receipt processing with enhanced analytics
func (a *AnalyticsTracker) TrackReceiptProcessing(ctx context.Context, userID string, duration time.Duration, success bool, details map[string]interface{}) {
	if !a.enabled || a.client == nil {
		return
	}

	properties := posthog.NewProperties().
		Set("duration_ms", duration.Milliseconds()).
		Set("success", success).
		Set("processing_stage", details["stage"]).
		Set("ocr_confidence", details["ocr_confidence"]).
		Set("category_confidence", details["category_confidence"])

	// Add error details if failed
	if !success {
		if err, ok := details["error"]; ok {
			properties.Set("error_message", fmt.Sprintf("%v", err))
			properties.Set("error_type", details["error_type"])
		}
	}

	_ = a.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "receipt_processed",
		Properties: properties,
	})
}

// TrackAuthEvent tracks authentication events
func (a *AnalyticsTracker) TrackAuthEvent(userID, event string, properties map[string]interface{}) {
	if !a.enabled || a.client == nil {
		return
	}

	// Standard auth events from ADR
	validEvents := map[string]bool{
		"user_signup_started":    true,
		"user_signup_completed":  true,
		"user_login":             true,
		"user_logout":            true,
		"password_reset_requested": true,
	}

	if !validEvents[event] {
		event = "auth_" + event // Prefix non-standard events
	}

	props := posthog.NewProperties()
	for k, v := range properties {
		props.Set(k, v)
	}

	_ = a.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      event,
		Properties: props,
	})
}

// TrackFeatureUsage tracks feature adoption
func (a *AnalyticsTracker) TrackFeatureUsage(userID, feature string, firstUse bool) {
	if !a.enabled || a.client == nil {
		return
	}

	event := "feature_used"
	if firstUse {
		event = "feature_first_use"
	}

	_ = a.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      event,
		Properties: posthog.NewProperties().
			Set("feature", feature).
			Set("first_use", firstUse),
	})
}

// IdentifyUser identifies a user with traits
func (a *AnalyticsTracker) IdentifyUser(userID string, traits map[string]interface{}) {
	if !a.enabled || a.client == nil {
		return
	}

	props := posthog.NewProperties()
	for k, v := range traits {
		props.Set(k, v)
	}

	_ = a.client.Enqueue(posthog.Identify{
		DistinctId: userID,
		Properties: props,
	})
}

// AddBreadcrumb adds a breadcrumb for crash context
func (a *AnalyticsTracker) AddBreadcrumb(correlationID string, crumb Breadcrumb) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if correlationID == "" {
		correlationID = "global"
	}

	if a.breadcrumbs[correlationID] == nil {
		a.breadcrumbs[correlationID] = []Breadcrumb{}
	}

	a.breadcrumbs[correlationID] = append(a.breadcrumbs[correlationID], crumb)

	// Keep only last 50 breadcrumbs per correlation ID
	if len(a.breadcrumbs[correlationID]) > 50 {
		a.breadcrumbs[correlationID] = a.breadcrumbs[correlationID][1:]
	}
}

// cleanupBreadcrumbs runs periodically to remove old breadcrumbs
func (a *AnalyticsTracker) cleanupBreadcrumbs() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.mu.Lock()
			cutoff := time.Now().Add(-1 * time.Hour) // Keep only last hour
			
			for correlationID, breadcrumbs := range a.breadcrumbs {
				// Filter out old breadcrumbs
				filtered := make([]Breadcrumb, 0, len(breadcrumbs))
				for _, crumb := range breadcrumbs {
					if crumb.Timestamp.After(cutoff) {
						filtered = append(filtered, crumb)
					}
				}
				
				// Remove correlation ID if no breadcrumbs remain
				if len(filtered) == 0 {
					delete(a.breadcrumbs, correlationID)
				} else {
					a.breadcrumbs[correlationID] = filtered
				}
			}
			a.mu.Unlock()
			
		case <-a.stopCleanup:
			return
		}
	}
}

// retryFailedEvents processes failed events from the buffer
func (a *AnalyticsTracker) retryFailedEvents() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case event := <-a.failedEvents:
			if a.client != nil && a.enabled {
				// Retry the event
				switch e := event.(type) {
				case posthog.Capture:
					if err := a.client.Enqueue(e); err != nil && a.logger != nil {
						a.logger.Debug("Failed to retry PostHog event",
							slog.String("event", e.Event),
						)
					}
				}
			}
			
		case <-ticker.C:
			// Periodic check to process any buffered events
			
		case <-a.stopCleanup:
			return
		}
	}
}

// GetBreadcrumbs retrieves breadcrumbs for a correlation ID
func (a *AnalyticsTracker) GetBreadcrumbs(correlationID string) []Breadcrumb {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if correlationID == "" {
		correlationID = "global"
	}

	return a.breadcrumbs[correlationID]
}

// ClearBreadcrumbs clears breadcrumbs for a correlation ID
func (a *AnalyticsTracker) ClearBreadcrumbs(correlationID string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if correlationID == "" {
		correlationID = "global"
	}

	delete(a.breadcrumbs, correlationID)
}

// Close closes the analytics client
func (a *AnalyticsTracker) Close() error {
	// Stop cleanup goroutines
	close(a.stopCleanup)
	
	// Close failed events channel
	close(a.failedEvents)
	
	// Close PostHog client
	if a.client != nil {
		return a.client.Close()
	}
	return nil
}