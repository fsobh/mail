package mail

type EmailSender interface {
	SendMail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

// SESSender implements the EmailSender interface for AWS SES.
type SESSender struct {
	Region      string
	SenderEmail string
}

// SendGridSender implements the EmailSender interface for SendGrid.
type SendGridSender struct {
	APIKey   string
	AppName  string
	AppEmail string
}

// NewSESSender initializes a new SESSender with optional credentials.
// If AccessKeyID and SecretKey are empty, the AWS SDK will use the default credentials chain.
func NewSESSender(region string, senderEmail string) EmailSender {
	return &SESSender{Region: region, SenderEmail: senderEmail}
}

// NewSendGridSender initializes a new SendGridSender.
func NewSendGridSender(apiKey string, appName string, appEmail string) EmailSender {
	return &SendGridSender{APIKey: apiKey, AppEmail: appEmail, AppName: appName}
}
