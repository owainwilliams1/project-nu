package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/validators"
)

func (a *App) SetTeamIcon(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if !validators.ValidateURL(options[0].StringValue()) {
		a.RespondWithError(i, responses.ValidationURL)
		return
	}

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	err = a.Database.SetTeamIcon(team.TeamID, options[0].StringValue())
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("could not update team icon", err)
		return
	}

	a.RespondWithMessage(i, responses.SetTeamIcon)
}
