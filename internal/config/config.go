package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

func InitConfig() {
	ExecPath, err := os.Executable()
	mainDir := path.Dir(ExecPath)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Can't find home directory: %v", err)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(mainDir)
	viper.AddConfigPath(home + "/.expense-tracker")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found. Using defaults")
	}
}

func CreateConfig() error {
	return nil
}
