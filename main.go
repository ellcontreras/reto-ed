package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/Obsinqsob01/reto-ed/actions"

	"github.com/Obsinqsob01/reto-ed/utils"

	"github.com/fatih/color"
)

var (
	clear  map[string]func()
	reader *bufio.Reader
	option chan string
)

func init() {
	clear = make(map[string]func())
	reader = bufio.NewReader(os.Stdin)
	option = make(chan string, 1)

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

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

// ClearScreen borra la terminal dejandola en blanco para escribir nueva información
func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Lo siento, tu plataforma no esta soportada aún para limpiar la pantalla :(")
	}
}

// HandleOption maneja la opción mandada por el usuario
func HandleOption() {
	var o string

	o = <-option

	oint, err := strconv.Atoi(o)
	utils.CheckErr(err, "Ocurrió un error al transformar el numero")

	switch oint {
	case 1:
		actions.List()
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	}
}

func main() {
	ClearScreen()
	go ShowMenu()
	HandleOption()
}
