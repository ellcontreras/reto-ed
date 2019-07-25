package actions

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

// Delete elimina una foto
func Delete() {
	color.Green("Ingresa el id de la foto para borrar:")
	i, _, _ := reader.ReadLine()
	id := string(i)

	request, err := http.NewRequest("DELETE", APIBase+"/"+id, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println(response.Status)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Se ha eliminado con exito!")
	} else {
		panic("HA ocurrido un error al tratar de eliminar la foto")
	}
}
