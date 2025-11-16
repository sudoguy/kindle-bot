package handlers

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v4"
)

const instructionsMessage = `Here is how to use Kindle Bot:

• Send /email <your_kindle_email> (or just the email in a private chat) to link your Kindle address.
• Upload .mobi or .epub documents. Other formats are rejected.
• Use /status any time to see the saved email and the number of books delivered.

Commands:
/start  – greet the bot and start onboarding.
/email  – set or update the Kindle email.
/status – show the saved email and statistics.
/help   – show these instructions again.`

// HelpHandler sends usage instructions
func HelpHandler(context tele.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Int64("telegram_id", context.Sender().ID).Str("username", context.Sender().Username).Msg("Help command")

	return context.Reply(instructionsMessage)
}
