package mailer

import (
	"fmt"

	"gopkg.in/gomail.v2"

	"github.com/sudoguy/kindle-bot/app/settings"
)

// Mailer represents an email sender.
type Mailer struct {
	From string
	To   string

	EmailUsername string
	EmailPassword string

	SMTPHost string
	SMTPPort int
}

// NewMailer creates a new Mailer instance.
func NewMailer(appSettings settings.Settings, toEmail string) *Mailer {
	return &Mailer{
		From:          appSettings.FromEmail,
		To:            toEmail,
		EmailUsername: appSettings.EmailUsername,
		EmailPassword: appSettings.EmailPassword,
		SMTPHost:      appSettings.SMTPHost,
		SMTPPort:      appSettings.SMTPPort,
	}
}

// SendBook sends a book to the recipient.
func (m *Mailer) SendBook(bookPath string) error {
	message := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Subject", "New book from Awesome Kindle Bot!")
	message.SetBody("text/html", "Get your book!")
	message.Attach(bookPath)

	dialer := gomail.NewDialer(m.SMTPHost, m.SMTPPort, m.EmailUsername, m.EmailPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
