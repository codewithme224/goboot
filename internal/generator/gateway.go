package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codewithme224/goboot/internal/context"
	"github.com/codewithme224/goboot/internal/filesystem"
	"github.com/codewithme224/goboot/internal/template"
	"gopkg.in/yaml.v3"
)

//go:embed all:templates/gateway
var gatewayTemplates embed.FS

type GatewayGenerator struct {
	*BaseGenerator
}

func NewGatewayGenerator(writer *filesystem.Writer, renderer *template.Renderer) *GatewayGenerator {
	return &GatewayGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func init() {
	Register(NewGatewayGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}

func (g *GatewayGenerator) Name() string {
	return "gateway"
}

func (g *GatewayGenerator) Supports(projectType string) bool {
	return projectType == "grpc"
}

func (g *GatewayGenerator) Generate(ctx *context.ProjectContext) error {
	if ctx.Config.Type != "grpc" {
		return fmt.Errorf("gateway plugin only supports gRPC projects")
	}

	// 1. Generate Gateway files
	if err := g.GenerateFromTemplates(ctx.Config, gatewayTemplates, "templates/gateway", "internal/gateway"); err != nil {
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

	return g.Tidy(ctx.Config)
}

func (g *GatewayGenerator) updateConfig(ctx *context.ProjectContext) error {
	configPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %w", err)
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	configMap["gateway"] = map[string]interface{}{
		"enabled": true,
		"port":    8081,
	}

	updatedData, err := yaml.Marshal(configMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated config: %w", err)
	}

	return os.WriteFile(configPath, updatedData, 0644)
}

func (g *GatewayGenerator) updateMain(ctx *context.ProjectContext) error {
	mainPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "cmd/grpc/main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(data)

	// Inject import
	importPath := fmt.Sprintf("\t\"%s/internal/gateway\"", ctx.Config.Module)
	content = strings.Replace(content, "// [goboot:import]", importPath+"\n\t// [goboot:import]", 1)

	// Inject config struct field
	configField := "\tGateway struct {\n\t\tEnabled bool `yaml:\"enabled\"`\n\t\tPort    int  `yaml:\"port\"`\n\t} `yaml:\"gateway\"`"
	content = strings.Replace(content, "// [goboot:config]", configField+"\n\t// [goboot:config]", 1)

	// Inject initialization (gateway run)
	initCode := "\t// Start Gateway\n\t// if cfg.Gateway.Enabled {\n\t// \tgo gateway.Run(fmt.Sprintf(\":%d\", cfg.Server.Port), cfg.Gateway.Port)\n\t// }"
	content = strings.Replace(content, "// [goboot:init]", initCode+"\n\t// [goboot:init]", 1)

	return os.WriteFile(mainPath, []byte(content), 0644)
}
