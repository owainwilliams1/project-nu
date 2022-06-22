package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) MakeMemberCoach(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	if len(team.Coaches) >= validators.MaxCoaches {
		m := fmt.Sprintf("You have the maximum number of %d coaches.", validators.MaxCoaches)
		a.RespondWithError(i, m)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is not in the team.")
		return
	}

	if utils.ContainsString(team.Coaches, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "Member is already a coach.")
		return
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Coach)
	if err != nil {
		a.RespondWithError(i, "There was an error making the member a coach.")
		a.Log.Error("error making coach", err)
		return
	}

	m := fmt.Sprintf("Successfully made <@%s> a coach.", options[0].UserValue(a.Session).ID)
	a.RespondWithMessage(i, m)
}
