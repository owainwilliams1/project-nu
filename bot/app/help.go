package app

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (a *App) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	e := &discordgo.MessageEmbed{
		Title: strings.Title(options[0].StringValue()),
	}

	features := a.GetFeatures()

	fields := []*discordgo.MessageEmbedField{}
	switch options[0].StringValue() {
	case "general":
		for _, feature := range features {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  feature.ApplicationCommand.Name,
				Value: feature.ApplicationCommand.Description,
			})
		}
	default:
		for _, feature := range features {
			if feature.ApplicationCommand.Name == options[0].StringValue() {
				for _, subCommand := range feature.ApplicationCommand.Options {
					fields = append(fields, &discordgo.MessageEmbedField{
						Name:  subCommand.Name,
						Value: subCommand.Description,
					})
				}
			}
		}
	}

	e.Fields = fields

	a.RespondWithEmbed(i, e)
}
