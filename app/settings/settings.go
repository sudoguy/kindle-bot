package settings

import (
	"errors"
	"os"
	"strconv"
)

type Settings struct {
	BotToken string `json:"bot_token"`
	Debug    bool   `json:"debug"`

	FromEmail     string `json:"from_email"`
	EmailUsername string `json:"email_username"`
	EmailPassword string `json:"email_password"`
	SmtpHost      string `json:"smtp_host"`
	SmtpPort      int    `json:"smtp_port"`
}

func NewSettings() *Settings {
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	return &Settings{
		BotToken: os.Getenv("BOT_TOKEN"),
		Debug:    os.Getenv("DEBUG") == "1",

		FromEmail:     os.Getenv("FROM_EMAIL"),
		EmailUsername: os.Getenv("EMAIL_USERNAME"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		SmtpHost:      os.Getenv("SMTP_HOST"),
		SmtpPort:      smtpPort,
	}
}

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
	if settings.SmtpHost == "" {
		return errors.New("SMTP_HOST is required")
	}
	if settings.SmtpPort == 0 {
		return errors.New("SMTP_PORT is required")
	}

	return nil
}
