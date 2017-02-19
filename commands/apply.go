package commands

import (
	"strings"

	"github.com/matthewellis/simulator-setup/simulator-setup"
	"github.com/spf13/cobra"
)

var version string

func init() {
	RootCmd.AddCommand(CmdApply)

	CmdApply.PersistentFlags().StringVar(&version, "version", "all", "iOS version to apply dump to")
}

var CmdApply = &cobra.Command{
	Use:   "apply",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		simulatorsetup.ApplyDataToSim(strings.Join(args, " "), version)
	},
}
