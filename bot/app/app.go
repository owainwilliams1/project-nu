package app

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
	api "hushclan.com/api/database"
	"hushclan.com/api/logging"
	"hushclan.com/pkg/responses"
	"hushclan.com/pkg/utils"
	"hushclan.com/types"
)

type App struct {
	Session  *discordgo.Session
	Manager  *scm.SCM
	Database *api.Database
	Log      *logging.Log
	Envs     Vars
}

type Vars struct {
	Token     string
	Guild     string
	ProjectID string
	LogName   string
}

func (a *App) TeamToEmbed(team types.Team) (embed *discordgo.MessageEmbed, err error) {
	embed = &discordgo.MessageEmbed{
		Title:       team.TeamName,
		Description: strings.Join([]string{team.Game, team.Sex, team.Region}, " / "),
	}

	fields := []*discordgo.MessageEmbedField{}
	for _, member := range team.Members {
		memberType, hasType := team.GetMemberType(member)
		if !hasType {
			continue
		}

		discordUser, err := a.Session.User(member)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   discordUser.Username,
			Value:  strings.Join(memberType, " / "),
			Inline: true,
		})
	}

	embed.Fields = fields

	if team.Icon != "" {
		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: team.Icon,
		}
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: team.TeamID,
	}

	embed.Color = team.Color

	return
}

func (a *App) RespondWithMessage(
	i *discordgo.InteractionCreate,
	m responses.Success,
	s ...string,
) {
	opt := utils.SAtoIA(s)
	err := a.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: fmt.Sprintf(string(m), opt...),
		},
	})
	if err != nil {
		a.Log.Error("could not respond to command", err)
	}
}

func (a *App) RespondWithError(
	i *discordgo.InteractionCreate,
	m responses.Failure,
	s ...string,
) {
	opt := utils.SAtoIA(s)
	err := a.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: fmt.Sprintf(string(":x: "+m), opt...),
		},
	})
	if err != nil {
		a.Log.Error("could not respond to command", err)
	}
}

func (a *App) RespondWithEmbed(
	i *discordgo.InteractionCreate,
	e *discordgo.MessageEmbed,
) {
	err := a.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "",
			Embeds:  []*discordgo.MessageEmbed{e},
		},
	})
	if err != nil {
		a.Log.Error("could not respond to command", err)
	}
}
