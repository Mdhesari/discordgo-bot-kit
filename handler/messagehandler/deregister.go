package messagehandler

import "github.com/bwmarrin/discordgo"

func (h Handler) DeRegister(session *discordgo.Session) {
	for _, remove := range h.actions {
		remove()
	}
}
