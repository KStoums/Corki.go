package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"github.com/bwmarrin/discordgo"
	"github.com/thoas/go-funk"
	"strings"
)

type NoteCommand struct {
}

func (n NoteCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !funk.ContainsString(m.Member.Roles, "1010346312459358269") {
		return
	}

	if len(args) == 0 {
		embed.SendEmbedAndDelete(messages.NoNoteDefine, s, m)
		return
	}

	channelId, err := s.Channel("1014890741853589537")
	if err != nil {
		return
	}

	noteEmbed, err := s.ChannelMessageSendEmbed(channelId.ID, messages.NoteResponse.SetDescription(strings.Join(args, " ")).ToMessageEmbed())
	if err != nil {
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.MessageReactionAdd(channelId.ID, noteEmbed.ID, "📤")
}

func (n NoteCommand) Name() string {
	return "note"
}
