package operation

import (
	"Expense_tracker/internal/storage"
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func DoctorChecking() error {
	ExecPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("Error getting Executable path: %w", err)
	}
	mainDir := path.Dir(ExecPath)
	DBpath := path.Join(mainDir, "Expenses.db")

	// check DB
	fmt.Println("Checking DB....")

	err = storage.InitDB()
	if err != nil {
		return err
	}

	fmt.Println("✅ Database connected: Yes at " + DBpath)

	// check Table

	fmt.Println("Checking Tables exists....")

	exists, err := checkTableExists(storage.DB, "expenses")
	if err != nil {
		return fmt.Errorf("❌ Error checking table: %w", err)
	} else if !exists {
		fmt.Println("❌ Table 'expenses' exists: No")
		fmt.Println("Creating 'expenses' table")
		err := createTable(storage.DB, "expenses",
			"id INTEGER PRIMARY KEY AUTOINCREMENT",
			"category TEXT NOT NULL",
			"amount REAL NOT NULL",
			"date TEXT NOT NULL",
			"description TEXT NOT NULL")
		if err != nil {
			fmt.Println("❌ Error creating 'expenses' table")
			return fmt.Errorf("Error creating 'expenses' table: %w", err)
		}
		fmt.Println("✅ Table 'expenses' created: Yes")
	} else {
		fmt.Println("✅ Table 'expenses' exists: Yes")
	}

	exists, err = checkTableExists(storage.DB, "budget")
	if err != nil {
		return fmt.Errorf("❌ Error checking budget table: %w", err)
	} else if !exists {
		fmt.Println("❌ Table 'budget' exists: NO")
		fmt.Println("Creating 'budget' table")
		err := createTable(storage.DB, "budget",
			"id INTEGER PRIMARY KEY AUTOINCREMENT",
			"month TEXT NOT NULL",
			"budget REAL NOT NULL")
		if err != nil {
			fmt.Println("❌ Error creating 'budget table'")
			return fmt.Errorf("Error creating 'budget' table: %w", err)
		}
		fmt.Println("✅ Table 'budget' created: YES")
	} else {
		fmt.Println("✅ Table 'budget', exists: YES")
	}

	// check file system

	fmt.Println("Checking file system writability....")

	err = canWriteToPath(mainDir)
	if err != nil {
		return fmt.Errorf("🚫 Cannot write to export directory: %w", err)
	}

	fmt.Println("✅ File system writable: Yes")

	return nil
}

func checkTableExists(db *sql.DB, tableName string) (bool, error) {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?`
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error checking table: %w", err)
	}

	return true, nil
}

func createTable(db *sql.DB, tableName string, columns ...string) error {
	if len(columns) == 0 {
		return fmt.Errorf("no columns specified")
	}

	columnsdef := strings.Join(columns, ",\n ")
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		%s
		);`, tableName, columnsdef)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func canWriteToPath(dir string) error {
	testFile := filepath.Join(dir, ".perm_test")

	f, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("write permission denied for %s: %w", dir, err)
	}
	defer os.Remove(testFile)
	defer f.Close()

	_, err = f.WriteString("test")
	if err != nil {
		return fmt.Errorf("failed to write to test file: %w", err)
	}
	return nil
}
