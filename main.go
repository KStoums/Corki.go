package main

import (
	"KBot/commands"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open(botToken)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + string(bytes))
	if err != nil {
		panic(err)
	}

	discord.AddHandler(messageCreate)
	commands.AddCommand(commands.LatencyCommand{}, commands.PingCommand{}, commands.ServerInfoCommand{},
		commands.ClearCommand{}, commands.KickCommand{}, commands.BanCommand{}, commands.NoteCommand{})

	discord.Identify.Intents = discordgo.IntentsAll
	group := sync.WaitGroup{}
	group.Add(1)
	discord.Open()
	group.Wait()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
