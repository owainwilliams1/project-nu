package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/validators"
)

func (a *App) SetTeamIcon(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if !validators.ValidateURL(options[0].StringValue()) {
		a.RespondWithError(i, "Not a valid URL.")
		return
	}

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	err = a.Database.SetTeamIcon(team.TeamID, options[0].StringValue())
	if err != nil {
		a.RespondWithError(i, "Could not update team icon.")
		a.Log.Error("could not update team icon", err)
		return
	}

	a.RespondWithMessage(i, "Successfully updated team icon.")
}
