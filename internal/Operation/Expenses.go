/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/

package operation

import (
	"Expense_tracker/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

func checkExists(db *sql.DB, id int) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM expenses WHERE id = ?)`, id).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("Error checking if expense exists: %w", err)
	}
	return exists, nil
}

func AddExepnse(db *sql.DB, date string, amount float64, category string, description string) error {
	stmt := `INSERT INTO expenses (category , amount, date, description) VALUES (?,?,?,?)`
	_, err := db.Exec(stmt, category, amount, date, description)
	if err != nil {
		return fmt.Errorf("Error adding expense: %w", err)
	}
	return nil
}

func ListExpenses(db *sql.DB, category, from, to string, month int) error {
	query := `SELECT id , category , amount , date , description FROM expenses`
	var conditions []string
	var args []any

	if category != "" {
		conditions = append(conditions, "category = ?")
		args = append(args, category)
	}
	if from != "" && to != "" {
		conditions = append(conditions, "date BETWEEN ? AND ?")
		args = append(args, from, to)
	}
	if month != -1 {
		conditions = append(conditions, "strftime('%m',date) = ?")
		args = append(args, fmt.Sprintf("%02d", month))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY id"

	rows, err := db.Query(query, args...)
	if err != nil {
		return fmt.Errorf("Error querying expenses: %w", err)
	}
	defer rows.Close()

	var expenses []models.Expenses
	for rows.Next() {
		var e models.Expenses
		err = rows.Scan(&e.ID, &e.Category, &e.Amount, &e.Date, &e.Description)
		if err != nil {
			return fmt.Errorf("Error scanning expenses rows: %w", err)
		}
		expenses = append(expenses, e)
	}
	for _, e := range expenses {
		fmt.Printf("ID: %d | %s | %.2f | %s | %s\n", e.ID, e.Category, e.Amount, e.Date, e.Description)
	}

	return nil
}

func DeleteExpense(db *sql.DB, id int) error {
	exists, err := checkExists(db, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Expense with ID %d not found", id)
	}
	stmt := `DELETE FROM expenses WHERE id = ?`
	_, err = db.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("Error deleting epxpense from table: %w", err)
	}
	return nil
}

func DeleteAllExpenses(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM expenses;`)
	if err != nil {
		return fmt.Errorf("Error deleting all rows from table: %w", err)
	}
	_, err = db.Exec(`DELETE FROM sqlite_sequence WHERE name='expenses';`)
	if err != nil {
		return fmt.Errorf("Error deleting AUTOINCREMENT id from table: %w", err)
	}
	return nil
}
