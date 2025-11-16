# Kindle Bot for Telegram

## Overview

Kindle Bot is a Telegram bot that forwards your books straight to the Kindle email assigned to your Amazon device. Upload a supported file in Telegram and the bot will relay it through the configured SMTP server.

## Getting Started

### Prerequisites

- Go 1.25 or newer.
- A Telegram account and a bot token from BotFather.
- An Amazon Kindle email address that trusts the sender address you plan to use.
- SMTP credentials for the address that will send the files.

### Configuration

Create a `.env` file in the project root (or export the variables) with the following keys:

- `BOT_TOKEN` – Telegram bot token.
- `FROM_EMAIL` – address that will send the books to Kindle.
- `EMAIL_USERNAME` / `EMAIL_PASSWORD` – login for the SMTP account.
- `SMTP_HOST` / `SMTP_PORT` – SMTP endpoint and port.

### Running locally

1. Install Go 1.25 and clone this repository.
2. `go mod download` to fetch dependencies (network access required once).
3. `go run ./app` starts the bot with the current terminal session.
4. Alternatively, `go build -o bot ./app` followed by `./bot` runs a compiled binary.

### Docker

The repository contains `docker/Dockerfile` and `docker-compose.yml`.

```bash
docker compose up --build
```

The compose file mounts the local `./storage` folder, so chats keep their state between restarts. Provide your `.env` file in the project root before launching the stack.

## Usage

Available commands registered by the bot:

- `/start` – greets the user and requests the Kindle email.
- `/email <address>` – saves or updates the Kindle email for the current chat (can also just send the email in a private chat).
- `/status` – shows the saved email and number of books already sent.

After the email is saved, send a `.mobi` or `.epub` document in a private chat. Other formats are rejected with a helpful message. Each upload is stored under `storage/<telegram_id>` and forwarded via SMTP.

## Contributing

Feel free to contribute by submitting pull requests or opening issues to suggest improvements or add new features.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
