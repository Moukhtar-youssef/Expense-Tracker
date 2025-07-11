/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export to CSV/JSON",
	Long: `Export all expenses into a CSV or JSON file for sharing or backup.
Example: expense-tracker export --format csv --output expenses.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// adding flags

	exportCmd.Flags().StringP("format", "f", "json", "Export format")
	exportCmd.Flags().StringP("output", "o", "", "File to write output to")
	exportCmd.Flags().IntP("month", "m", 0, "Export only for one month")
	exportCmd.Flags().StringP("category", "c", "", "Filter export by category")

	exportCmd.MarkFlagRequired("format")
	exportCmd.MarkFlagRequired("output")
}
