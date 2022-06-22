package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) Team(s *discordgo.Session, i *discordgo.InteractionCreate) {
	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You are not in a team.")
		return
	}

	team, err := a.Database.GetTeam(member.Team)
	if err != nil {
		a.RespondWithError(i, "Error retrieving team.")
		a.Log.Error("error retrieving team", err)
		return
	}

	embed, err := a.TeamToEmbed(team)
	if err != nil {
		a.RespondWithError(i, "Error retrieving team.")
		a.Log.Error("error retrieving team", err)
		return
	}

	a.RespondWithEmbed(i, embed)
}
