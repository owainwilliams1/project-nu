package app

import (
	"errors"
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
	Session      *discordgo.Session
	Manager      *scm.SCM
	Database     *api.Database
	Log          *logging.Log
	Envs         Vars
	JoinedGuilds []string
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
		Description: strings.Join([]string{team.Game, team.Sex, team.Region}, " "),
	}

	fields := []*discordgo.MessageEmbedField{}
	for _, member := range team.Members {
		memberType := team.GetMemberType(member)
		discordUser, err := a.Session.User(member)
		if err != nil {
			return nil, errors.New("could not get discord user")
		}
		memberData, err := a.Database.GetMember(member)
		if err != nil {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   discordUser.Username,
				Value:  "**Invited**",
				Inline: true,
			})
			continue
		}

		if memberData.Team != team.TeamID {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   memberData.Username,
				Value:  "*Invited*",
				Inline: true,
			})
			continue
		}

		if len(memberType) > 0 {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   memberData.Username,
				Value:  strings.Join(memberType, "\n"),
				Inline: true,
			})
			continue
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   memberData.Username,
			Value:  "Member",
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
		Text: fmt.Sprintf("Team ID: %s", team.TeamID),
	}

	embed.Color = team.Color

	return
}

func (a *App) TeamCreatedEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Welcome to team management!",
		Description: "Here are a few commands to help you get started. For more commands use `/help command:manage`. " +
			"Don't use these commands in DMs as they will not work",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "/manage invite",
				Value: "Invite a member to your team.",
			},
			{
				Name:  "/manage add",
				Value: "Give a player a role on the team.",
			},
			{
				Name:  "/manage set-icon",
				Value: "Give your team a beautiful icon.",
			},
		},
	}
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
