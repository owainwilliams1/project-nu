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
		Description: strings.Title(strings.Join([]string{team.Game, team.Sex, team.Region}, " ")),
	}

	players := []string{}
	substitutes := []string{}
	coaches := []string{}
	invited := []string{}
	unassigned := []string{}

	for _, member := range team.Members {
		memberType := team.GetMemberType(member)
		assigned := false

		memberData, err := a.Database.GetMember(member)
		if err != nil {
			assigned = true
			invited = append(invited, member)
		}

		if memberData.Team != team.TeamID {
			assigned = true
			invited = append(invited, member)
		}

		if utils.ContainsString(memberType, "Player") {
			assigned = true
			players = append(players, member)
		}

		if utils.ContainsString(memberType, "Sub") {
			assigned = true
			substitutes = append(substitutes, member)
		}

		if utils.ContainsString(memberType, "Coach") {
			assigned = true
			coaches = append(coaches, member)
		}

		if !assigned && member != team.OwnerID {
			unassigned = append(unassigned, member)
		}
	}

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "Manager",
		Value: fmt.Sprintf("<@%s>", team.OwnerID),
	})

	if len(players) > 0 {
		mentions := "<@" + strings.Join(players, ">\n<@") + ">"
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Players",
			Value: mentions,
		})
	}

	if len(substitutes) > 0 {
		mentions := "<@" + strings.Join(substitutes, ">\n<@") + ">"
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Substitutes",
			Value: mentions,
		})
	}

	if len(coaches) > 0 {
		mentions := "<@" + strings.Join(coaches, ">\n<@") + ">"
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Coaches",
			Value: mentions,
		})
	}

	if len(unassigned) > 0 {
		mentions := "<@" + strings.Join(unassigned, ">\n<@") + ">"
		mentions += "\nYou can assign these members with `/manage add`"
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Unassigned",
			Value: mentions,
		})
	}

	if len(invited) > 0 {
		mentions := "<@" + strings.Join(invited, ">\n<@") + ">"
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Invited",
			Value: mentions,
		})
	}

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
