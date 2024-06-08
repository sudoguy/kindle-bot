package handlers

import (
	"os"
	"time"

	"kindle-bot/app/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func TextHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Str("text", context.Text()).Msg("Received message")

	storage := utils.NewStorage()
	sender, err := storage.GetSenderById(context.Sender().ID)

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
	sender.TelegramId = context.Sender().ID

	storage.UpdateSender(sender)
	log.Info().Str("email", sender.Email).Int64("telegram_id", sender.TelegramId).Msg("Email saved")

	msg := "Email saved 🥳\nNow Send me book, please 🎉"

	err := context.Reply(msg)

	if err != nil {
		log.Error().Err(err)
	}
}

func RegisterNewUser(context tele.Context, storage *utils.Storage) {
	tg_sender := context.Sender()
	log.Info().Int64("telegram_id", tg_sender.ID).Str("username", tg_sender.Username).Msg("New user")
	msg := "Send me your email, please 📧"

	sender := &utils.SenderInfo{
		TelegramId: tg_sender.ID,
		UserName:   tg_sender.Username,
	}
	storage.Senders = append(storage.Senders, *sender)
	storage.WriteSendersToFile()

	err := context.Reply(msg)
	if err != nil {
		log.Error().Err(err)
	}
}
