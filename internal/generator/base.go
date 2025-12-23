package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/emmajones/goboot/internal/config"
	"github.com/emmajones/goboot/internal/filesystem"
	"github.com/emmajones/goboot/internal/template"
)

// BaseGenerator provides common functionality for all generators.
type BaseGenerator struct {
	writer   *filesystem.Writer
	renderer *template.Renderer
}

// NewBaseGenerator creates a new BaseGenerator instance.
func NewBaseGenerator(writer *filesystem.Writer, renderer *template.Renderer) *BaseGenerator {
	return &BaseGenerator{
		writer:   writer,
		renderer: renderer,
	}
}

// GenerateFromTemplates walks through the templates and renders them.
func (b *BaseGenerator) GenerateFromTemplates(cfg *config.ProjectConfig, templates fs.FS, templateRoot string, targetSubPath string) error {
	return fs.WalkDir(templates, templateRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Calculate relative path from templateRoot
		relPath, err := filepath.Rel(templateRoot, path)
		if err != nil {
			return err
		}

		// Skip Dockerfile if not requested
		if relPath == "Dockerfile.tmpl" && !cfg.Docker {
			return nil
		}

		// Remove .tmpl extension for the output file
		outputPath := strings.TrimSuffix(relPath, ".tmpl")

		// Prepend output directory and target subpath
		fullOutputPath := filepath.Join(cfg.Output, cfg.Name, targetSubPath, outputPath)

		// Read template content
		tmplContent, err := fs.ReadFile(templates, path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		// Render template
		rendered, err := b.renderer.Render(relPath, string(tmplContent), cfg)
		if err != nil {
			return err
		}

		// Write rendered content to file
		if err := b.writer.WriteFile(fullOutputPath, rendered); err != nil {
			return err
		}

		// Write metadata file
		metaPath := filepath.Join(cfg.Output, cfg.Name, ".goboot.yaml")
		metaContent := fmt.Sprintf("projectType: %s\ntemplateVersion: 0.2.0\n", cfg.Type)
		_ = os.WriteFile(metaPath, []byte(metaContent), 0644)

		return nil
	})
}
