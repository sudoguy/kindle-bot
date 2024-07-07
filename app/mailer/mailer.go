package mailer

import (
	"github.com/sudoguy/kindle-bot/app/settings"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	From string
	To   string

	EmailUsername string
	EmailPassword string

	SmtpHost string
	SmtpPort int
}

func NewMailer(settings settings.Settings, toEmail string) *Mailer {
	return &Mailer{
		From:          settings.FromEmail,
		To:            toEmail,
		EmailUsername: settings.EmailUsername,
		EmailPassword: settings.EmailPassword,
		SmtpHost:      settings.SmtpHost,
		SmtpPort:      settings.SmtpPort,
	}
}

func (m *Mailer) SendBook(bookPath string) error {
	message := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Subject", "New book from Awesome Kindle Bot!")
	message.SetBody("text/html", "Get your book!")
	message.Attach(bookPath)

	dialer := gomail.NewDialer(m.SmtpHost, m.SmtpPort, m.EmailUsername, m.EmailPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
