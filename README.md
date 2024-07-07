# Kindle Bot for Telegram

## Overview

Kindle Bot is a Telegram bot designed to facilitate the transfer of books directly to your Amazon Kindle device via email. By receiving book files through Telegram messages, it offers a convenient way to send reading material to your Kindle.

## Getting Started

### Prerequisites

- A Telegram account.
- An Amazon Kindle device with an active Kindle email address.
- A bot token from BotFather for your Telegram bot.
- An email account with SMTP settings.

### Configuration

Configure the application settings through environment variables. The following variables are required:

- BOT_TOKEN: Your Telegram bot token.
- FROM_EMAIL: The email address used to send books to Kindle.
- EMAIL_USERNAME: Username for the email account.
- EMAIL_PASSWORD: Password for the email account.
- SMTP_HOST: SMTP server host for the email account.
- SMTP_PORT: SMTP server port.

### Installation

1. Clone the repository.
2. Navigate to the project directory.
3. Install dependencies.
4. Set the required environment variables as described in the Configuration section.
5. Run the bot.

## Usage

- /start: Begin interaction with the bot by sending the /start command. The bot will prompt you to send your email address to set up the forwarding of books to your Kindle.
- Send a document to the bot to have it forwarded to your Kindle email.

## Contributing

Feel free to contribute by submitting pull requests or opening issues to suggest improvements or add new features.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
