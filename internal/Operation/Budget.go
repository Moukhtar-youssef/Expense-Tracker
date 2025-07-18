package operation

import (
	"database/sql"
	"fmt"

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
	return nil
}
