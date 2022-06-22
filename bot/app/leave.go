package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) Leave(s *discordgo.Session, i *discordgo.InteractionCreate) {
	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil || member.Team == "" {
		a.RespondWithError(i, "You are not in a team.")
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if team.OwnerID == i.Member.User.ID {
		a.RespondWithError(i, "You cannot leave your team, please transfer ownership first.")
		return
	}

	a.Database.RemoveTeamMember(member.Team, i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "There was an error leaving the team.")
		a.Log.Error("error leaving team", err)
		return
	}

	a.Database.RemoveMemberTeam(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "There was an error leaving the team.")
		a.Log.Error("error leaving team", err)
		return
	}

	a.RespondWithMessage(i, "Successfully left the team.")
}
