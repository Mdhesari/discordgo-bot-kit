package messagehandler

import "github.com/bwmarrin/discordgo"

func (h Handler) Register(session *discordgo.Session) {
	actions := []interface{}{
		h.ReplyCommands,
	}

	for _, a := range actions {
		h.actions = append(h.actions, session.AddHandler(a))
	}
}
