package upgrader

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const CurrentTemplateVersion = "0.2.0"

type Metadata struct {
	ProjectType     string   `yaml:"projectType"`
	EnabledFeatures []string `yaml:"enabledFeatures"`
	TemplateVersion string   `yaml:"templateVersion"`
}

func CheckUpgrade(projectDir string) (string, error) {
	metaPath := filepath.Join(projectDir, ".goboot.yaml")
	if _, err := os.Stat(metaPath); os.IsNotExist(err) {
		return "", fmt.Errorf(".goboot.yaml not found. Cannot determine version.")
	}

	data, err := os.ReadFile(metaPath)
	if err != nil {
		return "", err
	}

	var meta Metadata
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return "", err
	}

	if meta.TemplateVersion < CurrentTemplateVersion {
		return fmt.Sprintf("Upgrade available: %s -> %s", meta.TemplateVersion, CurrentTemplateVersion), nil
	}

	return "Project is up to date.", nil
}
