package clicommands

import (
	"os"

	"github.com/samekigor/quill-cli/cmd/clicommands/auths"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quill",
	Short: "Lightweight container system with flexible configuration management",
	Long:  `Quill is a CLI tool that allows users to manage and deploy containerized applications with ease.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auths.LoginCmd)
	rootCmd.AddCommand(auths.LogoutCmd)
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toglle")
}
