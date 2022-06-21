package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
)

func (a *App) Leave(s *discordgo.Session, i *discordgo.InteractionCreate) {
	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil || member.Team == "" {
		discordutils.RespondWithError(s, i, "You are not in a team.")
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if team.OwnerID == i.Member.User.ID {
		discordutils.RespondWithError(s, i, "You cannot leave your team, please transfer ownership first.")
		return
	}

	a.Database.RemoveTeamMember(member.Team, i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error leaving the team.")
		utils.LogError("error leaving team", err)
		return
	}

	a.Database.RemoveMemberTeam(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error leaving the team.")
		utils.LogError("error leaving team", err)
		return
	}

	discordutils.RespondWithMessage(s, i, "Successfully left the team.")
}
