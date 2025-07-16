/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var currency string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Setup config and data file",
	Long:    `Initialize the expense tracker configuration and data store.`,
	Example: `expense-tracker init`,
	Run: func(cmd *cobra.Command, args []string) {
		form := huh.NewForm(huh.NewGroup(
			huh.NewInput().Title("What currency do you wanna use").Value(&currency),
		))
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(currency)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
