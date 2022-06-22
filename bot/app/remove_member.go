package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (a *App) RemoveMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	if team.OwnerID == options[0].UserValue(a.Session).ID {
		a.RespondWithError(i, "You cannot remove yourself from a team, please transfer ownership first.")
		return
	}

	a.Database.RemoveTeamMember(team.TeamID, options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, "There was an error removing the user from the team.")
		a.Log.Error("error removing user from team", err)
		return
	}

	a.Database.RemoveMemberTeam(options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, "There was an error removing the user from the team.")
		a.Log.Error("error removing user from team", err)
		return
	}

	m := fmt.Sprintf("Successfully removed <@%s> from the team.", options[0].UserValue(a.Session).ID)
	a.RespondWithMessage(i, m)
}
