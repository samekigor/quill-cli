package tools

import (
	"github.com/spf13/cobra"
)

type FlagDetails struct {
	Name        string
	ShortName   string
	Default     interface{}
	Description string
}

type CommandDetails struct {
	Use   string
	Short string
	Long  string
	Flags []FlagDetails
	Run   func(cmd *cobra.Command, args []string)
}
