package main

import (
	"fmt"
	"github.com/fsobh/mail"
)

func main() {

	twilio := mail.NewTwilioSender("accountSid", "authToken", "fromNumber")
	sendGrid := mail.NewSendGridSender("apiKey", "appName", "appEmail")
	ses := mail.NewSESSender("region", "senderEmail")
	sns := mail.NewSNSSender("region", "senderID")

	fmt.Printf("Twilio Sender: %v\n", twilio)
	fmt.Printf("SendGrid Sender: %v\n", sendGrid)
	fmt.Printf("SES Sender: %v\n", ses)
	fmt.Printf("SNS Sender: %v\n", sns)
}
