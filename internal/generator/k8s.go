package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/k8s
var k8sTemplates embed.FS

type K8sGenerator struct {
	*BaseGenerator
}

func NewK8sGenerator(writer *filesystem.Writer, renderer *template.Renderer) *K8sGenerator {
	return &K8sGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *K8sGenerator) Name() string {
	return "k8s"
}

func (g *K8sGenerator) Supports(projectType string) bool {
	return true
}

func (g *K8sGenerator) Generate(ctx *context.ProjectContext) error {
	return g.GenerateFromTemplates(ctx.Config, k8sTemplates, "templates/k8s", "k8s")
}

func init() {
	Register(NewK8sGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
