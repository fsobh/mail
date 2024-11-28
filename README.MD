
# Mailing Module for Golang (AWS, SendGrid)

A Golang module for sending emails seamlessly using AWS or SendGrid.

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

```go
package main

import (
	"fmt"
	"github.com/fsobh/mail"
)

func main() {

	mailerSES := mail.NewSESSender("us-east-1", "AWSverifiedemail@example.com")
	mailerSendGrid := mail.NewSendGridSender("<your-api-key>", "<your-app-name>", "<SENDGRID_EMAIL_FROM@email.com>")

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


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
