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

//go:embed all:templates/auth
var authTemplates embed.FS

type AuthGenerator struct {
	*BaseGenerator
}

func NewAuthGenerator(writer *filesystem.Writer, renderer *template.Renderer) *AuthGenerator {
	return &AuthGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func init() {
	Register(NewAuthGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}

func (g *AuthGenerator) Name() string {
	return "auth"
}

func (g *AuthGenerator) Supports(projectType string) bool {
	return projectType == "rest"
}

func (g *AuthGenerator) Generate(ctx *context.ProjectContext) error {
	if ctx.Config.Type != "rest" {
		return fmt.Errorf("auth plugin only supports REST projects for now")
	}

	// 1. Generate Auth files
	if err := g.GenerateFromTemplates(ctx.Config, authTemplates, "templates/auth", "internal/middleware"); err != nil {
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

func (g *AuthGenerator) updateConfig(ctx *context.ProjectContext) error {
	configPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %w", err)
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	configMap["auth"] = map[string]string{
		"type":   "jwt",
		"secret": "your-secret-key",
	}

	updatedData, err := yaml.Marshal(configMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated config: %w", err)
	}

	return os.WriteFile(configPath, updatedData, 0644)
}

func (g *AuthGenerator) updateMain(ctx *context.ProjectContext) error {
	mainPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "cmd/api/main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(data)

	// Inject import
	importPath := fmt.Sprintf("\t\"%s/internal/middleware\"", ctx.Config.Module)
	content = strings.Replace(content, "// [goboot:import]", importPath+"\n\t// [goboot:import]", 1)

	// Inject config struct field
	configField := "\tAuth struct {\n\t\tType   string `yaml:\"type\"`\n\t\tSecret string `yaml:\"secret\"`\n\t} `yaml:\"auth\"`"
	content = strings.Replace(content, "// [goboot:config]", configField+"\n\t// [goboot:config]", 1)

	// Inject initialization (middleware setup)
	initCode := "\t// Initialize Auth Middleware\n\t// authMiddleware := middleware.JWTMiddleware(cfg.Auth.Secret)"
	content = strings.Replace(content, "// [goboot:init]", initCode+"\n\t// [goboot:init]", 1)

	return os.WriteFile(mainPath, []byte(content), 0644)
}
