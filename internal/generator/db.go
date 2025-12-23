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

//go:embed all:templates/db
var dbTemplates embed.FS

type DBGenerator struct {
	*BaseGenerator
}

func NewDBGenerator(writer *filesystem.Writer, renderer *template.Renderer) *DBGenerator {
	return &DBGenerator{
		BaseGenerator: NewBaseGenerator(writer, renderer),
	}
}

func init() {
	Register(NewDBGenerator(filesystem.NewWriter(false), template.NewRenderer()))
}

func (g *DBGenerator) Name() string {
	return "db"
}

func (g *DBGenerator) Supports(projectType string) bool {
	return true // Supports both rest and grpc
}

func (g *DBGenerator) Generate(ctx *context.ProjectContext) error {
	dbType := ctx.Config.DB
	if dbType == "none" || dbType == "" {
		return fmt.Errorf("database type not specified")
	}

	templatePath := fmt.Sprintf("templates/db/%s", dbType)

	// 1. Generate DB files
	if err := g.GenerateFromTemplates(ctx.Config, dbTemplates, templatePath, "internal/db"); err != nil {
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

func (g *DBGenerator) updateConfig(ctx *context.ProjectContext) error {
	configPath := filepath.Join(ctx.RootDir, ctx.Config.Name, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %w", err)
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	dbConfig := map[string]string{
		"driver": string(ctx.Config.DB),
	}
	if ctx.Config.DB == "postgres" {
		dbConfig["dsn"] = "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
	} else if ctx.Config.DB == "mysql" {
		dbConfig["dsn"] = "user:pass@tcp(localhost:3306)/dbname"
	} else {
		dbConfig["uri"] = "mongodb://localhost:27017"
	}

	configMap["database"] = dbConfig

	updatedData, err := yaml.Marshal(configMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated config: %w", err)
	}

	return os.WriteFile(configPath, updatedData, 0644)
}

func (g *DBGenerator) updateMain(ctx *context.ProjectContext) error {
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
	importPath := fmt.Sprintf("\t\"%s/internal/db\"", ctx.Config.Module)
	content = strings.Replace(content, "// [goboot:import]", importPath+"\n\t// [goboot:import]", 1)

	// Inject config struct field
	configField := "\tDatabase struct {\n\t\tDriver string `yaml:\"driver\"`"
	if ctx.Config.DB == "postgres" || ctx.Config.DB == "mysql" {
		configField += "\n\t\tDSN    string `yaml:\"dsn\"`"
	} else {
		configField += "\n\t\tURI    string `yaml:\"uri\"`"
	}
	configField += "\n\t} `yaml:\"database\"`"

	content = strings.Replace(content, "// [goboot:config]", configField+"\n\t// [goboot:config]", 1)

	// Inject initialization
	initCode := ""
	if ctx.Config.DB == "postgres" {
		initCode = "\t// Initialize Postgres\n\t// db, err := db.NewPostgresDB(cfg.Database.DSN)\n\t// if err != nil { log.Fatalf(\"Failed to connect to DB: %v\", err) }"
	} else if ctx.Config.DB == "mysql" {
		initCode = "\t// Initialize MySQL\n\t// db, err := db.NewMySQLDB(cfg.Database.DSN)\n\t// if err != nil { log.Fatalf(\"Failed to connect to DB: %v\", err) }"
	} else {
		initCode = "\t// Initialize Mongo\n\t// db, err := db.NewMongoDB(cfg.Database.URI)\n\t// if err != nil { log.Fatalf(\"Failed to connect to DB: %v\", err) }"
	}

	content = strings.Replace(content, "// [goboot:init]", initCode+"\n\t// [goboot:init]", 1)

	return os.WriteFile(mainPath, []byte(content), 0644)
}
