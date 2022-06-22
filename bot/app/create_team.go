package app

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
	"hushclan.com/types"
)

func (a *App) CreateTeam(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if !validators.ValidateTeamName(options[0].StringValue()) {
		a.RespondWithError(i, "Team names must only contain letters and spaces. They cannot be longer than 24 characters.")
		return
	}

	color, ok := validators.ValidateHexWithHashtag(options[1].StringValue())
	if !ok {
		a.RespondWithError(i, "That is not a valid hex string e.g. `#00FF66`.")
		return
	}

	region, ok := validators.ValidateRegion(options[2].StringValue())
	if !ok {
		m := fmt.Sprintf("That is not a valid region. Choose from `%s`", strings.Join(validators.Regions[:], ", "))
		a.RespondWithError(i, m)
		return
	}

	sex, ok := validators.ValidateSex(options[3].StringValue())
	if !ok {
		m := fmt.Sprintf("That is not a valid sex. Choose from `%s`", strings.Join(validators.Sexes[:], ", "))
		a.RespondWithError(i, m)
		return
	}

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, "You need to `/register` first.")
		return
	}

	if member.Team != "" {
		a.RespondWithError(i, "You are already in a team.")
		return
	}

	team := &types.Team{
		OwnerID:  i.Member.User.ID,
		TeamID:   utils.NameToID(options[0].StringValue()),
		TeamName: options[0].StringValue(),
		Color:    color,
		Game:     "Valorant",
		Sex:      sex,
		Region:   region,
	}
	err = a.Database.CreateTeam(team)
	if err != nil {
		a.RespondWithError(i, "A team with this name already exists.")
		return
	}

	m := fmt.Sprintf("Team `%s` created!", team.TeamID)
	a.RespondWithMessage(i, m)
}
