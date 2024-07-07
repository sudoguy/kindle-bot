package settings

import (
	"errors"
	"os"
	"strconv"
)

// Settings represents the application settings.
type Settings struct {
	BotToken string `json:"bot_token"`

	FromEmail     string `json:"from_email"`
	EmailUsername string `json:"email_username"`
	EmailPassword string `json:"email_password"`
	SMTPHost      string `json:"smtp_host"`
	SMTPPort      int    `json:"smtp_port"`
}

// NewSettings creates a new Settings instance.
func NewSettings() *Settings {
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	return &Settings{
		BotToken: os.Getenv("BOT_TOKEN"),

		FromEmail:     os.Getenv("FROM_EMAIL"),
		EmailUsername: os.Getenv("EMAIL_USERNAME"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      smtpPort,
	}
}

// Validate validates the settings.
func (settings *Settings) Validate() error {
	if settings.BotToken == "" {
		return errors.New("BOT_TOKEN is required")
	}
	if settings.FromEmail == "" {
		return errors.New("FROM_EMAIL is required")
	}
	if settings.EmailUsername == "" {
		return errors.New("EMAIL_USERNAME is required")
	}
	if settings.EmailPassword == "" {
		return errors.New("EMAIL_PASSWORD is required")
	}
	if settings.SMTPHost == "" {
		return errors.New("SMTP_HOST is required")
	}
	if settings.SMTPPort == 0 {
		return errors.New("SMTP_PORT is required")
	}

	return nil
}
