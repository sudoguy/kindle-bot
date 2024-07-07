package mailer

import (
	"testing"

	"github.com/sudoguy/kindle-bot/app/settings"

	"github.com/stretchr/testify/assert"
)

func TestNewMailer(t *testing.T) {
	mockSettings := settings.Settings{
		FromEmail:     "test@example.com",
		EmailUsername: "username",
		EmailPassword: "password",
		SmtpHost:      "smtp.example.com",
		SmtpPort:      587,
	}

	toEmail := "recipient@example.com"

	m := NewMailer(mockSettings, toEmail)

	assert.Equal(t, mockSettings.FromEmail, m.From, "Expected From to be %s, but got %s", mockSettings.FromEmail, m.From)
	assert.Equal(t, toEmail, m.To, "Expected To to be %s, but got %s", toEmail, m.To)
	assert.Equal(t, mockSettings.EmailUsername, m.EmailUsername, "Expected EmailUsername to be %s, but got %s", mockSettings.EmailUsername, m.EmailUsername)
	assert.Equal(t, mockSettings.EmailPassword, m.EmailPassword, "Expected EmailPassword to be %s, but got %s", mockSettings.EmailPassword, m.EmailPassword)
	assert.Equal(t, mockSettings.SmtpHost, m.SmtpHost, "Expected SmtpHost to be %s, but got %s", mockSettings.SmtpHost, m.SmtpHost)
	assert.Equal(t, mockSettings.SmtpPort, m.SmtpPort, "Expected SmtpPort to be %d, but got %d", mockSettings.SmtpPort, m.SmtpPort)
}
