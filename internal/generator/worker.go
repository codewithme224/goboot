package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/worker
var workerTemplates embed.FS

type WorkerGenerator struct {
	*BaseGenerator
}

func NewWorkerGenerator(writer *filesystem.Writer, renderer *template.Renderer) *WorkerGenerator {
	return &WorkerGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *WorkerGenerator) Name() string {
	return "worker"
}

func (g *WorkerGenerator) Supports(projectType string) bool {
	return projectType == "worker"
}

func (g *WorkerGenerator) Generate(ctx *context.ProjectContext) error {
	return g.GenerateFromTemplates(ctx.Config, workerTemplates, "templates/worker", "")
}

func init() {
	Register(NewWorkerGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
