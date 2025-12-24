package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/rest
var restTemplates embed.FS

// RESTGenerator handles REST project generation.
type RESTGenerator struct {
	*BaseGenerator
}

// NewRESTGenerator creates a new RESTGenerator instance.
func NewRESTGenerator(writer *filesystem.Writer, renderer *template.Renderer) *RESTGenerator {
	return &RESTGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func init() {
	Register(NewRESTGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}

func (g *RESTGenerator) Name() string {
	return "rest"
}

func (g *RESTGenerator) Supports(projectType string) bool {
	return projectType == "rest"
}

func (g *RESTGenerator) Generate(ctx *context.ProjectContext) error {
	if err := g.GenerateFromTemplates(ctx.Config, restTemplates, "templates/rest", ""); err != nil {
		return err
	}
	return g.Tidy(ctx.Config)
}
