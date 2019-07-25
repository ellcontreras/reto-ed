package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

// Create crea una foto
func Create() {

	color.Green("Ingresa el titulo:")
	t, _, _ := reader.ReadLine()
	title := string(t)

	color.Green("Ingresa el url de la imagen:")
	u, _, _ := reader.ReadLine()
	url := string(u)

	color.Green("Ingresa el url del thumbnail:")
	tu, _, _ := reader.ReadLine()
	thumbnailURL := string(tu)

	jsonData := map[string]string{
		"albumId":      "1",
		"title":        title,
		"url":          url,
		"thumbnailUrl": thumbnailURL,
	}

	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(APIBase, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
