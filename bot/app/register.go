package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/responses"
	"hushclan.com/types"
)

func (a *App) Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := a.Database.GetMember(i.Member.User.ID)
	if err == nil {
		a.RespondWithError(i, responses.ForbiddenAlreadyRegistered)
		return
	}

	member := &types.Member{
		MemberID: i.Member.User.ID,
	}

	err = a.Database.CreateMember(member)
	if err != nil {
		a.RespondWithError(i, responses.Unexpected)
		a.Log.Error("could not register", err)
		return
	}

	a.RespondWithMessage(i, responses.Register, member.MemberID)
}
