/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"Expense_tracker/internal/storage"
	"log"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "expense-tracker",
		Short: "A simple CLI expense tracker",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// config.InitConfig()
			err := storage.InitDB()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
