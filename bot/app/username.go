package app

import (
	"github.com/bwmarrin/discordgo"
	"hushclan.com/api/database"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/validators"
)

func (a *App) UsernameProjectNu(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	if !validators.ValidateUsername(options[0].StringValue()) {
		a.RespondWithError(i, responses.ValidationURL)
		return
	}

	err := a.Database.SetUsername(i.Member.User.ID, options[0].StringValue(), database.ProjectNu)
	if err != nil {
		a.RespondWithError(i, responses.RequireRegistration)
		return
	}

	a.RespondWithMessage(i, responses.UsernameSet)
}

func (a *App) UsernameValorant(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	if !validators.ValidateUsername(options[0].StringValue()) {
		a.RespondWithError(i, responses.ValidationURL)
		return
	}

	err := a.Database.SetUsername(i.Member.User.ID, options[0].StringValue(), database.Valorant)
	if err != nil {
		a.RespondWithError(i, responses.RequireRegistration)
		return
	}

	a.RespondWithMessage(i, responses.UsernameSet)
}

func (a *App) UsernameApex(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	if !validators.ValidateUsername(options[0].StringValue()) {
		a.RespondWithError(i, responses.ValidationURL)
		return
	}

	err := a.Database.SetUsername(i.Member.User.ID, options[0].StringValue(), database.ApexLegends)
	if err != nil {
		a.RespondWithError(i, responses.RequireRegistration)
		return
	}

	a.RespondWithMessage(i, responses.UsernameSet)
}
