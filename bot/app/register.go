package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"hushclan.com/types"
)

func (a *App) Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	_, err := a.Database.GetMember(i.Member.User.ID)
	if err == nil {
		a.RespondWithError(i, "You are already registered.")
		return
	}

	member := &types.Member{
		MemberID: i.Member.User.ID,
		Username: options[0].StringValue(),
	}

	err = a.Database.CreateMember(member)
	if err != nil {
		a.RespondWithError(i, "Could not register.")
		a.Log.Error("could not register", err)
		return
	}

	m := fmt.Sprintf("Welcome, you have been registered!")
	a.RespondWithMessage(i, m)
}
