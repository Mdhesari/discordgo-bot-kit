package interactionhandler

import (
	"mdhesari/discordgo-bot-kit/config"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	config     *config.Discord
	handlers   []func() // we need handlers func in order to remove them later
	commandIDs []string // we need command ids in order to remove them later
}

func New(cfg *config.Discord) *Handler {
	return &Handler{
		config: cfg,
	}
}

func (h Handler) Register(session *discordgo.Session) {
	h.TearUpCommands(session)

	actions := []interface{}{
		h.Ping,
	}

	for _, a := range actions {
		h.handlers = append(h.handlers, session.AddHandler(a))
	}
}

func (h Handler) DeRegister(session *discordgo.Session) {
	h.TearDownCommands(session)

	for _, remove := range h.handlers {
		remove()
	}
}
