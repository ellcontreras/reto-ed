package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

// ShowMenu muestra el menu en la pantalla
func ShowMenu() {
	b := color.New(color.FgBlue).Add(color.Bold)
	b.Println("EDteam API Information")

	fmt.Println("\nSelecciona una opción a realizar:")
	fmt.Println("1) Listar")
	fmt.Println("2) Consultar")
	fmt.Println("3) Crear")
	fmt.Println("4) Modificar")
	fmt.Println("5) Eliminar")

	op, _, _ := reader.ReadLine()

	option <- string(op)
}
