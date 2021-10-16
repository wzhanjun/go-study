package command

import (
	"github.com/spf13/cobra"
)

var CronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		StartJobs()

		return nil
	},
}
