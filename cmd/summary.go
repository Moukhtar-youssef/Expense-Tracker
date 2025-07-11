/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View totals",
	Long: `Show total expenses and optionally filter by month.
Example: expense-tracker summary --month 7`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("summary called")
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	// adding flags

	summaryCmd.Flags().IntP("month", "m", 1, "Show summary for a given month")
	summaryCmd.Flags().StringP("category", "c", "", "Filter summary by caregory")
}
