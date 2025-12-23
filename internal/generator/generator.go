package generator

import (
	"fmt"
	"io"

	"github.com/codewithme224/goboot/internal/config"
	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
)

// Generator handles the project scaffolding.
type Generator struct {
	out      io.Writer
	writer   *filesystem.Writer
	renderer *template.Renderer
}

// NewGenerator creates a new Generator instance.
func NewGenerator(out io.Writer, dryRun bool) *Generator {
	return &Generator{
		out:      out,
		writer:   filesystem.NewWriter(dryRun),
		renderer: template.NewRenderer(),
	}
}

// Generate scaffolds the project based on the configuration.
func (g *Generator) Generate(cfg *config.ProjectConfig) error {
	fmt.Fprintln(g.out, "--------------------------------------------------")
	fmt.Fprintln(g.out, "🚀 Generating Go project...")
	fmt.Fprintln(g.out, "--------------------------------------------------")
	fmt.Fprintf(g.out, "Project Name:    %s\n", cfg.Name)
	fmt.Fprintf(g.out, "Module Path:     %s\n", cfg.Module)
	fmt.Fprintf(g.out, "Project Type:    %s\n", cfg.Type)
	fmt.Fprintf(g.out, "Go Version:      %s\n", cfg.GoVersion)
	fmt.Fprintf(g.out, "Output Dir:      %s\n", cfg.Output)
	fmt.Fprintf(g.out, "Dry Run:         %v\n", cfg.DryRun)
	fmt.Fprintln(g.out, "--------------------------------------------------")

	ctx := &context.ProjectContext{
		Config:  cfg,
		RootDir: cfg.Output,
	}

	p, err := Get(string(cfg.Type))
	if err != nil {
		return fmt.Errorf("unsupported project type: %w", err)
	}

	if err := p.Generate(ctx); err != nil {
		return fmt.Errorf("failed to generate %s project: %w", cfg.Type, err)
	}

	fmt.Fprintln(g.out, "✅ Project scaffolded successfully!")
	return nil
}
