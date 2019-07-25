package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

// Update actualiza una foto
func Update() {
	color.Green("Ingresa el id de la foto:")
	i, _, _ := reader.ReadLine()
	id := string(i)

	color.Green("Ingresa el nuevo titulo:")
	t, _, _ := reader.ReadLine()
	title := string(t)

	color.Green("Ingresa el nuevo url de la imagen:")
	u, _, _ := reader.ReadLine()
	url := string(u)

	color.Green("Ingresa el nuevo url del thumbnail:")
	tu, _, _ := reader.ReadLine()
	thumbnailURL := string(tu)

	jsonData := map[string]string{
		"id":           id,
		"albumId":      "1",
		"title":        title,
		"url":          url,
		"thumbnailUrl": thumbnailURL,
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest("PUT", APIBase+"/"+id, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
