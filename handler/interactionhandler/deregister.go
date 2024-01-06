package interactionhandler

import "github.com/bwmarrin/discordgo"

// When we close the websocket we should make sure everything is cleaned up and removed
func (h Handler) DeRegister(session *discordgo.Session) {
	h.TearDownCommands(session)

	for _, remove := range h.handlers {
		remove()
	}
}
