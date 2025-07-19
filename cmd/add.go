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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new expense (wrap multi-word description in quotes)",
	Long:    `Add a new expense to your tracker with details like description, amount, category, and date.`,
	Example: `expense-tracker add --description "Lunch" --amount 15 --category "Food"`,

	Run: func(cmd *cobra.Command, args []string) {
		descriptionraw, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}
		description := utils.CleanStrings(descriptionraw)

		amountraw, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			log.Fatal(err)
		}
		amount, err := utils.ValidateAmount(amountraw)
		if err != nil {
			log.Fatal(err)
		}

		categoryraw, err := cmd.Flags().GetString("category")
		if err != nil {
			log.Fatal(err)
		}
		category := utils.CleanStrings(categoryraw)

		dateraw, err := cmd.Flags().GetString("date")
		if err != nil {
			log.Fatal(err)
		}
		date := utils.CleanStrings(dateraw)

		dateParsed, err := utils.ParseDate(date)
		if err != nil {
			log.Fatal(err)
		}

		if category == "" {
			category = "uncategorized"
		}

		err = operation.AddExepnse(storage.DB, dateParsed, amount, category, description)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// adding flags

	addCmd.Flags().StringP("description", "d", "", "What the expense was for")
	addCmd.Flags().Float64P("amount", "a", 0, "How much the expense cost")
	addCmd.Flags().StringP("category", "c", "uncategorized", "Category like Food, Transport")
	addCmd.Flags().String("date", "", "Defaults to today")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
