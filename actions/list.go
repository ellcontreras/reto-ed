package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIBase es el endpoint del api
const APIBase = "https://jsonplaceholder.typicode.com/photos"

// List obtiene todas las fotos del api y las muestra
func List() {
	response, err := http.Get(APIBase)
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}
