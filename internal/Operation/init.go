package operation

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

var currency string

func InitCommand() {
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().Title("What currency do you wanna use").Value(&currency),
	))
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currency)
}
