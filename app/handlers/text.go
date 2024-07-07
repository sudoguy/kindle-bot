package handlers

import (
	"os"
	"time"

	"github.com/sudoguy/kindle-bot/app/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

// TextHandler handles the text message
func TextHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Str("text", context.Text()).Msg("Received message")

	storage := utils.NewStorage()
	sender, err := storage.GetSenderByID(context.Sender().ID)

	if err != nil {
		RegisterNewUser(context, storage)
		return nil
	}

	if utils.IsValidMail(context.Text()) {
		saveEmailAndUpdateSender(sender, context, storage)
	} else {
		err := context.Reply("Invalid email, please try again 🥳")
		if err != nil {
			log.Error().Err(err)
		}
	}

	return nil
}

func saveEmailAndUpdateSender(sender *utils.SenderInfo, context tele.Context, storage *utils.Storage) {
	sender.Email = context.Text()
	sender.UserName = context.Sender().Username
	sender.TelegramID = context.Sender().ID

	storage.UpdateSender(sender)
	log.Info().Str("email", sender.Email).Int64("telegram_id", sender.TelegramID).Msg("Email saved")

	msg := "Email saved 🥳\nNow Send me book, please 🎉"

	err := context.Reply(msg)

	if err != nil {
		log.Error().Err(err)
	}
}

// RegisterNewUser registers a new user
func RegisterNewUser(context tele.Context, storage *utils.Storage) {
	tgSender := context.Sender()
	log.Info().Int64("telegram_id", tgSender.ID).Str("username", tgSender.Username).Msg("New user")
	msg := "Send me your email, please 📧"

	sender := &utils.SenderInfo{
		TelegramID: tgSender.ID,
		UserName:   tgSender.Username,
	}
	storage.Senders = append(storage.Senders, *sender)
	storage.WriteSendersToFile()

	err := context.Reply(msg)
	if err != nil {
		log.Error().Err(err)
	}
}
