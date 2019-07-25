package utils

import (
	"fmt"
	"strconv"

	"github.com/Obsinqsob01/reto-ed/actions"
)

var option = make(chan string, 1)

// HandleOption maneja la opción mandada por el usuario
func HandleOption() {
	var o string

	o = <-option

	oint, err := strconv.Atoi(o)
	CheckErr(err, "Ocurrió un error al transformar el numero")

	ClearScreen()
	switch oint {
	case 1:
		actions.List()
	case 2:
		actions.Consult()
	case 3:
		actions.Create()
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	}
}
