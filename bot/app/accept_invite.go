package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/utils"
)

func (a *App) AcceptInvite(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You need to `/register` first.")
		return
	}

	if member.Team != "" {
		a.RespondWithError(i, "You are already in a team.")
		return
	}

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		a.RespondWithError(i, "Team does not exist.")
		return
	}

	if !utils.ContainsString(team.Members, i.Member.User.ID) {
		a.RespondWithError(i, "You are not invited to this team.")
		return
	}

	err = a.Database.AddMemberTeam(i.Member.User.ID, team.TeamID)
	if err != nil {
		a.RespondWithError(i, "There was an error accepting this request.")
		a.Log.Error("error accepting request", err)
		return
	}

	a.RespondWithMessage(i, "Successfully accepted team invitation.")
}
