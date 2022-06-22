package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) MakeMemberSub(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	if len(team.Substitutes) >= validators.MaxSubs {
		m := fmt.Sprintf("You have the maximum number of %d subs.", validators.MaxSubs)
		a.RespondWithError(i, m)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is not in the team.")
		return
	}

	if utils.ContainsString(team.Substitutes, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is already a sub.")
		return
	}

	if utils.ContainsString(team.Players, options[0].UserValue(a.Session).ID) {
		err = a.Database.RemovePlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Player)
		if err != nil {
			a.RespondWithError(i, "There was an error removing the member from player.")
			a.Log.Error("error removing player", err)
			return
		}
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Substitute)
	if err != nil {
		a.RespondWithError(i, "There was an error making the member a sub.")
		a.Log.Error("error making sub", err)
		return
	}

	m := fmt.Sprintf("Successfully made <@%s> a sub.", options[0].UserValue(a.Session).ID)
	a.RespondWithMessage(i, m)
}
