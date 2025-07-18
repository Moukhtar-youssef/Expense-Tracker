/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"Expense_tracker/internal/config"
	"Expense_tracker/internal/storage"
	"errors"
	"fmt"
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
			err := config.InitConfig()
			if err != nil {
				if !errors.Is(err, config.NoconfigFile) {
					log.Fatal(err)
				}
				fmt.Println(err.Error())
			}
			err = storage.InitDB()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
