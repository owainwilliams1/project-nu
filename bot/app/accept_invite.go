package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) AcceptInvite(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.RequireRegistration)
		return
	}

	if member.Team != "" {
		a.RespondWithError(i, responses.ForbiddenAlreadyMember)
		return
	}

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		a.RespondWithError(i, responses.NotFoundTeam, team.TeamID)
		return
	}

	if !utils.ContainsString(team.Members, i.Member.User.ID) {
		a.RespondWithError(i, responses.ForbiddenNoInvite, team.TeamID)
		return
	}

	err = a.Database.AddMemberTeam(i.Member.User.ID, team.TeamID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error accepting request", err)
		return
	}

	a.RespondWithMessage(i, responses.AcceptInvite, team.TeamName)
}
