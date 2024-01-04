package interactionhandler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) TearUpCommands(session *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot.",
		},
	}

	for _, c := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, h.config.GuildID, c)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}

		h.commandIDs = append(h.commandIDs, cmd.ID)
	}
}

func (h Handler) TearDownCommands(session *discordgo.Session) {
	for _, cmd := range h.commandIDs {
		err := session.ApplicationCommandDelete(session.State.User.ID, h.config.GuildID, cmd)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", cmd, err)
		}
	}
}
