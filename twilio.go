package mail

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SendSMS sends an SMS message to a single recipient using Twilio
func (t *TwilioSender) SendSMS(message string, to []string) error {
	if len(to) == 0 {
		return fmt.Errorf("no recipients provided")
	}

	// Send to first recipient only
	recipient := to[0]
	return t.sendSingleSMS(message, recipient)
}

// SendBulkSMS sends SMS messages to multiple recipients using Twilio
func (t *TwilioSender) SendBulkSMS(message string, to []string) error {
	if len(to) == 0 {
		return fmt.Errorf("no recipients provided")
	}

	var errs []string
	
	for _, recipient := range to {
		err := t.sendSingleSMS(message, recipient)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to send to %s: %v", recipient, err))
		}
	}
	
	if len(errs) > 0 {
		return fmt.Errorf("bulk SMS send had errors: %s", strings.Join(errs, "; "))
	}
	
	return nil
}

// sendSingleSMS is a helper function to send an SMS to a single recipient
func (t *TwilioSender) sendSingleSMS(message string, to string) error {
	endpoint := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.AccountSID)
	
	data := url.Values{}
	data.Set("To", to)
	data.Set("From", t.SenderPhoneNumber)
	data.Set("Body", message)
	
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	req.SetBasicAuth(t.AccountSID, t.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("SMS sending failed with status: %d", resp.StatusCode)
	}
	
	return nil
} 