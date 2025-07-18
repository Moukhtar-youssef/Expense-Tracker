/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"Expense_tracker/internal/storage"
	"Expense_tracker/internal/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export to CSV/JSON",
	Long: `Export all expenses into a CSV or JSON file for sharing or backup.
Example: expense-tracker export --format csv --output expenses.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		formatraw, _ := cmd.Flags().GetString("format")
		format := utils.CleanStrings(formatraw)

		outputraw, _ := cmd.Flags().GetString("output")
		output := utils.CleanStrings(outputraw)

		categoryraw, _ := cmd.Flags().GetString("category")
		category := utils.CleanStrings(categoryraw)

		uncheckedmonth, _ := cmd.Flags().GetInt("month")
		month, err := utils.ValidateMonth(uncheckedmonth)
		if err != nil {
			log.Fatal(err)
		}

		err = operation.ExportExpenses(storage.DB, format, output, category, month)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Exported")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// adding flags

	exportCmd.Flags().StringP("format", "f", "json", "Export format")
	exportCmd.Flags().StringP("output", "o", "", "File to write output to")
	exportCmd.Flags().IntP("month", "m", -1, "Export only for one month")
	exportCmd.Flags().StringP("category", "c", "", "Filter export by category")

	exportCmd.MarkFlagRequired("format")
	exportCmd.MarkFlagRequired("output")
}
