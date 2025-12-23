package doctor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Diagnosis struct {
	Issues   []string
	Warnings []string
}

func Check(projectDir string) (*Diagnosis, error) {
	diag := &Diagnosis{}

	// 1. Check config.yaml
	configPath := filepath.Join(projectDir, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		diag.Issues = append(diag.Issues, "config.yaml is missing")
	} else {
		data, err := os.ReadFile(configPath)
		if err != nil {
			diag.Issues = append(diag.Issues, fmt.Sprintf("failed to read config.yaml: %v", err))
		} else {
			var cfg map[string]interface{}
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				diag.Issues = append(diag.Issues, "config.yaml is not valid YAML")
			}
		}
	}

	// 2. Check go.mod
	goModPath := filepath.Join(projectDir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		diag.Issues = append(diag.Issues, "go.mod is missing")
	}

	// 3. Check Dockerfile
	dockerfilePath := filepath.Join(projectDir, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		diag.Warnings = append(diag.Warnings, "Dockerfile is missing. Use --docker when creating project or add it manually.")
	} else {
		data, err := os.ReadFile(dockerfilePath)
		if err == nil && !strings.Contains(string(data), "AS builder") {
			diag.Warnings = append(diag.Warnings, "Dockerfile is not using multi-stage builds. Consider upgrading.")
		}
	}

	// 4. Check for .goboot.yaml
	metaPath := filepath.Join(projectDir, ".goboot.yaml")
	if _, err := os.Stat(metaPath); os.IsNotExist(err) {
		diag.Warnings = append(diag.Warnings, ".goboot.yaml is missing. This project might not be fully compatible with 'goboot upgrade'.")
	}

	return diag, nil
}
