package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
)

func (a *App) TeamLeave(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil || member.Team == "" {
		a.RespondWithError(i, responses.ForbiddenNotMember)
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if team.OwnerID == i.Member.User.ID {
		a.RespondWithError(i, responses.ForbiddenOwnerAction)
		return
	}

	a.Database.RemoveTeamMember(member.Team, i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error leaving team", err)
		return
	}

	a.Database.RemoveMemberTeam(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error leaving team", err)
		return
	}

	a.RespondWithMessage(i, responses.Leave)
}
