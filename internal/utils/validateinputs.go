package utils

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func ValidateMonth(month int) (int, error) {
	if month > 12 || month <= 0 {
		return 0, errors.New("please enter a valid month number")
	}
	return month, nil
}

func CleanStrings(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func TitleCaseStrings(input string) string {
	input = strings.TrimSpace(input)
	words := strings.Fields(input)

	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func ValidateAmount(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("please enter an amount larger than 0")
	}
	return amount, nil
}

func IspositiveInt(n int, fieldName string) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("please enter a positive number for %s", fieldName)
	}
	return n, nil
}
