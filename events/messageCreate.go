package events

import (
	"KBot/commands"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")
	command := args[0]
	args = args[1:]

	for _, c := range commands.Commands {
		if "!"+c.Name() == command {
			c.Run(s, m, args)
			break
		}
	}
}
