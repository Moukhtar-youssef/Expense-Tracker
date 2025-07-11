/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system status",
	Long: `Run a diagnostic check on your system and config to ensure everything is working.
Example: expense-tracker doctor`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doctor called")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
