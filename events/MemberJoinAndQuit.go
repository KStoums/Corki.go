package events

import (
	"KBot/messages"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "1014833151115268096")
	welcomeChannel, err := s.Channel("1014831045616599091")
	if err != nil {
		return
	}

	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		return
	}

	s.ChannelMessageSendEmbed(welcomeChannel.ID, messages.WelcomeMessage.
		SetDescription(fmt.Sprintf("Bienvenue à %s ! \nNous sommes désormais %d", m.User.Username, len(members))).
		SetTitle("Nouveau membre:").
		ToMessageEmbed())
}

func MemberQuit(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	welcomeChannel, err := s.Channel("1014831045616599091")
	if err != nil {
		return
	}

	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		return
	}

	s.ChannelMessageSendEmbed(welcomeChannel.ID, messages.WelcomeMessage.
		SetDescription(fmt.Sprintf("%s nous a quitté ! \nNous sommes désormais %d", m.User.Username, len(members))).
		SetTitle("A quitter le serveur:").
		ToMessageEmbed())
}
