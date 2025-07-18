/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"Expense_tracker/internal/storage"
	"Expense_tracker/internal/utils"
	"log"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View totals",
	Long: `Show total expenses and optionally filter by month.
Example: expense-tracker summary --month 7`,
	Run: func(cmd *cobra.Command, args []string) {
		uncheckedmonth, _ := cmd.Flags().GetInt("month")
		month, err := utils.ValidateMonth(uncheckedmonth)
		if err != nil {
			log.Fatal(err)
		}

		categoryraw, _ := cmd.Flags().GetString("category")
		category := utils.CleanStrings(categoryraw)

		err = operation.SummarizeExpenses(storage.DB, category, month)
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
