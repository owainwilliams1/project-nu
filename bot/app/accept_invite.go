package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
)

func (a *App) AcceptInvite(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You need to `/register` first.")
		return
	}

	if member.Team != "" {
		discordutils.RespondWithError(s, i, "You are already in a team.")
		return
	}

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		discordutils.RespondWithError(s, i, "Team does not exist.")
		return
	}

	if !utils.ContainsString(team.Members, i.Member.User.ID) {
		discordutils.RespondWithError(s, i, "You are not invited to this team.")
		return
	}

	err = a.Database.AddMemberTeam(i.Member.User.ID, team.TeamID)
	if err != nil {
		discordutils.RespondWithError(s, i, "There was an error accepting this request.")
		return
	}

	discordutils.RespondWithMessage(s, i, "Successfully accepted team invitation.")
}
