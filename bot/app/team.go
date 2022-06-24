package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) Team(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if options != nil {
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
		return
	}

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotMember)
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error retrieving team", err)
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
