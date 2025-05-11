package main

import (
	"fmt"

	"github.com/fsobh/mail"
)

func main() {
	// Initialize SNS sender with a proper alphanumeric sender ID
	sns := mail.NewSNSSender("us-east-2", "+1-us-number")

	// Test message
	message := "Welcome to RAA! Thank you for subscribing to service updates."
	recipients := []string{"+1-us-number"}

	// Test SNS
	fmt.Println("Testing SNS SMS...")
	err := sns.SendSMS(message, recipients)
	if err != nil {
		fmt.Printf("SNS SMS Error: %v\n", err)
	} else {
		fmt.Println("SNS SMS sent successfully!")
	}
}
