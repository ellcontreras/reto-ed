package main

import "github.com/Obsinqsob01/reto-ed/utils"

func main() {
	utils.ClearScreen()
	go utils.ShowMenu()
	utils.HandleOption()
}
