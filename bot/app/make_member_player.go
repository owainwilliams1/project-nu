package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) MakeMemberPlayer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You are not in a team or you do not own the team.")
		return
	}

	if len(team.Players) >= validators.MaxPlayers {
		m := fmt.Sprintf("You have the maximum number of %d players.", validators.MaxPlayers)
		discordutils.RespondWithError(s, i, m)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		discordutils.RespondWithError(s, i, "Member is not in the team.")
		return
	}

	if utils.ContainsString(team.Players, options[0].UserValue(a.Session).ID) {
		discordutils.RespondWithError(s, i, "Member is already a player.")
		return
	}

	if utils.ContainsString(team.Substitutes, options[0].UserValue(a.Session).ID) {
		err = a.Database.RemovePlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Substitute)
		if err != nil {
			discordutils.RespondWithError(s, i, "There was an error removing the member from sub.")
			utils.LogError("error removing sub", err)
			return
		}
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Player)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error making the member a player.")
		utils.LogError("error making player", err)
		return
	}

	m := fmt.Sprintf("Successfully made <@%s> a player.", options[0].UserValue(a.Session).ID)
	discordutils.RespondWithMessage(s, i, m)
}
