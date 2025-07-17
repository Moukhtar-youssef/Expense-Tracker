/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/

package models

type Expenses struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

type Budget struct {
	Month  int
	Amount float64
}
