package interactionhandler

import "github.com/bwmarrin/discordgo"

func (h Handler) Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "ping" {

		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong " + h.config.Name,
		},
	})
}
