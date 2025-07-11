/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete expense",
	Long: `Delete an expense by its ID.
Example: expense-tracker delete --id 5`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// adding flags
	deleteCmd.Flags().IntP("id", "i", 0, "ID of the expense to delete")

	deleteCmd.MarkFlagRequired("id")
}
