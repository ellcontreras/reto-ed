package actions

import (
	"fmt"
	"net/http"

	"github.com/ashwanthkumar/slack-go-webhook"
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

		attachment1 := slack.Attachment{}
		attachment1.AddField(slack.Field{Title: "Author", Value: "Luis Contreras"}).AddField(slack.Field{Title: "Hey", Value: "Se ha creado una foto!"})
		payload := slack.Payload{
			Text:        "Se ha creado una foto en el API",
			Username:    "robot",
			Channel:     "#general",
			IconEmoji:   ":monkey_face:",
			Attachments: []slack.Attachment{attachment1},
		}
		er := slack.Send(webhookUrl, "", payload)
		if len(er) > 0 {
			fmt.Printf("error: %s\n", err)
		}
	} else {
		panic("HA ocurrido un error al tratar de eliminar la foto")
	}
}
