package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
)

func (a *App) SetTeamIcon(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if !validators.ValidateURL(options[0].StringValue()) {
		discordutils.RespondWithError(s, i, "Not a valid URL.")
		return
	}

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You are not in a team or you do not own the team.")
		return
	}

	err = a.Database.SetTeamIcon(team.TeamID, options[0].StringValue())
	if err != nil {
		discordutils.RespondWithError(s, i, "Could not update team icon.")
		utils.LogError("could not update team icon", err)
		return
	}

	discordutils.RespondWithMessage(s, i, "Successfully updated team icon.")
}
