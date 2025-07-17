/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"Expense_tracker/internal/storage"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update existing expense",
	Long:    `Update details of an existing expense by ID.`,
	Example: `expense-tracker update --id 3 --amount 25 --description "Dinner"`,

	Run: func(cmd *cobra.Command, args []string) {
		descriptionraw, _ := cmd.Flags().GetString("description")
		descriptiontrimmed := strings.TrimSpace(descriptionraw)
		description := strings.ToLower(descriptiontrimmed)

		amount, _ := cmd.Flags().GetFloat64("amount")

		categoryraw, _ := cmd.Flags().GetString("category")
		categorytrimmed := strings.TrimSpace(categoryraw)
		category := strings.ToLower(categorytrimmed)

		date, _ := cmd.Flags().GetString("date")

		id, _ := cmd.Flags().GetInt("id")

		err := operation.UpdateExpense(storage.DB, id, date, amount, category, description)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// adding flags

	updateCmd.Flags().IntP("id", "i", 0, "ID of the expense to update")
	updateCmd.Flags().StringP("description", "d", "", "New description")
	updateCmd.Flags().StringP("category", "c", "", "New category")
	updateCmd.Flags().String("date", "", "New date")
	updateCmd.Flags().Float64P("amount", "a", 0, "New amount")

	updateCmd.MarkFlagRequired("id")
}
