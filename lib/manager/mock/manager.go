package mock

import (
	"net/http"

	"github.com/Jeffail/benthos/v3/internal/component"
	"github.com/Jeffail/benthos/v3/internal/component/cache"
	"github.com/Jeffail/benthos/v3/internal/component/input"
	"github.com/Jeffail/benthos/v3/internal/component/output"
	"github.com/Jeffail/benthos/v3/internal/component/processor"
	"github.com/Jeffail/benthos/v3/internal/component/ratelimit"
	"github.com/Jeffail/benthos/v3/lib/message"
)

// Manager provides a mock benthos manager that components can use to test
// interactions with fake resources.
type Manager struct {
	Inputs     map[string]*Input
	Caches     map[string]map[string]CacheItem
	RateLimits map[string]RateLimit
	Outputs    map[string]OutputWriter
	Processors map[string]Processor
	Pipes      map[string]<-chan message.Transaction
}

// NewManager provides a new mock manager.
func NewManager() *Manager {
	return &Manager{
		Inputs:     map[string]*Input{},
		Caches:     map[string]map[string]CacheItem{},
		RateLimits: map[string]RateLimit{},
		Outputs:    map[string]OutputWriter{},
		Processors: map[string]Processor{},
		Pipes:      map[string]<-chan message.Transaction{},
	}
}

// RegisterEndpoint registers a server wide HTTP endpoint.
func (m *Manager) RegisterEndpoint(path, desc string, h http.HandlerFunc) {}

// GetCache attempts to find a service wide cache by its name.
func (m *Manager) GetCache(name string) (cache.V1, error) {
	values, ok := m.Caches[name]
	if !ok {
		return nil, component.ErrCacheNotFound
	}
	return &Cache{Values: values}, nil
}

// GetRateLimit attempts to find a service wide rate limit by its name.
func (m *Manager) GetRateLimit(name string) (ratelimit.V1, error) {
	fn, ok := m.RateLimits[name]
	if !ok {
		return nil, component.ErrRateLimitNotFound
	}
	return fn, nil
}

// GetInput attempts to find a service wide input by its name.
func (m *Manager) GetInput(name string) (input.Streamed, error) {
	if i, exists := m.Inputs[name]; exists {
		return i, nil
	}
	return nil, component.ErrInputNotFound
}

// GetProcessor attempts to find a service wide processor by its name.
func (m *Manager) GetProcessor(name string) (processor.V1, error) {
	fn, ok := m.Processors[name]
	if !ok {
		return nil, component.ErrProcessorNotFound
	}
	return fn, nil
}

// GetOutput attempts to find a service wide output by its name.
func (m *Manager) GetOutput(name string) (output.Sync, error) {
	if o, exists := m.Outputs[name]; exists {
		return o, nil
	}
	return nil, component.ErrOutputNotFound
}

// GetPipe attempts to find a service wide transaction chan by its name.
func (m *Manager) GetPipe(name string) (<-chan message.Transaction, error) {
	if p, ok := m.Pipes[name]; ok {
		return p, nil
	}
	return nil, component.ErrPipeNotFound
}

// SetPipe registers a transaction chan under a name.
func (m *Manager) SetPipe(name string, t <-chan message.Transaction) {
	m.Pipes[name] = t
}

// UnsetPipe removes a named transaction chan.
func (m *Manager) UnsetPipe(name string, t <-chan message.Transaction) {
	delete(m.Pipes, name)
}
