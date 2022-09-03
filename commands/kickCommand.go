package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"github.com/bwmarrin/discordgo"
)

type KickCommand struct {
}

func (k KickCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		s.ChannelMessageSendEmbed(m.ChannelID, messages.KickNoUserDefine.ToMessageEmbed())
		return
	}

	if len(args) > 1 {
		embed.SendEmbedAndDelete(messages.CommandSyntaxe, s, m)
		return
	}

	if len(m.Mentions) == 0 || len(m.Mentions) > 1 {
		embed.SendEmbedAndDelete(messages.CommandSyntaxe, s, m)
		return
	}

	var memberToKick = m.Mentions[0]
	s.GuildMemberDelete(m.GuildID, memberToKick.ID)
	embed.SendEmbedAndDelete(messages.MemberKicked, s, m)
}

func (k KickCommand) Name() string {
	return "kick"
}
