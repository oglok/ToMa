package cmd

import (
	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var deleteTokenCmd = &cobra.Command{
	Use:              "delete --name myToken",
	Short:            `Command to delete an existing token.`,
	Long:             `Command to delete an existing token by name`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		tm.DeleteToken(name)
	},
}

func init() {

	deleteTokenCmd.Flags().StringVarP(&name, "name", "n", "", "Token name.")
	deleteTokenCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(deleteTokenCmd)

}
