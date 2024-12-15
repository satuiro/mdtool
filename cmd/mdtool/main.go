package main

import (
	"fmt"
	"os"

	"github.com/satuiro/mdtool/internal/config"
	"github.com/satuiro/mdtool/internal/readme"
	"github.com/spf13/cobra"
)

var version = "v0.1.1"

var rootCmd = &cobra.Command{
	Use:     "mdtool",
	Version: version,
	Short:   "A CLI tool to generate and render README for a GitHub project",
	Long:    `mdtool is a comprehensive tool that helps you generate and manage README files for GitHub repositories.`,
}

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Generate README for the repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoName, _ := cmd.Flags().GetString("repo")
		output, _ := cmd.Flags().GetString("output")
		save, _ := cmd.Flags().GetBool("save")

		generator := readme.NewGenerator(
			config.GetConfig(),
			repoName,
		)

		content, err := generator.Generate()
		if err != nil {
			return fmt.Errorf("failed to generate README: %w", err)
		}

		if content != "" {
			generator.DisplayReadme(content, output)

			if save {
				if err := os.WriteFile("README.md", []byte(content), 0644); err != nil {
					return fmt.Errorf("failed to save README: %w", err)
				}
				fmt.Println("\nREADME.md saved successfully!")
			}
		}

		return nil
	},
}

func init() {
	readmeCmd.Flags().StringP("repo", "r", "", "Repository name in format 'owner/repo'")
	readmeCmd.Flags().StringP("output", "o", "preview", "Output format (preview/raw)")
	readmeCmd.Flags().BoolP("save", "s", false, "Save README to file")
	readmeCmd.MarkFlagRequired("repo")

	rootCmd.AddCommand(readmeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
