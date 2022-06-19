package app

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/pkg/utils"
)

func (a *App) TeamInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	teamID := utils.NameToID(options[0].StringValue())
	team, err := a.Database.GetTeam(teamID)
	if err != nil {
		log.Print(err)
		m := fmt.Sprintf("`%s` team does not exist.", options[0].StringValue())
		discordutils.RespondWithError(s, i, m)
		return
	}

	embed, err := a.TeamToEmbed(team)
	if err != nil {
		discordutils.RespondWithError(s, i, "Error retrieving team.")
		return
	}

	discordutils.RespondWithEmbed(s, i, embed)
}
