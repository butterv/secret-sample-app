package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var failCmd = &cobra.Command{
	Use:   "fail",
	Short: "Output `NG!` to stdout and return an error",
	Args:  cobra.RangeArgs(0, 0), // No arguments are accepted.
	PreRun: func(cmd *cobra.Command, args []string) {
		// TODO(istsh): Notify Slack of the start.
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		time.Sleep(10 * time.Second)

		fmt.Println("NG!")
		return errors.New("an error occurred")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// TODO(istsh): Notify Slack of the finish.
	},
}
