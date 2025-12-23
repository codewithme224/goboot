package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/emmajones/goboot/internal/context"
	"github.com/emmajones/goboot/internal/filesystem"
	"github.com/emmajones/goboot/internal/template"
	"gopkg.in/yaml.v3"
)

//go:embed all:templates/observability
var observabilityTemplates embed.FS

type ObservabilityGenerator struct {
	*BaseGenerator
}

func NewObservabilityGenerator(writer *filesystem.Writer, renderer *template.Renderer) *ObservabilityGenerator {
	return &ObservabilityGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func (g *ObservabilityGenerator) Name() string {
	return "observability"
}

func (g *ObservabilityGenerator) Supports(projectType string) bool {
	return true
}

func (g *ObservabilityGenerator) Generate(ctx *context.ProjectContext) error {
	// 1. Generate Observability files
	if err := g.GenerateFromTemplates(ctx.Config, observabilityTemplates, "templates/observability", "internal/observability"); err != nil {
		return err
	}

	// 2. Update config.yaml
	if err := g.updateConfig(ctx); err != nil {
		return err
	}

	// 3. Update main.go
	if err := g.updateMain(ctx); err != nil {
		return err
	}

	return nil
}

func (g *ObservabilityGenerator) updateConfig(ctx *context.ProjectContext) error {
	configPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %w", err)
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	configMap["observability"] = map[string]interface{}{
		"enabled": false,
		"metrics": map[string]interface{}{
			"enabled": true,
			"port":    9090,
		},
		"tracing": map[string]interface{}{
			"enabled": false,
		},
	}

	updatedData, err := yaml.Marshal(configMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated config: %w", err)
	}

	return os.WriteFile(configPath, updatedData, 0644)
}

func (g *ObservabilityGenerator) updateMain(ctx *context.ProjectContext) error {
	mainPath := ""
	if ctx.Config.Type == "rest" {
		mainPath = filepath.Join(ctx.RootDir, ctx.Config.Name, "cmd/api/main.go")
	} else {
		mainPath = filepath.Join(ctx.RootDir, ctx.Config.Name, "cmd/grpc/main.go")
	}

	data, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(data)

	// Inject import
	importPath := fmt.Sprintf("\t\"%s/internal/observability\"", ctx.Config.Module)
	content = strings.Replace(content, "// [goboot:import]", importPath+"\n\t// [goboot:import]", 1)

	// Inject config struct field
	configField := "\tObservability struct {\n\t\tEnabled bool `yaml:\"enabled\"`\n\t\tMetrics struct {\n\t\t\tEnabled bool `yaml:\"enabled\"`\n\t\t\tPort    int  `yaml:\"port\"`\n\t\t} `yaml:\"metrics\"`\n\t\tTracing struct {\n\t\t\tEnabled bool `yaml:\"enabled\"`\n\t\t} `yaml:\"tracing\"`\n\t} `yaml:\"observability\"`"
	content = strings.Replace(content, "// [goboot:config]", configField+"\n\t// [goboot:config]", 1)

	// Inject initialization
	initCode := "\t// Initialize Observability\n\t// if cfg.Observability.Enabled {\n\t// \tcleanup, err := observability.Setup(cfg.App.Name, cfg.Observability.Metrics.Port)\n\t// \tif err != nil { log.Fatalf(\"Failed to setup observability: %v\", err) }\n\t// \tdefer cleanup()\n\t// }"
	content = strings.Replace(content, "// [goboot:init]", initCode+"\n\t// [goboot:init]", 1)

	return os.WriteFile(mainPath, []byte(content), 0644)
}

func init() {
	Register(NewObservabilityGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}
