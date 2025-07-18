package utils

import (
	"log"
)

func Info(msg string) {
	log.Println(msg)
}

func Warn(msg string) {
	log.Println(msg)
}

func Error(err error) {
	log.Fatal(err)
}
