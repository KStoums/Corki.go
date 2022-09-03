package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type LatencyCommand struct {
}

func (l LatencyCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var latencyValue = s.HeartbeatLatency()
	s.ChannelMessageDelete(m.ChannelID, m.ID)
	embed.SendEmbedAndDelete(messages.LatencyResponse.SetDescription(fmt.Sprintf("Latence: %d ms", latencyValue.Milliseconds())), s, m)
}

func (l LatencyCommand) Name() string {
	return "latence"
}
