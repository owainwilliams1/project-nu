package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
)

func (a *App) DeleteTeam(s *discordgo.Session, i *discordgo.InteractionCreate) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You are not in a team or you do not own the team.")
		return
	}

	err = a.Database.DeleteTeam(team.TeamID)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error deleting the team.")
		utils.LogError("error deleting team", err)
		return
	}

	discordutils.RespondWithMessage(s, i, "Successfully deleted your team.")
}
