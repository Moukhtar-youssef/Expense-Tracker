/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version",
	Long:    `Display the current version of the expense-tracker CLI tool.`,
	Example: `expense-tracker version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Expense Tracker CLI version: %s \n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
