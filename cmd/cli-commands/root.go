package clicommands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quill",
	Short: "Lightweight container system with flexible configuration management",
	Long:  `Quill is a CLI tool that allows users to manage and deploy containerized applications with ease.`,
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
