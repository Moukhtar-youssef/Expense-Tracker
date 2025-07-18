/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/

package operation

import (
	"Expense_tracker/internal/models"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
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

func UpdateExpense(db *sql.DB, id int, date string, amount float64, category string, description string) error {
	stmt := `UPDATE expenses`
	var updates []string
	var args []any

	if date != "" {
		updates = append(updates, "date = ?")
		args = append(args, date)
	}
	if amount != -1 {
		updates = append(updates, "amount = ?")
		args = append(args, amount)
	}
	if category != "" {
		updates = append(updates, "category = ?")
		args = append(args, category)
	}
	if description != "" {
		updates = append(updates, "description = ?")
		args = append(args, description)
	}

	if len(updates) > 0 {
		stmt += " SET " + strings.Join(updates, ",")
	}
	stmt += " WHERE id = ?"
	args = append(args, id)

	_, err := db.Exec(stmt, args...)
	if err != nil {
		return fmt.Errorf("Error updating expemse: %w", err)
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
		year := time.Now().Year()
		conditions = append(conditions, "strftime('%m',date) = ?", "strftime('%Y',date) = ?")
		args = append(args, fmt.Sprintf("%02d", month), fmt.Sprintf("%d", year))
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
		fmt.Printf("ID: %d | %s | %.2f %s | %s | %s\n", e.ID, e.Category, e.Amount, viper.GetString("currency_symbol"), e.Date, e.Description)
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

func ExportExpenses(db *sql.DB, exportType string, exportFilename string, category string, month int) error {
	query := `SELECT id, date, amount, category, description FROM expenses`
	var conditions []string
	var args []any
	if category != "" {
		conditions = append(conditions, " category = ? ")
		args = append(args, category)
	}
	if month != -1 {
		conditions = append(conditions, "strftime('%m',date) = ?")
		args = append(args, fmt.Sprintf("%02d", month))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += "ORDER BY id"

	rows, err := db.Query(query, args...)
	if err != nil {
		return fmt.Errorf("Error querying expenses for export: %w", err)
	}
	defer rows.Close()

	var expenses []models.Expenses

	for rows.Next() {
		var e models.Expenses
		err = rows.Scan(&e.ID, &e.Date, &e.Amount, &e.Category, &e.Description)
		if err != nil {
			return fmt.Errorf("Error scanning row for export: %w", err)
		}
		expenses = append(expenses, e)
	}
	switch exportType {
	case "json":
		data, err := json.MarshalIndent(expenses, "", "  ")
		if err != nil {
			return fmt.Errorf("Error marshaling expenses to JSON: %w", err)
		}
		err = os.WriteFile(fmt.Sprintf("%s.json", exportFilename), data, 0644)
		if err != nil {
			return fmt.Errorf("Error writing JSON to file: %w", err)
		}
	case "csv":
		file, err := os.Create(fmt.Sprintf("%s.csv", exportFilename))
		if err != nil {
			return fmt.Errorf("Error creating csv file: %w", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		header := []string{"id", "date", "amount", "category", "description"}
		err = writer.Write(header)
		if err != nil {
			return fmt.Errorf("Error writing csv header: %w", err)
		}

		for _, expense := range expenses {
			record := []string{
				fmt.Sprintf("%d", expense.ID),
				expense.Date,
				fmt.Sprintf("%.2f %s", expense.Amount, viper.GetString("currency_symbol")),
				expense.Category,
				expense.Description,
			}
			err := writer.Write(record)
			if err != nil {
				return fmt.Errorf("Error writing csv records: %w", err)
			}
		}
	}

	return nil
}

func SummarizeExpenses(db *sql.DB, category string, month int) error {
	query := `SELECT SUM(amount) FROM expenses`
	var conditions []string
	var args []any

	if category != "" {
		conditions = append(conditions, "category = ?")
		args = append(args, category)
	}

	if month != -1 {
		conditions = append(conditions, "strftime('%m', date) = ?")
		args = append(args, fmt.Sprintf("%02d", month))
	}

	year := time.Now().Year()
	conditions = append(conditions, "strftime('%Y', date) = ?")
	args = append(args, fmt.Sprintf("%d", year))

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var total sql.NullFloat64
	err := db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return fmt.Errorf("error getting summary: %w", err)
	}
	fmt.Printf("Total expenses")
	if category != "" {
		fmt.Printf(" in category '%s'", category)
	}
	if month != -1 {
		fmt.Printf(" for month %d", month)
	}
	fmt.Printf(" in %d: %.2f\n", year, total.Float64)
	return nil
}
