package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudoguy/kindle-bot/app/settings"
)

func TestNewMailer(t *testing.T) {
	mockSettings := settings.Settings{
		FromEmail:     "test@example.com",
		EmailUsername: "username",
		EmailPassword: "password",
		SMTPHost:      "smtp.example.com",
		SMTPPort:      587,
	}

	toEmail := "recipient@example.com"

	m := NewMailer(mockSettings, toEmail)

	assert.Equal(t, mockSettings.FromEmail, m.From, "Expected From to be %s, but got %s", mockSettings.FromEmail, m.From)
	assert.Equal(t, toEmail, m.To, "Expected To to be %s, but got %s", toEmail, m.To)
	assert.Equal(t, mockSettings.EmailUsername, m.EmailUsername, "Expected EmailUsername to be %s, but got %s", mockSettings.EmailUsername, m.EmailUsername)
	assert.Equal(t, mockSettings.EmailPassword, m.EmailPassword, "Expected EmailPassword to be %s, but got %s", mockSettings.EmailPassword, m.EmailPassword)
	assert.Equal(t, mockSettings.SMTPHost, m.SMTPHost, "Expected SmtpHost to be %s, but got %s", mockSettings.SMTPHost, m.SMTPHost)
	assert.Equal(t, mockSettings.SMTPPort, m.SMTPPort, "Expected SmtpPort to be %d, but got %d", mockSettings.SMTPPort, m.SMTPPort)
}
