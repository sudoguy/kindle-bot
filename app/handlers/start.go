package handlers

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func StartHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Msg("Start command")

	text := "Send me your email, please 📧"

	return context.Reply(text)
}
