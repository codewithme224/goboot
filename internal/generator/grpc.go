package generator

import (
	"embed"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

//go:embed all:templates/grpc
var grpcTemplates embed.FS

// GRPCGenerator handles gRPC project generation.
type GRPCGenerator struct {
	*BaseGenerator
}

// NewGRPCGenerator creates a new GRPCGenerator instance.
func NewGRPCGenerator(writer *filesystem.Writer, renderer *template.Renderer) *GRPCGenerator {
	return &GRPCGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func init() {
	Register(NewGRPCGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}

func (g *GRPCGenerator) Name() string {
	return "grpc"
}

func (g *GRPCGenerator) Supports(projectType string) bool {
	return projectType == "grpc"
}

func (g *GRPCGenerator) Generate(ctx *context.ProjectContext) error {
	if err := g.GenerateFromTemplates(ctx.Config, grpcTemplates, "templates/grpc", ""); err != nil {
		return err
	}
	return g.Tidy(ctx.Config)
}
