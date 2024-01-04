package messagehandler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) ReplyCommands(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if msg.Author.ID == s.State.User.ID {
		return
	}

	// Just sample command check
	if msg.Content == "/ping" {
		_, err := s.ChannelMessageSend(msg.ChannelID, "pong "+h.config.Name)
		if err != nil {
			log.Println(err)
		}
	}
}
