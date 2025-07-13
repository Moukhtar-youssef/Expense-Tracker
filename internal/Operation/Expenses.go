package operation

import (
	"Expense_tracker/internal/models"
	"database/sql"
	"fmt"
	"log"
)

func AddExepnse(db *sql.DB, date string, amount float64, category string, description string) {
	stmt := `INSERT INTO expenses (category , amount, date, description) VALUES (?,?,?,?)`
	_, err := db.Exec(stmt, category, amount, date, description)
	if err != nil {
		log.Fatal(fmt.Errorf("Error adding expense: %w", err))
	}
}

func ListExpenses(db *sql.DB) {
	rows, err := db.Query(`SELECT id , category , amount , date , description from expenses ORDER BY ID`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var expenses []models.Expenses
	for rows.Next() {
		var e models.Expenses

		err = rows.Scan(&e.ID, &e.Category, &e.Amount, &e.Date, &e.Description)
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, e)
	}
	for _, expense := range expenses {
		fmt.Printf("ID: %d | %s | %.2f | %s | %s\n", expense.ID, expense.Category, expense.Amount, expense.Date, expense.Description)
	}
}

func DeleteExpense(db *sql.DB, id int) {
	stmt := `DELETE FROM expenses WHERE id = ?`
	_, err := db.Exec(stmt, id)
	if err != nil {
		log.Fatal(err)
	}
}
