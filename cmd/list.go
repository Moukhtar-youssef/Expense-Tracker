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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View expenses",
	Long: `List all recorded expenses optionally filtered by category or date.
Example: expense-tracker list --category "Food"`,

	Run: func(cmd *cobra.Command, args []string) {
		categoryraw, _ := cmd.Flags().GetString("category")
		categorytrimmed := strings.TrimSpace(categoryraw)
		category := strings.ToLower(categorytrimmed)
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		month, _ := cmd.Flags().GetInt("month")
		if month == -1 {
			log.Fatal("please provide the number of the month you wanna filter with")
		}

		err := operation.ListExpenses(storage.DB, category, from, to, month)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// adding flags

	listCmd.Flags().StringP("category", "c", "", "Filter by category")
	listCmd.Flags().String("from", "", "Start date")
	listCmd.Flags().String("to", "", "End date")
	listCmd.Flags().IntP("month", "m", -1, "Filter by month (current year)")
}
