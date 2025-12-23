package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
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
