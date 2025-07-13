package models

type Expenses struct {
	ID          int
	Date        string
	Amount      float64
	Category    string
	Description string
}

type Budget struct {
	Month  int
	Amount float64
}
