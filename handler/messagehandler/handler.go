package messagehandler

import (
	"mdhesari/discordgo-bot-kit/config"

	"github.com/bwmarrin/discordgo"
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

func (h Handler) Register(session *discordgo.Session) {
	actions := []interface{}{
		h.ReplyCommands,
	}

	for _, a := range actions {
		h.actions = append(h.actions, session.AddHandler(a))
	}
}

func (h Handler) DeRegister(session *discordgo.Session) {
	for _, remove := range h.actions {
		remove()
	}
}
