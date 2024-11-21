package main

import (
	"fmt"
	"satuiro/mdtool/internal/parser"
	"satuiro/mdtool/internal/render"
)

func previewMarkdown(filePath string, theme string, raw bool) error {
	// Read the markdown file
	content, err := parser.ReadMarkdownFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Render the markdown
	rendered, err := render.RenderMarkdown(content, render.RenderOptions{
		Theme: theme,
		Raw:   raw,
	})
	if err != nil {
		return fmt.Errorf("failed to render markdown: %w", err)
	}

	// Print the rendered output
	fmt.Print(rendered)
	return nil
}
