/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"Expense_tracker/internal/storage"
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A simple CLI expense tracker",
}

func Execute() {
	err := storage.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	cobra.CheckErr(rootCmd.Execute())
}
