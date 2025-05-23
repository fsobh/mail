
# Mailing Module for Golang

A Golang module for sending emails seamlessly using [AWS](https://aws.amazon.com/ses/) or [Send Grid](https://sendgrid.com/).

## Features
- **AWS SES Integration**: Send transactional emails via AWS SES.
- **Send Grid Support**: Integrate with SendGrid for email delivery.
- **Simple API**: Easy-to-use interface for developers.

---

## Installation
Get started by installing the package:

```sh
go get github.com/fsobh/mail
```

---

## Usage

### Initialize the Mailer
Before sending emails, configure the mailer with your provider's credentials.

#### Ensure that the sender email has been verified with its respective provider :
1. AWS SES guide [here](https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html)
2. Send Grid guide [here](https://www.twilio.com/docs/sendgrid/ui/sending-email/sender-verification)
```go
package main

import (
	"fmt"
	"github.com/fsobh/mail"
)

func main() {

	/*
	 Never paste your credentials directly in the code for safety reasons
	 Recommended : Store these in environment variables using [viper](https://github.com/spf13/viper)
	*/

	senderEmail := "example@sender.com"
	sesRegion :=  "us-east-1"
	sgAPIKey := "<your-api-key>"
	sgAppName := "<your-app-name>"

	//Initialize a new AWS SES mailer
	mailerSES := mail.NewSESSender(sesRegion, senderEmail)

	//Initialize a new Send Grid mailer
	mailerSendGrid := mail.NewSendGridSender(sgAPIKey, sgAppName, senderEmail)

	subject := "Email Subject"
	contentHtml := fmt.Sprintf("<h1>Email HTML content</h1>")
	to := []string{"reciever@email.com"}
	cc := []string{"cc@email.com"}
	bcc := []string{"bcc@email.com>"}
	attachments := []string{"../files/example.txt"}

	// SES example
	err := mailerSES.SendMail(subject, contentHtml, to, cc, bcc, nil) //can't do attachments on SES

	if err != nil {
		_ = fmt.Errorf("error sending email via SES: %s", err)
	}

	// Send Grid example
	err = mailerSendGrid.SendMail(subject, contentHtml, to, cc, bcc, attachments)

	if err != nil {
		_ = fmt.Errorf("error sending email via Send Grid: %s", err)
	}

	fmt.Println("Email sent successfully")

}
```
## Notes
#### To test AWS SES sends locally, you must export your IAM roles Access ID & Secret :
1. Follow [this](https://aws.amazon.com/blogs/security/wheres-my-secret-access-key/) guide to generate the credentials for your IAM account **(make sure you give it the right access policies to send Emails via SES)**.
2. Export your credentials locally in your terminal
```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_DEFAULT_REGION="us-west-2" # Change to your region
```

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
