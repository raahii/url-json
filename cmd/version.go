package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of url-json",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(resultWriter, "%s (rev: %s)\n", Version, Revision)
	},
}
