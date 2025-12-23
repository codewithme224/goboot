package generator

import (
	"github.com/emmajones/goboot/internal/context"
)

// Plugin defines the interface for a generator feature.
type Plugin interface {
	Name() string
	Supports(projectType string) bool
	Generate(ctx *context.ProjectContext) error
}
