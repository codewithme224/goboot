package generator

import (
	"fmt"
)

var (
	registry = make(map[string]Plugin)
)

// Register adds a plugin to the registry.
func Register(p Plugin) {
	registry[p.Name()] = p
}

// Get returns a plugin by name.
func Get(name string) (Plugin, error) {
	p, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("plugin not found: %s", name)
	}
	return p, nil
}

// List returns all registered plugins.
func List() []Plugin {
	var plugins []Plugin
	for _, p := range registry {
		plugins = append(plugins, p)
	}
	return plugins
}
