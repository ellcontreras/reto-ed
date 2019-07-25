package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Obsinqsob01/reto-ed/utils"
)

// APIBase es el endpoint del api
const APIBase = "https://jsonplaceholder.typicode.com/photos"

// List obtiene todas las fotos del api y las muestra
func List() {
	response, err := http.Get(APIBase)
	utils.CheckErr(err, "Ocurrió un error al tratar de llamar al api")

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}
