package main

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/sudoguy/kindle-bot/app/handlers"
	"github.com/sudoguy/kindle-bot/app/settings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().Msg("Starting an app..")

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg(".env file not found")
	}

	appSettings := settings.NewSettings()
	if err := appSettings.Validate(); err != nil {
		log.Fatal().Err(err).Msg("Invalid settings")
	}

	pref := tele.Settings{
		Token:  appSettings.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create a bot")
		return
	}

	log.Info().Msg("Authorized on account " + bot.Me.Username)
	err = bot.RemoveWebhook()
	if err != nil {
		log.Info().Msg("Webhook is not set")
	}

	bot.Handle("/start", handlers.StartHandler)
	bot.Handle("/status", handlers.StatusHandler)
	bot.Handle(tele.OnText, handlers.TextHandler)
	bot.Handle(tele.OnDocument, handlers.DocumentHandler)
	err = bot.SetCommands([]tele.Command{
		{Text: "start", Description: "Start working with bot"},
		// {Text: "status", Description: "Show current status"},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to set commands")
	}

	bot.Start()
}
