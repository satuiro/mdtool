package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mdtool",
		Short: "A markdown preview and generation tool",
		Long:  `A CLI tool to preview markdown files with beautiful formatting and generate READMEs for git repositories.`,
	}

	rootCmd.AddCommand(previewCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func previewCmd() *cobra.Command {
	var (
		theme string
		raw   bool
	)

	cmd := &cobra.Command{
		Use:   "preview [file]",
		Short: "Preview a markdown file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return previewMarkdown(args[0], theme, raw)
		},
	}

	cmd.Flags().StringVarP(&theme, "theme", "t", "dark", "Theme to use (dark/light)")
	cmd.Flags().BoolVarP(&raw, "raw", "r", false, "Show raw markdown")

	return cmd
}
