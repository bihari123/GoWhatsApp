package main

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+918750906449")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}
