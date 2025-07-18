package operation

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
)

func checkbudgetExists(db *sql.DB, month int) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM budget WHERE month = ?)`, month).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("Error checking if budget exists: %w", err)
	}
	return exists, nil
}

func SetBudget(db *sql.DB, month int, amount float64) error {
	var overwrite bool
	exists, err := checkbudgetExists(db, month)
	if err != nil {
		return err
	} else if !exists {

		stmt := `INSERT INTO budget (month,amount) VALUES (?,?)`
		_, err := db.Exec(stmt, month, amount)
		if err != nil {
			return fmt.Errorf("Error adding budget: %w", err)
		}
		return nil
	} else {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().Title("There is a budget already set do you wanna overwrite it?").Affirmative("yes").Negative("no").Value(&overwrite),
			),
		)
		form.Run()
		if overwrite {
			stmt := `UPDATE budget SET amount = ? where month = ?`
			_, err := db.Exec(stmt, amount, month)
			if err != nil {
				return fmt.Errorf("Error overwriting budget: %w", err)
			}
			return nil
		} else {
			fmt.Println("Exiting without overwriting budget")
			return nil
		}
	}
}

func CheckBudget(db *sql.DB, month int) error {
	var budget sql.NullFloat64
	query := `SELECT amount FROM budget WHERE month = ?`
	err := db.QueryRow(query, month).Scan(&budget)
	if err != nil {
		return fmt.Errorf("error getting budget amount: %w", err)
	}

	var total sql.NullFloat64
	query = `SELECT SUM(amount) FROM expenses WHERE  strftime('%m',date) = ? AND strftime('%Y',date) = ? `
	year := time.Now().Year()
	err = db.QueryRow(query, fmt.Sprintf("%02d", month), fmt.Sprintf("%d", year)).Scan(&total)
	if err != nil {
		return fmt.Errorf("error getting sum of expenses in the month:%w", err)
	}

	fmt.Printf("The budget limit is: %.2f \n This month expenses is: %.2f", budget.Float64, total.Float64)

	return nil
}
