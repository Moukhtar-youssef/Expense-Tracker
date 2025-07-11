package internal

import (
	"fmt"
	"os"
	"path"
)

func FindExeDir() (string, error) {
	Exec, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("Error finding where the Executable dir is : %w", err)
	}

	ExecDir := path.Dir(Exec)

	return path.Join(ExecDir, "Expenses.db"), nil
}
