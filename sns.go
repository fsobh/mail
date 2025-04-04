package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SendSMS sends an SMS message to a single recipient using AWS SNS
func (s *SNSSender) SendSMS(message string, to []string) error {
	if len(to) == 0 {
		return fmt.Errorf("no recipients provided")
	}

	// Send to first recipient only
	recipient := to[0]
	return s.sendSingleSMS(message, recipient)
}

// SendBulkSMS sends SMS messages to multiple recipients using AWS SNS
func (s *SNSSender) SendBulkSMS(message string, to []string) error {
	if len(to) == 0 {
		return fmt.Errorf("no recipients provided")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.Region),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	snsClient := sns.New(sess)
	
	var errs []string
	for _, recipient := range to {
		_, err := snsClient.Publish(&sns.PublishInput{
			Message:     aws.String(message),
			PhoneNumber: aws.String(recipient),
			MessageAttributes: map[string]*sns.MessageAttributeValue{
				"AWS.SNS.SMS.SenderID": {
					DataType:    aws.String("String"),
					StringValue: aws.String(s.SenderID),
				},
				"AWS.SNS.SMS.SMSType": {
					DataType:    aws.String("String"),
					StringValue: aws.String("Transactional"),
				},
			},
		})
		
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to send to %s: %v", recipient, err))
		}
	}
	
	if len(errs) > 0 {
		return fmt.Errorf("bulk SMS send had errors: %v", errs)
	}
	
	return nil
}

// sendSingleSMS is a helper function to send an SMS to a single recipient
func (s *SNSSender) sendSingleSMS(message string, to string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.Region),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	snsClient := sns.New(sess)
	
	_, err = snsClient.Publish(&sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(to),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(s.SenderID),
			},
			"AWS.SNS.SMS.SMSType": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Transactional"),
			},
		},
	})
	
	if err != nil {
		return fmt.Errorf("failed to send SMS: %w", err)
	}
	
	return nil
} 