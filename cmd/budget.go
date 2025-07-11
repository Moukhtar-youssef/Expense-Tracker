/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// budgetCmd represents the parent budget command
var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "Manage your monthly budget",
	Long: `The 'budget' command allows you to manage monthly budgets.

You can set a new budget for a month or check if your spending has exceeded it.`,
}

// setcmd represents the 'budget set' command
var setcmd = &cobra.Command{
	Use:   "set",
	Short: "Set monthly budget",
	Long: `Set a monthly budget limit.

Example:
  expense-tracker budget set --month 7 --amount 500`,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")
		amount, _ := cmd.Flags().GetFloat64("amount")
		fmt.Printf("Set budget: Month %d, Amount %.2f\n", month, amount)
	},
}

// checkcmd represents the 'budget check' command
var checkcmd = &cobra.Command{
	Use:   "check",
	Short: "Check if budget is exceeded",
	Long: `Check whether your spending has exceeded the set monthly budget.

Example:
  expense-tracker budget check --month 7`,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")
		fmt.Printf("Checking budget for Month %d...\n", month)
	},
}

func init() {
	rootCmd.AddCommand(budgetCmd)

	// Subcommands
	budgetCmd.AddCommand(setcmd)
	budgetCmd.AddCommand(checkcmd)

	// Flags for both subcommands
	setcmd.Flags().IntP("month", "m", 1, "Target month (1-12)")
	setcmd.Flags().Float64P("amount", "a", 0, "Budget amount")
	setcmd.MarkFlagRequired("month")
	setcmd.MarkFlagRequired("amount")

	checkcmd.Flags().IntP("month", "m", 1, "Target month (1-12)")
	checkcmd.MarkFlagRequired("month")
}
