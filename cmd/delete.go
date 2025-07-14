/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"Expense_tracker/internal/storage"
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete expense",
	Long: `Delete an expense by its ID.
Example: expense-tracker delete --id 5`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		err := operation.DeleteExpense(storage.DB, id)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Delete all expenses",
	Long:  "Delete all expenses and reset autoincrement id",
	Run: func(cmd *cobra.Command, args []string) {
		err := operation.DeleteAllExpenses(storage.DB)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteAllCmd)

	// adding flags
	deleteCmd.Flags().IntP("id", "i", 0, "ID of the expense to delete")

	deleteCmd.MarkFlagRequired("id")
}
