/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Amount float64

// addCmd represents the add command

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunAddExpenseCmd(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "Amount of the expense (required)")
	addCmd.MarkFlagRequired("amount")
}

func RunAddExpenseCmd(args []string) error {
	if Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}
	fmt.Println(Amount)
	return nil
}
