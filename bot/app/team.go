package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
)

func (a *App) Team(s *discordgo.Session, i *discordgo.InteractionCreate) {
	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		discordutils.RespondWithError(s, i, "You are not in a team.")
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if err != nil {
		discordutils.RespondWithError(s, i, "Error retrieving team.")
		return
	}

	embed, err := a.TeamToEmbed(team)
	if err != nil {
		discordutils.RespondWithError(s, i, "Error retrieving team.")
		return
	}

	discordutils.RespondWithEmbed(s, i, embed)
}
