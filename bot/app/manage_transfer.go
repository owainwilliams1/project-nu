package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) ManageTransferOwnership(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	if !utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenUserNotMember)
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

	err = a.Database.TransferOwnership(team.TeamID, options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error transferring ownership", err)
		return
	}

	a.RespondWithMessage(i, responses.TransferOwnership, options[0].UserValue(a.Session).ID)
}
