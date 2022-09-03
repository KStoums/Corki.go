package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Run(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	Name() string
}

var Commands []Command

func AddCommand(command ...Command) {
	Commands = append(Commands, command...)
}
