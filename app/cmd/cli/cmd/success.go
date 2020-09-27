package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var successCmd = &cobra.Command{
	Use:   "success",
	Short: "Output `OK!` to stdout and ends normally",
	Args:  cobra.RangeArgs(0, 0), // No arguments are accepted.
	PreRun: func(cmd *cobra.Command, args []string) {
		// TODO(istsh): Notify Slack of the start.
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		time.Sleep(10 * time.Second)

		fmt.Println("OK!")
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// TODO(istsh): Notify Slack of the finish.
	},
}
