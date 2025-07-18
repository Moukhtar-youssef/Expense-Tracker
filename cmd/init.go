/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Setup config and data file",
	Long:    `Initialize the expense tracker configuration and data store.`,
	Example: `expense-tracker init`,
	Run: func(cmd *cobra.Command, args []string) {
		operation.InitCommand()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
