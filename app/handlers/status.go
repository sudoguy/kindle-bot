package handlers

import (
	"kindle-bot/app/utils"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func StatusHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Msg("Status command")

	storage := utils.NewStorage()
	sender, err := storage.GetSenderById(context.Sender().ID)

	if err != nil {
		RegisterNewUser(context, storage)
		return nil
	}

	text := "Attached email: " + sender.Email + "\n"
	text += "Books sent: " + strconv.Itoa(sender.BooksSent)

	return context.Reply(text)
}
