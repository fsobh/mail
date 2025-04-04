package mail

// SMSSender interface defines methods for sending SMS messages
type SMSSender interface {
	SendSMS(
		message string,
		to []string,
	) error

	SendBulkSMS(
		message string,
		to []string,
	) error
}

// TwilioSender implements the SMSSender interface for Twilio
type TwilioSender struct {
	AccountSID       string
	AuthToken        string
	SenderPhoneNumber string
}

// SNSSender implements the SMSSender interface for AWS SNS
type SNSSender struct {
	Region           string
	SenderID         string
}

// NewTwilioSender initializes a new TwilioSender
func NewTwilioSender(accountSID string, authToken string, senderPhoneNumber string) SMSSender {
	return &TwilioSender{
		AccountSID:       accountSID,
		AuthToken:        authToken,
		SenderPhoneNumber: senderPhoneNumber,
	}
}

// NewSNSSender initializes a new SNSSender
func NewSNSSender(region string, senderID string) SMSSender {
	return &SNSSender{
		Region:   region,
		SenderID: senderID,
	}
} 