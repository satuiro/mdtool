package render

import (
	"github.com/charmbracelet/glamour"
)

type RenderOptions struct {
	Theme string
	Raw   bool
}

func RenderMarkdown(input string, opts RenderOptions) (string, error) {
	if opts.Raw {
		return input, nil
	}

	// Use built-in styles instead of manual configuration
	var style string
	if opts.Theme == "light" {
		style = "light"
	} else {
		style = "dark"
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle(style),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return "", err
	}

	return r.Render(input)
}
