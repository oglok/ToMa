package cmd

import (
	"fmt"
	"os"

	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var importConfigCmd = &cobra.Command{
	Use:              "import",
	Short:            `Command to import a backup file.`,
	Long:             `Command to import a backup file in JSON format.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) < 3 {
			fmt.Println("ERROR: add path to file to be imported.")
			os.Exit(1)
		}
		configFile := os.Args[2]
		tm.ImportConfigFile(configFile)
	},
}

func init() {

	rootCmd.AddCommand(importConfigCmd)

}
