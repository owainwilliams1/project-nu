package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
)

func (a *App) Leaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	page := 0
	if options != nil {
		page = int(options[0].IntValue())
	}

	teams, err := a.Database.GetTeams(5, page)
	if err != nil {
		a.RespondWithError(i, responses.NotFoundTeams)
		return
	}

	fields := []*discordgo.MessageEmbedField{}
	for i, team := range teams {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("Team %d", i),
			Value: team.TeamID,
		})
	}
	embed := &discordgo.MessageEmbed{
		Title:  "Leaderboard",
		Fields: fields,
	}
	a.RespondWithEmbed(i, embed)
}
