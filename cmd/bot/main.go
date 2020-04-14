package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/imroc/req"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Starting an app..")
	token := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic().Err(err)
	}

	bot.Debug = os.Getenv("DEBUG") == "1"

	log.Info().Msg("Authorized on account " + bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Error().Err(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Document == nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Send me book, please 🥳")
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				log.Error().Err(err)
			}

			continue
		}

		document := update.Message.Document

		if document.MimeType != "application/x-mobipocket-ebook" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unfortunately, I can receive only .mobi books 😭")
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				log.Error().Err(err)
			}

			continue
		}

		fileUrl, err := bot.GetFileDirectURL(document.FileID)
		if err != nil {
			log.Error().Err(err)
			continue
		}

		r, _ := req.Get(fileUrl)
		userId := strconv.Itoa(update.Message.From.ID)
		_ = os.MkdirAll(userId, os.ModePerm)

		err = r.ToFile(userId + "/" + document.FileName)
		if err != nil {
			log.Error().Err(err)
			continue
		} else {
			log.Info().Msg("Download complete: " + document.FileName)
		}

		from := os.Getenv("FROM_EMAIL")
		to := os.Getenv("TO_EMAIL")

		emailUsername := os.Getenv("EMAIL_USERNAME")
		emailPassword := os.Getenv("EMAIL_PASSWORD")

		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", to)
		m.SetHeader("Subject", "New book from Awesome Kindle Bot!")
		m.Attach(userId + "/" + document.FileName)

		d := gomail.NewDialer(smtpHost, smtpPort, emailUsername, emailPassword)

		if err := d.DialAndSend(m); err != nil {
			log.Error().Err(err)
		} else {
			log.Info().Msg("Book " + document.FileName + " was sent to " + to)
		}

	}
}
