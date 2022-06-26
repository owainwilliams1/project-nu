package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) TeamInfo(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	if len(options) > 0 {
		teamID := utils.NameToID(options[0].StringValue())
		team, err := a.Database.GetTeam(teamID)
		if err != nil {
			a.RespondWithError(i, responses.NotFoundTeam, options[0].StringValue())
			return
		}

		embed, err := a.TeamToEmbed(team)
		if err != nil {
			a.RespondWithError(i, responses.Unexpected)
			a.Log.Error("error creating team embed", err)
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
		a.RespondWithError(i, responses.ForbiddenNotMember)
		return
	}

	embed, err := a.TeamToEmbed(team)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error creating team embed", err)
		return
	}

	a.RespondWithEmbed(i, embed)
}
