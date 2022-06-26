package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
	"hushclan.com/pkg/validators"
	"hushclan.com/types"
)

func (a *App) TeamCreate(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	if !validators.ValidateTeamName(options[0].StringValue()) {
		a.RespondWithError(i, responses.ValidationTeamName)
		return
	}

	color, ok := validators.ValidateHexHashtag(options[1].StringValue())
	if !ok {
		a.RespondWithError(i, responses.ValidationHex, options[1].StringValue())
		return
	}

	member, err := a.Database.GetMember(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.RequireRegistration)
		return
	}

	if member.Team != "" {
		a.RespondWithError(i, responses.ForbiddenAlreadyMember)
		return
	}

	team := &types.Team{
		OwnerID:  i.Member.User.ID,
		TeamID:   utils.NameToID(options[0].StringValue()),
		TeamName: options[0].StringValue(),
		Color:    color,
		Region:   options[2].StringValue(),
		Sex:      options[3].StringValue(),
		Game:     options[4].StringValue(),
	}
	err = a.Database.CreateTeam(team)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenTeamExists, team.TeamName)
		return
	}

	dmChannel, err := s.UserChannelCreate(i.Member.User.ID)
	if err != nil {
		a.Log.Error("could not create dm with user", err)
	} else {
		_, err = s.ChannelMessageSendEmbed(dmChannel.ID, a.TeamCreatedEmbed())
		if err != nil {
			a.Log.Error("could not send embed to user", err)
		}
	}

	a.RespondWithMessage(i, responses.CreateTeam, team.TeamID)
}
