package cmd

import (
	"fmt"
	"github.com/Bnei-Baruch/jukfs/pkg/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of jukfs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
