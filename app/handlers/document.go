package handlers

import (
	"fmt"
	"os"
	"path"
	"slices"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"

	"github.com/sudoguy/kindle-bot/app/mailer"
	"github.com/sudoguy/kindle-bot/app/settings"
)

// DocumentHandler handles the document message
func DocumentHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	document := context.Message().Document
	contextedLog := log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Str("document_name", document.FileName)

	storage := getStorage()
	sender, err := storage.GetSenderByID(context.Sender().ID)

	if err != nil || sender.Email == "" {
		RegisterNewUser(context)
		return nil
	}

	if !slices.Contains([]string{"application/x-mobipocket-ebook", "application/epub+zip"}, document.MIME) {
		msg := "Unfortunately, I can receive only .mobi and .epub books 😭"

		err := context.Reply(msg)
		if err != nil {
			log.Error().Err(err)
			return fmt.Errorf("failed to send reply: %w", err)
		}

		return err
	}

	var msg string
	if context.Chat().Type == tele.ChatPrivate {
		msg = "⏳ Received a book, working..."
		go context.Reply(msg)
	}

	filePath := path.Join(sender.Path(), document.FileName)
	err = context.Bot().Download(&document.File, filePath)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	contextedLog.Msg("Download complete: " + document.FileName)

	toEmail := sender.Email
	appSettings := settings.NewSettings()
	mail := mailer.NewMailer(*appSettings, toEmail)
	err = mail.SendBook(filePath)

	if err != nil {
		log.Error().Err(err)
		msg = "❌ An error has occurred. " + document.FileName + " was not sent"
	} else {
		contextedLog.Msg("Book " + document.FileName + " was sent to " + toEmail)
		msg = "✅ " + document.FileName + " is successfully sent"
		sender.BooksSent++
		_ = storage.UpdateSender(sender)
	}

	if err := context.Reply(msg); err != nil {
		log.Err(err)
		return fmt.Errorf("failed to send reply: %w", err)
	}

	return nil
}
