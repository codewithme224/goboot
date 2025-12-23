package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/emmajones/goboot/internal/config"
	"github.com/emmajones/goboot/internal/context"
	"gopkg.in/yaml.v3"
)

// loadProjectContext loads the project context from the current directory.
func loadProjectContext() (*context.ProjectContext, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(cwd, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config.yaml not found. Are you in the project root?")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return nil, err
	}

	// Get project name from config or directory
	projectName := filepath.Base(cwd)
	if app, ok := configMap["app"].(map[string]interface{}); ok {
		if name, ok := app["name"].(string); ok {
			projectName = name
		}
	}

	// Get module from go.mod
	module := ""
	goModPath := filepath.Join(cwd, "go.mod")
	if data, err := os.ReadFile(goModPath); err == nil {
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
			module = strings.TrimPrefix(lines[0], "module ")
		}
	}

	cfg := &config.ProjectConfig{
		Name:   projectName,
		Module: module,
		Output: filepath.Dir(cwd),
	}

	// Infer project type
	if _, err := os.Stat(filepath.Join(cwd, "cmd/api")); err == nil {
		cfg.Type = config.TypeREST
	} else if _, err := os.Stat(filepath.Join(cwd, "cmd/grpc")); err == nil {
		cfg.Type = config.TypeGRPC
	}

	return &context.ProjectContext{
		Config:  cfg,
		RootDir: filepath.Dir(cwd),
	}, nil
}
