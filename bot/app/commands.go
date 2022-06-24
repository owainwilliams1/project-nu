package app

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
)

func (a *App) GetFeatures() []*scm.Feature {
	return []*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.CreateTeam,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "create-team",
				Description: "Create a team. You must have a team manager role to do this.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "The name of your team.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "color",
						Description: "A hex color code for the team.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "region",
						Description: "The team's region: EU, NA.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "sex",
						Description: "The team's sex: Male, Female, Mixed.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.DeleteTeam,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "delete-team",
				Description: "Delete a team. You must own a team to do this.",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.InviteMember,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "invite",
				Description: "Invite a member to your team. You must own a team to do this.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you want to invite to your team.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.RemoveMember,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "remove",
				Description: "Remove a member from your team. You must own a team to do this.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you want to remove from your team.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.AcceptInvite,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "accept-invite",
				Description: "Accept an invite to a team. You must not be in a team to do this.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "team-id",
						Description: "The ID of the team you wish to join.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.Leave,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "leave",
				Description: "Leave your team. You must be in a team to do this.",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.Team,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "team",
				Description: "Get your team's info. You must be in a team to do this.",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.TeamInfo,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "team-info",
				Description: "Get info on a team.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "team-id",
						Description: "The ID of the team you wish to look up.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.Leaderboard,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "leaderboard",
				Description: "Get the team leaderboard.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "page",
						Description: "Leaderboard page.",
						Required:    false,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.Register,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "register",
				Description: "Register yourself.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "username",
						Description: "Valorant Username.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.MakeMemberCoach,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "make-coach",
				Description: "Make a team member a coach.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to make a coach.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.MakeMemberPlayer,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "make-player",
				Description: "Make a team member a player.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to make a player.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.MakeMemberSub,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "make-sub",
				Description: "Make a team member a sub.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to make a sub.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.RemoveMemberCoach,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "remove-coach",
				Description: "Remove a member from being a coach.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to remove.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.RemoveMemberPlayer,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "remove-player",
				Description: "Remove a member from being a player.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to remove.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.RemoveMemberSub,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "remove-sub",
				Description: "Remove a member from being a sub.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you would like to remove.",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.SetTeamIcon,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "set-team-icon",
				Description: "Set your team's icon.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "url",
						Description: "The URL of the image you would like to use.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (a *App) PopulateSCM() {
	features := a.GetFeatures()
	a.Manager.AddFeatures(features)
}

func (a *App) RegisterCommands(guildID string) {
	a.Session.AddHandler(a.Manager.HandleInteraction)

	err := a.Manager.CreateCommands(a.Session, guildID)
	if err != nil {
		log.Fatal("could not create commands", err)
	}
}

func (a *App) DeleteCommands() {
	for _, guildID := range a.JoinedGuilds {
		err := a.Manager.DeleteCommands(a.Session, guildID)
		if err != nil {
			m := fmt.Sprintf("could not delete commands for server %s", guildID)
			a.Log.Error(m, err)
		}
	}
}
