package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/emmajones/goboot/internal/config"
)

// Renderer handles template rendering.
type Renderer struct {
	funcs template.FuncMap
}

// NewRenderer creates a new Renderer instance.
func NewRenderer() *Renderer {
	return &Renderer{
		funcs: template.FuncMap{
			// Add custom template functions here if needed in the future
		},
	}
}

// Render renders a template string with the provided config.
func (r *Renderer) Render(name string, tmplStr string, cfg *config.ProjectConfig) ([]byte, error) {
	tmpl, err := template.New(name).Funcs(r.funcs).Parse(tmplStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %w", name, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return nil, fmt.Errorf("failed to execute template %s: %w", name, err)
	}

	return buf.Bytes(), nil
}
