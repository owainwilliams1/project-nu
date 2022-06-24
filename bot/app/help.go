package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (a *App) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	e := &discordgo.MessageEmbed{
		Title: "Here are a list of some commands. Use `help [page]` to see the next pages.",
	}

	page := 0
	if options != nil {
		page = int(options[0].IntValue()) - 1
	}

	pageSize := 5
	start := pageSize * page
	end := pageSize*page + pageSize

	features := a.GetFeatures()
	featuresSub := features[start:end]
	for _, feature := range featuresSub {
		e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
			Name:  feature.ApplicationCommand.Name,
			Value: feature.ApplicationCommand.Description,
		})
	}

	totalPages := len(features) / 5
	if r := totalPages % 1; r > 0 {
		totalPages -= r
		totalPages += 1
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("Page %d/%d", page+1, totalPages),
	}

	a.RespondWithEmbed(i, e)
}
