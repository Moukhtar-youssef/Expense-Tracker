/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"Expense_tracker/internal/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new expense (wrap multi-word description in quotes)",
	Long: `Add a new expense to your tracker with details like description, amount, category, and date.
Example: expense-tracker add --description "Lunch" --amount 15 --category "Food"`,

	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetFloat64("amount")
		category, _ := cmd.Flags().GetString("category")
		date, _ := cmd.Flags().GetString("date")

		dateParsed, err := utils.ParseDate(date)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Add called with:", description, amount, category, dateParsed)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
