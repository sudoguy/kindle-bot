package handlers

import (
	"kindle-bot/app/mailer"
	"kindle-bot/app/settings"
	"kindle-bot/app/utils"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func DocumentHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	document := context.Message().Document
	contexted_log := log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Str("document_name", document.FileName)

	storage := utils.NewStorage()
	sender, err := storage.GetSenderById(context.Sender().ID)

	if err != nil || sender.Email == "" {
		RegisterNewUser(context, storage)
		return nil
	}

	if !slices.Contains([]string{"application/x-mobipocket-ebook", "application/epub+zip"}, document.MIME) {
		msg := "Unfortunately, I can receive only .mobi and .epub books 😭"

		err := context.Reply(msg)
		if err != nil {
			log.Error().Err(err)
		}

		return err
	}

	userId := strconv.Itoa(int(sender.TelegramId))
	_ = os.MkdirAll("books/"+userId, os.ModePerm)
	filePath := "books/" + userId + "/" + document.FileName
	err = context.Bot().Download(&document.File, filePath)

	if err != nil {
		log.Error().Err(err)
		return err
	}

	contexted_log.Msg("Download complete: " + document.FileName)

	toEmail := sender.Email
	settings := settings.NewSettings()
	mailer := mailer.NewMailer(*settings, toEmail)
	err = mailer.SendBook("books/" + userId + "/" + document.FileName)

	var msg string
	if err != nil {
		log.Error().Err(err)
		msg = "❌ An error has occurred. " + document.FileName + " was not sent"
	} else {
		contexted_log.Msg("Book " + document.FileName + " was sent to " + toEmail)
		msg = "✅ " + document.FileName + " is successfully sent"
		sender.BooksSent += 1
		storage.UpdateSender(sender)
	}

	if err := context.Reply(msg); err != nil {
		log.Err(err)
	}

	return nil

}
