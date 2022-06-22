package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/utils"
)

func (a *App) TeamInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		m := fmt.Sprintf("`%s` team does not exist.", options[0].StringValue())
		a.RespondWithError(i, m)
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
