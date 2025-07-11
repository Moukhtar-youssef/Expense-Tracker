/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View expenses",
	Long: `List all recorded expenses optionally filtered by category or date.
Example: expense-tracker list --category "Food"`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// adding flags

	listCmd.Flags().StringP("category", "c", "", "Filter by category")
	listCmd.Flags().String("from", "", "Start date")
	listCmd.Flags().String("to", "", "End date")
	listCmd.Flags().IntP("month", "m", 1, "Filter by month (current year)")
}
