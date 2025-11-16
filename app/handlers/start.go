package handlers

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v4"
)

// StartHandler handles the /start command
func StartHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Msg("Start command")

	storage := getStorage()
	sender, err := storage.GetSenderByID(context.Sender().ID)
	if err != nil {
		RegisterNewUser(context)
		return nil
	}

	var text string

	if sender.Email == "" {
		text = "Welcome to the Kindle Bot 🥳\nPlease send me your email 🎉"
	} else {
		text = "You are already registered 🥳\nSend me a book, please 🎉"
	}

	return context.Reply(text)
}
