package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var token = os.Getenv("WHATSAPP_TOKEN")

func SendMessage(ad IkmanAd) {
	body := []byte(`{
        "messaging_product": "whatsapp",
        "recipient_type": "individual",
        "to": "",
        "type": "template",
        "template": {
            "name": "ikman_template",
            "language": {
                "code": "en"
            },
            "components": [
				{
					"type": "header",
                    "parameters": []
				},
                {
                    "type": "body",
                    "parameters": []
                }
            ]
        }
    }`)

	var message TemplateMessage
	if err := json.Unmarshal(body, &message); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	var title = Parameter{Type: "text", Text: ad.Title, Image: nil}
	var price = Parameter{Type: "text", Text: ad.Price, Image: nil}
	var link = Parameter{Type: "text", Text: ad.Link, Image: nil}
	var image = Parameter{Type: "image", Image: ImageParameterType{Link: ad.Image}, Text: nil}
	message.Template.Components[1].Parameters = append(message.Template.Components[1].Parameters, title, price, link)
	message.Template.Components[0].Parameters = append(message.Template.Components[0].Parameters, image)

	phoneNumbers := [2]string{"+94742400690", "+94768038766"}
	for _, phoneNumber := range phoneNumbers {
		message.To = phoneNumber

		// Marshal the map into JSON format
		requestBodyBytes, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshaling request body:", err)
			return
		}

		req, err := http.NewRequest("POST", "https://graph.facebook.com/v17.0/114095055017089/messages",
			bytes.NewBuffer(requestBodyBytes))

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		// Create an HTTP client
		client := &http.Client{}

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		resp.Body.Close()
		fmt.Println("Response Status:", resp.Status)
	}
}
