package generator

import (
	"embed"

	"github.com/emmajones/goboot/internal/context"
	"github.com/emmajones/goboot/internal/filesystem"
	"github.com/emmajones/goboot/internal/template"
)

//go:embed all:templates/ci
var ciTemplates embed.FS

type CIGenerator struct {
	*BaseGenerator
}

func NewCIGenerator(writer *filesystem.Writer, renderer *template.Renderer) *CIGenerator {
	return &CIGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *CIGenerator) Name() string {
	return "ci"
}

func (g *CIGenerator) Supports(projectType string) bool {
	return true
}

func (g *CIGenerator) Generate(ctx *context.ProjectContext) error {
	// For now, we only support GitHub Actions
	return g.GenerateFromTemplates(ctx.Config, ciTemplates, "templates/ci/github", ".github/workflows")
}

func init() {
	Register(NewCIGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
