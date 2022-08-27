package main

import (
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/imroc/req"
	"golang.org/x/exp/slices"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
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

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")

	webhookHost := "https://" + appName + ".fly.dev/" + bot.Token
	log.Info().Msg("Webhook host: " + webhookHost)

	resp, err := bot.SetWebhook(tgbotapi.NewWebhook(webhookHost))
	if err != nil {
		log.Panic().Err(err)
	} else {
		log.Info().Msg(resp.Description)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+port, nil) //nolint:errcheck

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

		if !slices.Contains([]string{"application/x-mobipocket-ebook", "application/epub+zip"}, document.MimeType) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unfortunately, I can receive only .mobi and .epub books 😭")
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
		_ = os.MkdirAll("books/"+userId, os.ModePerm)

		err = r.ToFile("books/" + userId + "/" + document.FileName)
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

		m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
		m.SetHeader("From", from)
		m.SetHeader("To", to)
		m.SetHeader("Subject", "New book from Awesome Kindle Bot!")
		m.SetBody("text/html", "Get your book!")
		m.Attach("books/" + userId + "/" + document.FileName)

		d := gomail.NewDialer(smtpHost, smtpPort, emailUsername, emailPassword)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if err := d.DialAndSend(m); err != nil {
			log.Error().Err(err)
			msg.Text = "❌ An error has occurred. " + document.FileName + " was not sent"
		} else {
			log.Info().Msg("Book " + document.FileName + " was sent to " + to)
			msg.Text = "✅ " + document.FileName + " is successfully sent"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Err(err)
		}

	}
}
