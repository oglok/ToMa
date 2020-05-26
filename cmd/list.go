package cmd

import (
	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var listTokenCmd = &cobra.Command{
	Use:              "list",
	Short:            `Command to list existing tokens.`,
	Long:             `Command to list all existing tokens.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		tm.ListTokens()
	},
}

func init() {
	rootCmd.AddCommand(listTokenCmd)

}
