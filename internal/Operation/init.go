package operation

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

var (
	currency        string
	currency_symbol string
	locale          string
	date_format     string
	timezone        string

	default_category  string
	custom_categories string

	color_output        bool
	show_totals_on_list bool

	enable_budgeting         bool
	warn_if_budget_exceeded  bool
	budget_warning_threshold string

	auto_backup      bool
	backup_location  string
	backup_frequency string
)

func InitCommand() error {
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().Title("What is your default currency code? (e.g., USD, EUR)").Value(&currency),
		huh.NewInput().Title("What symbol should represent your currency? (e.g., $, €, ¥)").Value(&currency_symbol),
		huh.NewInput().Title("What is your locale? (e.g., en_US, fr_FR)").Value(&locale),
		huh.NewInput().Title("What date format do you prefer? (e.g., 2006-01-02)").Value(&date_format),
		huh.NewInput().Title("What timezone should be used? (e.g., Local, UTC, America/New_York)").Value(&timezone),
	), huh.NewGroup(
		huh.NewInput().Title("What is your default expense category?").Value(&default_category),
		huh.NewInput().Title("List your custom categories (comma-separated, e.g., food, transport, rent):").Value(&custom_categories),
	), huh.NewGroup(
		huh.NewConfirm().Title("Do you want colorful output in the terminal? (true/false)").Affirmative("yes").Negative("no").Value(&color_output),
		huh.NewConfirm().Title("Show total expenses when listing? (true/false)").Affirmative("yes").Negative("no").Value(&show_totals_on_list),
	), huh.NewGroup(
		huh.NewConfirm().Title("Enable budgeting features? (true/false)").Affirmative("yes").Negative("no").Value(&enable_budgeting),
		huh.NewConfirm().Title("Warn if budget is exceeded? (true/false)").Affirmative("yes").Negative("no").Value(&warn_if_budget_exceeded),
		huh.NewInput().Title("At what percentage should a warning appear? (e.g., 90 for 90%)").Value(&budget_warning_threshold),
	), huh.NewGroup(
		huh.NewConfirm().Title("Enable automatic backups? (true/false)").Affirmative("yes").Negative("no").Value(&auto_backup),
		huh.NewInput().Title("Where should backups be stored? (e.g., ./backups)").Value(&backup_location),
		huh.NewSelect[string]().Options(huh.NewOption("none", "none"), huh.NewOption("daily", "daily"), huh.NewOption("weekly", "weekly"), huh.NewOption("monthly", "monthly")).Title("How often should backups run? (none, daily, weekly, month)").Value(&backup_frequency),
	))
	err := form.Run()
	if err != nil {
		return fmt.Errorf("Error running the huh init form: %w", err)
	}
	return nil
}
