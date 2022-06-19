package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/pkg/discordutils"
	"hushclan.com/types"
)

func (a *App) Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	_, err := a.Database.GetMember(i.Member.User.ID)
	if err == nil {
		discordutils.RespondWithError(s, i, "You are already registered.")
		return
	}

	member := &types.Member{
		MemberID: i.Member.User.ID,
		Username: options[0].StringValue(),
	}

	err = a.Database.CreateMember(member)
	if err != nil {
		discordutils.RespondWithError(s, i, "Could not register.")
		return
	}

	m := fmt.Sprintf("Welcome, you have been registered!")
	discordutils.RespondWithMessage(s, i, m)
}
