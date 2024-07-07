package handlers

import (
	"os"
	"strconv"
	"time"

	"github.com/sudoguy/kindle-bot/app/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

// StatusHandler handles the /status command
func StatusHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Msg("Status command")

	storage := utils.NewStorage()
	sender, err := storage.GetSenderByID(context.Sender().ID)

	if err != nil {
		RegisterNewUser(context, storage)
		return nil
	}

	text := "Attached email: " + sender.Email + "\n"
	text += "Books sent: " + strconv.Itoa(sender.BooksSent)

	return context.Reply(text)
}
