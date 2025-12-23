package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// Writer handles file system operations.
type Writer struct {
	dryRun bool
}

// NewWriter creates a new Writer instance.
func NewWriter(dryRun bool) *Writer {
	return &Writer{
		dryRun: dryRun,
	}
}

// WriteFile writes content to a file.
func (w *Writer) WriteFile(path string, content []byte) error {
	if w.dryRun {
		fmt.Printf("[DRY RUN] Would write file: %s\n", path)
		return nil
	}

	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	return os.WriteFile(path, content, 0644)
}

// CreateDir creates a directory if it doesn't exist.
func (w *Writer) CreateDir(path string) error {
	if w.dryRun {
		fmt.Printf("[DRY RUN] Would create directory: %s\n", path)
		return nil
	}

	return os.MkdirAll(path, 0755)
}
