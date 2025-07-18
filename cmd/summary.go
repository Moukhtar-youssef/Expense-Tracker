/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"Expense_tracker/internal/storage"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View totals",
	Long: `Show total expenses and optionally filter by month.
Example: expense-tracker summary --month 7`,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")

		categoryraw, _ := cmd.Flags().GetString("category")
		categorytrimmed := strings.TrimSpace(categoryraw)
		category := strings.ToLower(categorytrimmed)

		err := operation.SummarizeExpenses(storage.DB, category, month)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	// adding flags

	summaryCmd.Flags().IntP("month", "m", -1, "Show summary for a given month")
	summaryCmd.Flags().StringP("category", "c", "", "Filter summary by caregory")
}
