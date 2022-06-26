package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
)

func (a *App) DeleteTeam(s *discordgo.Session, i *discordgo.InteractionCreate) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	err = a.Database.DeleteTeam(team.TeamID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error deleting team", err)
		return
	}

	a.RespondWithMessage(i, responses.DeleteTeam)
}
