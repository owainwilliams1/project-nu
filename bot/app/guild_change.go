package app

import "github.com/bwmarrin/discordgo"

func (a *App) HandleGuildJoin(s *discordgo.Session, c *discordgo.GuildCreate) {
	a.RegisterCommands(c.Guild.ID)
}

func (a *App) HandleGuildLeave(s *discordgo.Session, c *discordgo.GuildDelete) {
	a.DeleteCommands(c.Guild.ID)
}
