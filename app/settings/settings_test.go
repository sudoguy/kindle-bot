package settings

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSettings(t *testing.T) {
	// Set up test environment variables
	os.Setenv("BOT_TOKEN", "test_token")
	os.Setenv("FROM_EMAIL", "test@example.com")
	os.Setenv("EMAIL_USERNAME", "test_username")
	os.Setenv("EMAIL_PASSWORD", "test_password")
	os.Setenv("SMTP_HOST", "test_host")
	os.Setenv("SMTP_PORT", "587")

	// Call the function being tested
	settings := NewSettings()

	// Verify the expected values
	assert.Equal(t, "test_token", settings.BotToken, "Expected BotToken to be 'test_token'")
	assert.Equal(t, "test@example.com", settings.FromEmail, "Expected FromEmail to be 'test@example.com'")
	assert.Equal(t, "test_username", settings.EmailUsername, "Expected EmailUsername to be 'test_username'")
	assert.Equal(t, "test_password", settings.EmailPassword, "Expected EmailPassword to be 'test_password'")
	assert.Equal(t, "test_host", settings.SMTPHost, "Expected SmtpHost to be 'test_host'")
	assert.Equal(t, 587, settings.SMTPPort, "Expected SmtpPort to be 587")

	// Clean up test environment variables
	os.Unsetenv("BOT_TOKEN")
	os.Unsetenv("FROM_EMAIL")
	os.Unsetenv("EMAIL_USERNAME")
	os.Unsetenv("EMAIL_PASSWORD")
	os.Unsetenv("SMTP_HOST")
	os.Unsetenv("SMTP_PORT")
}

func TestSettings_Validate(t *testing.T) {
	settings := &Settings{
		BotToken:      "test_token",
		FromEmail:     "test@example.com",
		EmailUsername: "test_username",
		EmailPassword: "test_password",
		SMTPHost:      "test_host",
		SMTPPort:      587,
	}

	err := settings.Validate()
	require.NoError(t, err, "Expected no error")

	// Test missing BotToken
	settings.BotToken = ""
	err = settings.Validate()
	require.EqualError(t, err, "BOT_TOKEN is required", "Expected error 'BOT_TOKEN is required'")

	// Test missing FromEmail
	settings.BotToken = "test_token"
	settings.FromEmail = ""
	err = settings.Validate()
	require.EqualError(t, err, "FROM_EMAIL is required", "Expected error 'FROM_EMAIL is required'")

	// Test missing EmailUsername
	settings.FromEmail = "test@example.com"
	settings.EmailUsername = ""
	err = settings.Validate()
	require.EqualError(t, err, "EMAIL_USERNAME is required", "Expected error 'EMAIL_USERNAME is required'")

	// Test missing EmailPassword
	settings.EmailUsername = "test_username"
	settings.EmailPassword = ""
	err = settings.Validate()
	require.EqualError(t, err, "EMAIL_PASSWORD is required", "Expected error 'EMAIL_PASSWORD is required'")

	// Test missing SmtpHost
	settings.EmailPassword = "test_password"
	settings.SMTPHost = ""
	err = settings.Validate()
	require.EqualError(t, err, "SMTP_HOST is required", "Expected error 'SMTP_HOST is required'")

	// Test missing SmtpPort
	settings.SMTPHost = "test_host"
	settings.SMTPPort = 0
	err = settings.Validate()
	require.EqualError(t, err, "SMTP_PORT is required", "Expected error 'SMTP_PORT is required'")
}
