package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/types"
)

func (a *App) Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	_, err := a.Database.GetMember(i.Member.User.ID)
	if err == nil {
		a.RespondWithError(i, responses.ForbiddenAlreadyRegistered)
		return
	}

	member := &types.Member{
		MemberID: i.Member.User.ID,
		Username: options[0].StringValue(),
	}

	err = a.Database.CreateMember(member)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("could not register", err)
		return
	}

	a.RespondWithMessage(i, responses.Register, member.MemberID)
}
