package pipeline

import "fmt"

// Registry maps stage names to factory functions.
type Registry struct {
	factories map[string]func() Stage
}

// NewRegistry creates an empty registry.
func NewRegistry() *Registry {
	return &Registry{factories: map[string]func() Stage{}}
}

// Register adds a stage factory. Panics on duplicate name.
func (r *Registry) Register(name string, factory func() Stage) {
	if _, exists := r.factories[name]; exists {
		panic(fmt.Sprintf("stage %q already registered", name))
	}
	r.factories[name] = factory
}

// MustRegister is like Register but returns the registry for chaining.
func (r *Registry) MustRegister(name string, factory func() Stage) *Registry {
	r.Register(name, factory)
	return r
}

// Build instantiates the named stages in the requested order.
func (r *Registry) Build(names []string) ([]Stage, error) {
	var stages []Stage
	for _, name := range names {
		factory, ok := r.factories[name]
		if !ok {
			return nil, fmt.Errorf("unknown stage %q", name)
		}
		stages = append(stages, factory())
	}
	return stages, nil
}

// Names returns all registered stage names.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.factories))
	for name := range r.factories {
		names = append(names, name)
	}
	return names
}
