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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View expenses",
	Long: `List all recorded expenses optionally filtered by category or date.
Example: expense-tracker list --category "Food"`,

	Run: func(cmd *cobra.Command, args []string) {
		categoryraw, _ := cmd.Flags().GetString("category")
		category := utils.CleanStrings(categoryraw)

		fromraw, _ := cmd.Flags().GetString("from")
		from := utils.CleanStrings(fromraw)

		toraw, _ := cmd.Flags().GetString("to")
		to := utils.CleanStrings(toraw)

		uncheckedmonth, _ := cmd.Flags().GetInt("month")
		month, err := utils.ValidateMonth(uncheckedmonth)
		if err != nil {
			log.Fatal(err)
		}

		err = operation.ListExpenses(storage.DB, category, from, to, month)
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
