package cmd

import (
	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var getTotpCmd = &cobra.Command{
	Use:              "getTotp --name myToken",
	Short:            `Command to get Time OTP using  existing tokens.`,
	Long:             `Command to get Time OTP using  existing tokens.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		tm.GetTotp(name)
	},
}

func init() {
	getTotpCmd.Flags().StringVarP(&name, "name", "n", "", "Token name.")
	getTotpCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getTotpCmd)

}
