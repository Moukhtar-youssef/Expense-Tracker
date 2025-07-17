/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	operation "Expense_tracker/internal/Operation"
	"log"

	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Short:   "Check system status",
	Long:    `Run a diagnostic check on your system and config to ensure everything is working.`,
	Example: `expense-tracker doctor`,
	Run: func(cmd *cobra.Command, args []string) {
		err := operation.DoctorChecking()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
