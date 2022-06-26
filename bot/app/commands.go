package app

import (
	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
)

// register
// member		username	user
//							valorant
//							apex
// team 		create
// 				info
//				accept
//				leave
//
// manage 		delete
// 				transfer
//				icon
//				invite
//				add			player
//							sub
//							coach
//				remove		player
//							sub
//							coach
//				kick
// leaderboard
// help

func (a *App) GetFeaturesOld() []*scm.Feature {
	return []*scm.Feature{
		// {
		// 	Type:    discordgo.InteractionApplicationCommand,
		// 	Handler: a.CreateTeam,
		// 	ApplicationCommand: &discordgo.ApplicationCommand{
		// 		Name:        "create-team",
		// 		Description: "Create a team. You must have a team manager role to do this.",
		// 		Options: []*discordgo.ApplicationCommandOption{
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "name",
		// 				Description: "The name of your team.",
		// 				Required:    true,
		// 			},
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "color",
		// 				Description: "A hex color code for the team.",
		// 				Required:    true,
		// 			},
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "region",
		// 				Description: "The team's region: EU, NA.",
		// 				Required:    true,
		// 			},
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "sex",
		// 				Description: "The team's sex: Male, Female, Mixed.",
		// 				Required:    true,
		// 			},
		// 		},
		// 	},
		// },
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
		// {
		// 	Type:    discordgo.InteractionApplicationCommand,
		// 	Handler: a.AcceptInvite,
		// 	ApplicationCommand: &discordgo.ApplicationCommand{
		// 		Name:        "accept-invite",
		// 		Description: "Accept an invite to a team. You must not be in a team to do this.",
		// 		Options: []*discordgo.ApplicationCommandOption{
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "team-id",
		// 				Description: "The ID of the team you wish to join.",
		// 				Required:    true,
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	Type:    discordgo.InteractionApplicationCommand,
		// 	Handler: a.Leave,
		// 	ApplicationCommand: &discordgo.ApplicationCommand{
		// 		Name:        "leave",
		// 		Description: "Leave your team. You must be in a team to do this.",
		// 	},
		// },
		// {
		// 	Type:    discordgo.InteractionApplicationCommand,
		// 	Handler: a.Team,
		// 	ApplicationCommand: &discordgo.ApplicationCommand{
		// 		Name:        "team",
		// 		Description: "Get your team's info or another team's.",
		// 		Options: []*discordgo.ApplicationCommandOption{
		// 			{
		// 				Type:        discordgo.ApplicationCommandOptionString,
		// 				Name:        "team-name",
		// 				Description: "The name of the team you would like to view.",
		// 				Required:    false,
		// 			},
		// 		},
		// 	},
		// },
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
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.TransferOwnership,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "transfer-ownership",
				Description: "Transfer ownership of your team.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "The user you want to make .",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.Help,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "help",
				Description: "Get some help.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "page",
						Description: "The help page you would like to view.",
						Required:    false,
					},
				},
			},
		},
	}
}

func (a *App) GetFeatures() []*scm.Feature {
	return []*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.TeamRouter,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "team",
				Description: "Team commands.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "create",
						Description: "Create a team",
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
								Choices: []*discordgo.ApplicationCommandOptionChoice{
									{
										Name:  "EU",
										Value: "eu",
									},
									{
										Name:  "NA",
										Value: "na",
									},
								},
								Required: true,
							},
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "game",
								Description: "The game your team plays.",
								Choices: []*discordgo.ApplicationCommandOptionChoice{
									{
										Name:  "Valorant",
										Value: "valorant",
									},
									{
										Name:  "Apex Legends",
										Value: "apex",
									},
								},
								Required: true,
							},
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "sex",
								Description: "The team's sex: Male, Female, Mixed.",
								Choices: []*discordgo.ApplicationCommandOptionChoice{
									{
										Name:  "Male",
										Value: "male",
									},
									{
										Name:  "Female",
										Value: "female",
									},
									{
										Name:  "Mixed",
										Value: "mixed",
									},
								},
								Required: true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "info",
						Description: "Get your team's info or another team's.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "team-name",
								Description: "The name of the team you wish to view.",
								Required:    false,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "accept",
						Description: "Accept an invite to a team. You must not be in a team to do this.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "team-id",
								Description: "The ID of the team you wish to join.",
								Required:    false,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "leave",
						Description: "Leave your team.",
					},
				},
			},
		},
	}
}

func (a *App) TeamRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	subcommand := options[0].Name
	o := options[1:]
	switch subcommand {
	case "create":
		a.TeamCreate(s, i, o)
	case "info":
		a.TeamInfo(s, i, o)
	case "accept":
		a.TeamAccept(s, i, o)
	case "leave":
		a.TeamLeave(s, i, o)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) PopulateSCM() {
	features := a.GetFeatures()
	a.Manager.AddFeatures(features)
}
