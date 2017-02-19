package commands

import (
	"strings"

	"github.com/matthewellis/simulator-setup/simulator-setup"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CmdDump)
}

var CmdDump = &cobra.Command{
	Use:   "dump",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		simulatorsetup.DumpSimData(strings.Join(args, " "))
	},
}
