/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update existing expense",
	Long: `Update details of an existing expense by ID.
Example: expense-tracker update --id 3 --amount 25 --description "Dinner"`,

	Run: func(cmd *cobra.Command, args []string) {
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
