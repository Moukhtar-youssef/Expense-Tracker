/*
Copyright © 2025 Moukhtar youssef moukhtar.youssef06@gmail.com
*/

package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func getDBfilepath() (string, error) {
	ExecPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("Error getting Executable path: %w", err)
	}
	mainDir := path.Dir(ExecPath)
	return path.Join(mainDir, "Expenses.db"), nil
}

func InitDB() error {
	DBfilepath, err := getDBfilepath()
	if err != nil {
		return err
	}
	DB, err = sql.Open("sqlite3", DBfilepath)
	if err != nil {
		return fmt.Errorf("Error opening sql db: %w", err)
	}
	createTable := `
    CREATE TABLE IF NOT EXISTS expenses (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        category TEXT NOT NULL,
        amount REAL NOT NULL,
        date TEXT NOT NULL,
    	description TEXT NOT NULL
    );`
	_, err = DB.Exec(createTable)
	if err != nil {
		return fmt.Errorf("Error creating table in sqlite3: %w", err)
	}
	return DB.Ping()
}
