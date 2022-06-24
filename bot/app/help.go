package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
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

	features := a.GetFeatures()
	totalLen := len(features)

	pageSize := 5
	start := pageSize * page

	if start > totalLen {
		a.RespondWithError(i, responses.ForbiddenPageOutOfRange)
		return
	}

	left := totalLen - start
	var end int
	if left > pageSize {
		end = pageSize*page + pageSize
	} else {
		end = pageSize*page + left
	}

	featuresSub := features[start:end]
	for _, feature := range featuresSub {
		e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
			Name:  feature.ApplicationCommand.Name,
			Value: feature.ApplicationCommand.Description,
		})
	}

	totalPages := totalLen / 5
	if r := totalPages % 1; r > 0 {
		totalPages -= r
		totalPages += 1
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("Page %d/%d", page+1, totalPages),
	}

	a.RespondWithEmbed(i, e)
}
