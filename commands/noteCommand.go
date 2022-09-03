package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"KBot/utils/emoji"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type NoteCommand struct {
}

func (n NoteCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
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

	err = s.MessageReactionAdd(m.ChannelID, noteEmbed.ID, emoji.ConvertEmoji("📤"))
}

func (n NoteCommand) Name() string {
	return "note"
}
