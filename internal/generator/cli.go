package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/cli
var cliTemplates embed.FS

type CLIGenerator struct {
	*BaseGenerator
}

func NewCLIGenerator(writer *filesystem.Writer, renderer *template.Renderer) *CLIGenerator {
	return &CLIGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *CLIGenerator) Name() string {
	return "cli"
}

func (g *CLIGenerator) Supports(projectType string) bool {
	return projectType == "cli"
}

func (g *CLIGenerator) Generate(ctx *context.ProjectContext) error {
	return g.GenerateFromTemplates(ctx.Config, cliTemplates, "templates/cli", "")
}

func init() {
	Register(NewCLIGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
