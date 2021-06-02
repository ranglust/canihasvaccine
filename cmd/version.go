package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VERSION = "DEV"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", VERSION)
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
