package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/utils"
)

func (a *App) RemoveMemberCoach(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is not in the team.")
		return
	}

	if !utils.ContainsString(team.Coaches, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is not a coach already.")
		return
	}

	err = a.Database.RemovePlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Coach)
	if err != nil {
		a.RespondWithError(i, "There was an error removing the coach.")
		a.Log.Error("error removing coach", err)
		return
	}

	m := fmt.Sprintf("Successfully removed <@%s> from coach.", options[0].UserValue(a.Session).ID)
	a.RespondWithMessage(i, m)
}
