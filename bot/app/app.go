package app

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
	api "hushclan.com/api/database"
	"hushclan.com/types"
)

type App struct {
	Session  *discordgo.Session
	Manager  *scm.SCM
	Database *api.Database
	Envs     Vars
}

type Vars struct {
	Token string
	Guild string
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
		embed.Image = &discordgo.MessageEmbedImage{
			URL: team.Icon,
		}
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: team.TeamID,
	}

	embed.Color = team.Color

	return
}
