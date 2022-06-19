package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
)

func (a *App) RemoveMemberPlayer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You are not in a team or you do not own the team.")
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		discordutils.RespondWithError(s, i, "Member is not in the team.")
		return
	}

	if !utils.ContainsString(team.Players, options[0].UserValue(a.Session).ID) {
		discordutils.RespondWithError(s, i, "Member is not a player already.")
		return
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Player)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error removing the player.")
		return
	}

	m := fmt.Sprintf("Successfully removed <@%s> from player.", options[0].UserValue(a.Session).ID)
	discordutils.RespondWithMessage(s, i, m)
}
