package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func FormatAmount(amount float64) string {
	symbol := viper.GetString("currency_symbol")
	if symbol == "" {
		symbol = "$" // fallback default
	}
	return fmt.Sprintf("%10.2f %s", amount, symbol)
}
