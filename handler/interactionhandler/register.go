package interactionhandler

import "github.com/bwmarrin/discordgo"


// When we open the websocket we should make sure we have initialized commands and integrated our handlers
func (h Handler) Register(session *discordgo.Session) {
	h.TearUpCommands(session)

	actions := []interface{}{
		h.Ping,
	}

	for _, a := range actions {
		h.handlers = append(h.handlers, session.AddHandler(a))
	}
}
