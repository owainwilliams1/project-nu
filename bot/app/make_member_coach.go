package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) MakeMemberCoach(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	if len(team.Coaches) >= validators.MaxCoaches {
		a.RespondWithError(i, responses.ForbiddenMaxCoaches)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenUserNotMember)
		return
	}

	if utils.ContainsString(team.Coaches, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenAlreadyCoach)
		return
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Coach)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error making coach", err)
		return
	}

	a.RespondWithMessage(i, responses.MakeMemberCoach, options[0].UserValue(a.Session).ID)
}
