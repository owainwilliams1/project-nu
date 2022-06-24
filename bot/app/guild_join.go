package app

import "github.com/bwmarrin/discordgo"

func (a *App) HandleGuildJoin(s *discordgo.Session, c *discordgo.GuildCreate) {
	a.RegisterCommands(c.Guild.ID)
}
