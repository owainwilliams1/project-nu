package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
)

func (a *App) InviteMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	team, err := a.Database.GetTeamByOwner(i.Member.User.ID)
	if err != nil {
		a.RespondWithError(i, responses.ForbiddenNotOwner)
		return
	}

	if utils.ContainsString(team.Members, options[0].UserValue(a.Session).ID) {
		a.RespondWithError(i, responses.ForbiddenAlreadyInvited, options[0].UserValue(a.Session).ID)
		return
	}

	messageGuild, err := s.Guild(i.GuildID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("guild not found", err)
		return
	}

	dmChannel, err := s.UserChannelCreate(options[0].UserValue(a.Session).ID)
	m := fmt.Sprintf("You have been invited to join team `%[1]s`, please type "+
		"`/accept-invite %[1]s` in %[2]s to join the team! You can view the team's info with"+
		"`/team-info %[1]s`.", team.TeamID, messageGuild.Name)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("could not create dm with user", err)
		return
	}

	_, err = s.ChannelMessageSend(dmChannel.ID, m)
	if err != nil {
		a.RespondWithError(i, responses.NotFoundUser)
		return
	}

	err = a.Database.AddTeamMember(team.TeamID, options[0].UserValue(a.Session).ID)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("error inviting member", err)
		return
	}

	a.RespondWithMessage(i, responses.InviteMember, options[0].UserValue(a.Session).ID, team.TeamID)
}
