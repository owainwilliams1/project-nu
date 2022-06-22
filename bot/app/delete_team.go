package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) DeleteTeam(s *discordgo.Session, i *discordgo.InteractionCreate) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	err = a.Database.DeleteTeam(team.TeamID)
	if err != nil {
		a.RespondWithError(i, "There was an error deleting the team.")
		a.Log.Error("error deleting team", err)
		return
	}

	a.RespondWithMessage(i, "Successfully deleted your team.")
}
