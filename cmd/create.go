package cmd

import (
	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var name string
var digits int
var seedLen int
var hashF string

var createTokenCmd = &cobra.Command{
	Use:              "create --name myToken --seed-length 16 --digits 6 --hash-function sha256",
	Short:            `Command to create a new token.`,
	Long:             `Command to create a new token using a set of parameters such name, number of digits and hash function`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		tm.CreateToken(name, digits, seedLen, hashF)
	},
}

func init() {

	createTokenCmd.Flags().StringVarP(&name, "name", "n", "", "Token name.")
	createTokenCmd.MarkFlagRequired("name")
	createTokenCmd.Flags().IntVarP(&seedLen, "seed-length", "s", 32, "Length of the seed use for generating the OTP.")
	createTokenCmd.Flags().IntVarP(&digits, "digits", "d", 6, "Number of digits of the generated One-Time-Password")
	createTokenCmd.Flags().StringVarP(&hashF, "hash-function", "f", "sha256", "Hash function used to get the OTP. Values are: sha1, sha256, sha512")

	rootCmd.AddCommand(createTokenCmd)

}
