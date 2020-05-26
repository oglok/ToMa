package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ToMa",
	Short: "ToMa is a light-weight token manager to create soft tokens and get One-Time-Passwords",
	Long: `Token Manager (ToMa) is a light-weight application to generate One-Time-Passwords (OTP) as a mechanism
		   for second factor authentication. You can create, store, backup and delete multiple tokens and
		   generate OTPs.`,
}

// Execute is...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
