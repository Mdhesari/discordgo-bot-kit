package messagehandler

import (
	"mdhesari/discordgo-bot-kit/config"
)

type Handler struct {
	config  *config.Discord
	actions []func()
}

func New(cfg *config.Discord) *Handler {
	return &Handler{
		config: cfg,
	}
}
