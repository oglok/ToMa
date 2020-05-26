package cmd

import (
	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var getHotpCmd = &cobra.Command{
	Use:              "getHotp --name myToken",
	Short:            `Command to get HMAC OTP using  existing tokens.`,
	Long:             `Command to get HMAC OTP using  existing tokens.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		tm.GetHotp(name)
	},
}

func init() {
	getHotpCmd.Flags().StringVarP(&name, "name", "n", "", "Token name.")
	getHotpCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getHotpCmd)

}
