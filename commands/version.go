package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CmdVersion)
}

var CmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Simulator Setup",
	Long:  `Does what it says on the tin!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Simulator Setup v0.1.0")
	},
}
