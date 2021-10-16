package command

import (
	"github.com/spf13/cobra"
)

func RunCommand() error {
	var rootCmd = &cobra.Command{
		Use:   "stock",
		Short: "stock 命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	addCommand(rootCmd)

	return rootCmd.Execute()
}

func addCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(QuoteCommand, CronCommand)
}
