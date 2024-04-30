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

	var title = Parameter{Type: "text", Text: ad.Title}
	var price = Parameter{Type: "text", Text: ad.Price}
	var link = Parameter{Type: "text", Text: ad.Link}
	message.Template.Components[0].Parameters = append(message.Template.Components[0].Parameters, title, price, link)

	message.To = "+94742400690"

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
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}
