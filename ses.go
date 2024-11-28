package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func (s *SESSender) SendMail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.Region),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := ses.New(sess)

	// Construct the message
	dest := &ses.Destination{
		ToAddresses:  aws.StringSlice(to),
		CcAddresses:  aws.StringSlice(cc),
		BccAddresses: aws.StringSlice(bcc),
	}

	// Prepare body
	body := &ses.Content{
		Charset: aws.String("UTF-8"),
		Data:    aws.String(content),
	}

	// Prepare subject
	subj := &ses.Content{
		Charset: aws.String("UTF-8"),
		Data:    aws.String(subject),
	}

	msg := &ses.Message{
		Body: &ses.Body{
			Html: body,
		},
		Subject: subj,
	}

	// Add attachments
	if len(attachFiles) > 0 {
		return fmt.Errorf("attachments are not supported in the SES API via this method")
	}

	input := &ses.SendEmailInput{
		Destination: dest,
		Message:     msg,
		Source:      aws.String(s.SenderEmail),
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
