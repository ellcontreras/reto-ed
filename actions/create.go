package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/fatih/color"
)

const webhookUrl = "https://hooks.slack.com/services/foo/bar/baz"

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

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Author", Value: "Luis Contreras"}).AddField(slack.Field{Title: "Hey", Value: "Se ha creado una foto!"})
	payload := slack.Payload{
		Text:        "Se ha creado una foto en el API",
		Username:    "robot",
		Channel:     "#general",
		IconEmoji:   ":yum:",
		Attachments: []slack.Attachment{attachment1},
	}
	er := slack.Send(webhookUrl, "", payload)
	if len(er) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
