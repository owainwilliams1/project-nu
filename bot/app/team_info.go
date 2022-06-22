package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) TeamInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		a.RespondWithError(i, responses.NotFoundTeam, options[0].StringValue())
		return
	}

	embed, err := a.TeamToEmbed(team)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error retrieving team", err)
		return
	}

	a.RespondWithEmbed(i, embed)
}
