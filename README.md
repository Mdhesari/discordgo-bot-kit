# Your DiscordGo Bot Kit

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A DiscordGo bot kit for jumpstarting your Discord bot development.

## Overview

This DiscordGo bot kit provides a solid foundation for building feature-rich Discord bots in Go. With a modular structure and easy customization, it allows developers to quickly create and deploy their own bots tailored to their needs.

## Getting Started

Follow these steps to get your DiscordGo bot up and running:

```bash
git clone https://github.com/your-username/your-discordgo-bot-kit.git example
cd example
go mod tidy
```

### Run Locally

```bash
go run cmd/main.go
```

### Run With Docker

```bash
docker compose up -d
```

### Project Structure

#### cmd/main.go

Config and initialize our discord bot web socket server.

#### config/config.go, load.go

Load and create configurations like bot name and token.

#### delivery/websocketserver/server.go

Create and initialize server websocket connection.

#### handler/messagehandler, interactionhandler, \*…handler

Server handler business logic layer where we register different handlers for different sake.

## How to Contribute

Anyone and everyone is welcome to [contribute](.github/CONTRIBUTING.md). Start
by checking out the list of [open issues](https://github.com/mdhesari/discordgo-bot-kit/issues)
marked [help wanted](https://github.com/mdhesari/discordgo-bot-kit/issues?q=label:"help+wanted").
However, if you decide to get involved, please take a moment to review the
[guidelines](.github/CONTRIBUTING.md).

## License

This source code is licensed under the MIT license found in the
[LICENSE](https://opensource.org/license/mit) file.
