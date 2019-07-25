package actions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

// Consult obtiene una foto basada en su id ingresado por el usuario
func Consult() {
	fmt.Println("Ingrese el id para realizar la consulta:")
	id, _, _ := reader.ReadLine()

	response, err := http.Get(APIBase + "/" + string(id))
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}
