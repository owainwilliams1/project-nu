package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) ManageAddPlayer(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	if len(team.Players) >= validators.MaxPlayers {
		a.RespondWithError(i, responses.ForbiddenMaxPlayers)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenUserNotMember)
		return
	}

	if utils.ContainsString(team.Players, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenAlreadyPlayer)
		return
	}

	member, err := a.Database.GetMember(options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenUserNotJoined)
		return
	}

	if member.Team != team.TeamID {
		a.RespondWithError(i, responses.ForbiddenUserNotJoined)
		return
	}

	if utils.ContainsString(team.Substitutes, options[0].UserValue(a.Session).ID) {
		err = a.Database.RemovePlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Substitute)
		if err != nil {
			a.RespondWithError(i, responses.Unexpected)
			a.Log.Error("error removing sub", err)
			return
		}
	}

	err = a.Database.AddPlayerType(team.TeamID, options[0].UserValue(a.Session).ID, database.Player)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error making player", err)
		return
	}

	a.RespondWithMessage(i, responses.MakeMemberPlayer, options[0].UserValue(a.Session).ID)
}
