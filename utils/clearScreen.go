package utils

import (
	"os"
	"os/exec"
	"runtime"
)

var (
	clear = make(map[string]func())
)

func init() {
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

// ClearScreen borra la terminal dejandola en blanco para escribir nueva información
func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Lo siento, tu plataforma no esta soportada aún para limpiar la pantalla :(")
	}
}
