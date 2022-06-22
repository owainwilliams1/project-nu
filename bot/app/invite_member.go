package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/utils"
)

func (a *App) InviteMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team or you do not own the team.")
		return
	}

	if utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, "User has already been invited to the team.")
		return
	}

	err = a.Database.AddTeamMember(team.TeamID, options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, "There was an error inviting the user.")
		a.Log.Error("error inviting member", err)
		return
	}

	m := fmt.Sprintf("Successfully invited <@%s>.", options[0].UserValue(a.Session).ID)
	a.RespondWithMessage(i, m)
}
