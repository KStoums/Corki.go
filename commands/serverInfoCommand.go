package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type ServerInfoCommand struct {
}

func (s2 ServerInfoCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var server, err = s.Guild(m.GuildID)
	if err != nil {
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		return
	}

	embed.SendEmbedAndDelete(messages.ServerInfoResponse.SetDescription(fmt.Sprintf(`Serveur: %s
Membres: %d`, server.Name, len(members))), s, m)
}

func (s2 ServerInfoCommand) Name() string {
	return "serverInfo"
}
