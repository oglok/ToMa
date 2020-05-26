package cmd

import (
	"fmt"
	"os"

	tm "github.com/oglok/ToMa/pkg/tokenmanager"
	"github.com/spf13/cobra"
)

var exportConfigCmd = &cobra.Command{
	Use:              "export backupFile",
	Short:            `Command to export a backup file.`,
	Long:             `Command to export a backup file in JSON format.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		configFile := os.Args[2]
		if configFile == "" {
			fmt.Println("ERROR: add backup file to save existing tokens.")
		}
		tm.ExportConfigFile(configFile)
	},
}

func init() {

	rootCmd.AddCommand(exportConfigCmd)

}
