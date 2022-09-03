package embed

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func SendEmbedAndDelete(embed *Embed, s *discordgo.Session, m *discordgo.MessageCreate) {
	sendEmbed, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.ToMessageEmbed())
	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)
	s.ChannelMessageDelete(m.ChannelID, sendEmbed.ID)
}
