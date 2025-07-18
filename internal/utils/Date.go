/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/

package utils

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func ParseDate(input string) (string, error) {
	whenparser := when.New(nil)
	whenparser.Add(en.All...)
	whenparser.Add(common.All...)

	switch input {
	case "", "today":
		return time.Now().Format("2006-01-02"), nil

	case "yesterday":
		return time.Now().AddDate(0, 0, -1).Format("2006-01-02"), nil

	case "tomorrow":
		return time.Now().AddDate(0, 0, 1).Format("2006-01-02"), nil

	default:
		r, err := whenparser.Parse(input, time.Now())
		if err != nil {
			return "", fmt.Errorf("Error while parsing natrual language: %w", err)
		}

		if r == nil {

			t, err := now.Parse(input)
			if err != nil {
				return "", fmt.Errorf("Invalid date format or keyword")
			}
			return t.Format("2006-01-02"), nil
		}

		return r.Time.Format("2006-01-02"), nil

	}
}

func GetCurrentMonth() int {
	return int(time.Now().Month())
}

func GetCurrentYear() int {
	return time.Now().Year()
}
