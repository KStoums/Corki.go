package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"github.com/bwmarrin/discordgo"
	"github.com/thoas/go-funk"
)

type BanCommand struct {
}

func (b BanCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !funk.ContainsString(m.Member.Roles, "1010346312459358269") {
		return
	}

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
	s.GuildBan(m.GuildID, memberToKick.ID)
	embed.SendEmbedAndDelete(messages.MemberKicked, s, m)
}

func (b BanCommand) Name() string {
	return "ban"
}
