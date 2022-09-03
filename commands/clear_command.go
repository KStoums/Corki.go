package commands

import (
	"KBot/messages"
	"KBot/utils/embed"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type ClearCommand struct {
}

func (c ClearCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageDelete(m.ChannelID, m.ID)
	if len(args) == 0 {
		embed.SendEmbedAndDelete(messages.NoIntDefineToClear, s, m)
		return
	}

	messageToDelete, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	if messageToDelete > 100 {
		return
	}

	roleFound := false
	for _, role := range m.Member.Roles {
		if role == "1010346312459358269" {
			roleFound = true
		}
	}

	if !roleFound {
		return
	}

	msg, err := s.ChannelMessages(m.ChannelID, messageToDelete, "", "", "")
	if err != nil {
		return
	}

	var messagesId []string

	for _, s2 := range msg {
		messagesId = append(messagesId, s2.ID)
	}

	s.ChannelMessagesBulkDelete(m.ChannelID, messagesId)
	embed.SendEmbedAndDelete(messages.ClearSuccess, s, m)
}

func (c ClearCommand) Name() string {
	return "clear"
}
