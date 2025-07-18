package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Currency       string `mapstructure:"currency"`
	CurrencySymbol string `mapstructure:"currency_symbol"`
	Locale         string `mapstructure:"locale"`
	DateFormat     string `mapstructure:"date_format"`
	Timezone       string `mapstructure:"timezone"`

	DefaultCategory  string   `mapstructure:"default_category"`
	CustomCategories []string `mapstructure:"custom_categories"`

	ColorOutput      bool `mapstructure:"color_output"`
	ShowTotalsOnList bool `mapstructure:"show_totals_on_list"`

	EnableBudgeting        bool `mapstructure:"enable_budgeting"`
	WarnIfBudgetExceeded   bool `mapstructure:"warn_if_budget_exceeded"`
	BudgetWarningThreshold int  `mapstructure:"budget_warning_threshold"`

	AutoBackup      bool   `mapstructure:"auto_backup"`
	BackupLocation  string `mapstructure:"backup_location"`
	BackupFrequency string `mapstructure:"backup_frequency"`
}

var AppConfig Config

var NoconfigFile error = errors.New("No config file found. Using defaults , use init commands to create config file")

func InitConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Can't find home directory: %w", err)
	}

	viper.AddConfigPath(home + "/.expense-tracker")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		return NoconfigFile
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}

func CreateConfig(app_config *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Can't find home directory: %w", err)
	}
	configDir := filepath.Join(home, ".expense-tracker")
	configFile := filepath.Join(configDir, "config.yaml")

	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("Error creating config dir: %w", err)
	}
	data, err := yaml.Marshal(app_config)
	if err != nil {
		return fmt.Errorf("Error marshaling config into yaml: %w", err)
	}
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		return fmt.Errorf("Error creating config file: %w", err)
	}
	return nil
}
