package discordutils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RespondWithMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	m string,
) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: m,
		},
	})
	if err != nil {
		log.Print("could not respond to command", err)
	}
}

func RespondWithError(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	m string,
) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: ":x: " + m,
		},
	})
	if err != nil {
		log.Print("could not respond to command", err)
	}
}

func RespondWithEmbed(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	e *discordgo.MessageEmbed,
) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "",
			Embeds:  []*discordgo.MessageEmbed{e},
		},
	})
	if err != nil {
		log.Print("could not respond to command", err)
	}
}
