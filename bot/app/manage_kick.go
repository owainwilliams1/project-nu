package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
)

func (a *App) ManageKick(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	if team.OwnerID == options[0].UserValue(a.Session).ID {
		a.RespondWithError(i, responses.ForbiddenOwnerAction)
		return
	}

	a.Database.RemoveTeamMember(team.TeamID, options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error removing user from team", err)
		return
	}

	a.Database.RemoveMemberTeam(options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error removing user from team", err)
		return
	}

	a.RespondWithMessage(i, responses.RemoveMember, options[0].UserValue(a.Session).ID)
}
