package mail

import (
	"encoding/base64"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

func (s *SendGridSender) SendMail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	from := mail.NewEmail(s.AppName, s.AppEmail)

	// Build recipients
	toRecipients := []*mail.Email{}
	for _, recipient := range to {
		toRecipients = append(toRecipients, mail.NewEmail("", recipient))
	}

	// Create email message
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.Subject = subject

	// Add content
	message.AddContent(mail.NewContent("text/html", content))

	// Add recipients
	personalization := mail.NewPersonalization()
	for _, recipient := range toRecipients {
		personalization.AddTos(recipient)
	}
	for _, ccRecipient := range cc {
		personalization.AddCCs(mail.NewEmail("", ccRecipient))
	}
	for _, bccRecipient := range bcc {
		personalization.AddBCCs(mail.NewEmail("", bccRecipient))
	}
	message.AddPersonalizations(personalization)

	for _, filePath := range attachFiles {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read attachment: %w", err)
		}
		encodedContent := base64.StdEncoding.EncodeToString(fileContent)

		attachment := mail.NewAttachment()
		attachment.SetContent(encodedContent)
		attachment.SetFilename(filePath)
		attachment.SetType("application/octet-stream")
		attachment.SetDisposition("attachment")
		message.AddAttachment(attachment)
	}

	client := sendgrid.NewSendClient(s.APIKey)
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("email sending failed with status: %d, body: %s", response.StatusCode, response.Body)
	}

	return nil
}
