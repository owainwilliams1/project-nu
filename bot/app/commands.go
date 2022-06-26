package app

import (
	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
)

// username		project-nu
//				valorant
//				apex
// team 		create
// 				info
//				accept
//				leave
//
// manage 		delete
// 				transfer-ownership
//				set-icon
//				invite
//				add					player
//									sub
//									coach
//				remove				player
//									sub
//									coach
//				kick
// register
// leaderboard
// help

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
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.ManageRouter,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "manage",
				Description: "Team management commands.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "delete",
						Description: "Delete your team.",
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "transfer-ownership",
						Description: "Transfer ownership of your team.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "user",
								Description: "The user you want to make owner.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "set-icon",
						Description: "Set your team's icon. Send a URL to the image.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "url",
								Description: "The URL of the image you would like to use.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "invite",
						Description: "Invite a member to your team.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "user",
								Description: "The user you want to invite to the team.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "kick",
						Description: "Kick a member from your team.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "user",
								Description: "The user you want to kick from the team.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
						Name:        "add",
						Description: "Add members to specific roles on your team.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "player",
								Description: "Make a member a player.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to make a player.",
										Required:    true,
									},
								},
							},
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "sub",
								Description: "Make a member a sub.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to make a sub.",
										Required:    true,
									},
								},
							},
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "coach",
								Description: "Make a member a coach.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to make a coach.",
										Required:    true,
									},
								},
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
						Name:        "remove",
						Description: "Remove members from specific roles on your team.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "player",
								Description: "Remove a member from player.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to remove from player.",
										Required:    true,
									},
								},
							},
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "sub",
								Description: "Remove a member from sub.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to remove from sub.",
										Required:    true,
									},
								},
							},
							{
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Name:        "coach",
								Description: "Remove a member from coach.",
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionUser,
										Name:        "user",
										Description: "The user you want to remove from coach.",
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: a.ManageRouter,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "username",
				Description: "Change your username.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "project-nu",
						Description: "Set your Project Nu username.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "username",
								Description: "The username you would like to have set.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "valorant",
						Description: "Set your Valorant username.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "username",
								Description: "The username you would like to have set.",
								Required:    true,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "apex-legends",
						Description: "Set your Apex username.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionUser,
								Name:        "username",
								Description: "The username you would like to have set.",
								Required:    true,
							},
						},
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
						Description: "The username you would like to use on Project Nu.",
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
				Description: "Get a leaderboard of teams.",
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

func (a *App) TeamRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	subcommand := options[0].Name
	args := options[0].Options
	switch subcommand {
	case "create":
		a.TeamCreate(s, i, args)
	case "info":
		a.TeamInfo(s, i, args)
	case "accept":
		a.TeamAccept(s, i, args)
	case "leave":
		a.TeamLeave(s, i, args)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) ManageRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	subcommand := options[0].Name
	args := options[0].Options
	switch subcommand {
	case "delete":
		a.ManageDelete(s, i, args)
	case "transfer-ownership":
		a.ManageTransferOwnership(s, i, args)
	case "set-icon":
		a.ManageSetIcon(s, i, args)
	case "invite":
		a.ManageInvite(s, i, args)
	case "kick":
		a.ManageKick(s, i, args)
	case "add":
		a.ManageAddRouter(s, i, args)
	case "remove":
		a.ManageRemoveRouter(s, i, args)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) UsernameRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	subcommand := options[0].Name
	args := options[0].Options
	switch subcommand {
	case "project-nu":
		a.ManageDelete(s, i, args)
	case "valorant":
		a.ManageTransferOwnership(s, i, args)
	case "apex-legends":
		a.ManageSetIcon(s, i, args)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) ManageAddRouter(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	subcommand := options[0].Name
	args := options[0].Options
	switch subcommand {
	case "player":
		a.ManageAddPlayer(s, i, args)
	case "sub":
		a.ManageAddSub(s, i, args)
	case "coach":
		a.ManageAddCoach(s, i, args)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) ManageRemoveRouter(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	options []*discordgo.ApplicationCommandInteractionDataOption,
) {
	subcommand := options[0].Name
	args := options[0].Options
	switch subcommand {
	case "player":
		a.ManageRemovePlayer(s, i, args)
	case "sub":
		a.ManageRemoveSub(s, i, args)
	case "coach":
		a.ManageRemoveCoach(s, i, args)
	default:
		a.RespondWithError(i, "That is not a subcommand.")
	}
}

func (a *App) PopulateSCM() {
	features := a.GetFeatures()
	a.Manager.AddFeatures(features)
}
