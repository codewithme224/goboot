package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/test
var testTemplates embed.FS

type TestGenerator struct {
	*BaseGenerator
}

func NewTestGenerator(writer *filesystem.Writer, renderer *template.Renderer) *TestGenerator {
	return &TestGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *TestGenerator) Name() string {
	return "test"
}

func (g *TestGenerator) Supports(projectType string) bool {
	return true
}

func (g *TestGenerator) Generate(ctx *context.ProjectContext) error {
	// 1. Generate Test files
	if err := g.GenerateFromTemplates(ctx.Config, testTemplates, "templates/test/testify", "internal/service"); err != nil {
		return err
	}

	// 2. Run Tidy to download testify
	return g.Tidy(ctx.Config)
}

func init() {
	Register(NewTestGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
