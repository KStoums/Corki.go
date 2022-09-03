package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct {
}

func (p PingCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageDelete(m.ChannelID, m.ID)
	embed.SendEmbedAndDelete(messages.PongResponse, s, m)
}

func (p PingCommand) Name() string {
	return "ping"
}
