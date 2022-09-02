package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
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

	if m.Content == "!ping" {
		s.ChannelMessageSendEmbed(m.ChannelID, pongResponse)
	}

	if m.Content == "!serverInfo" || m.Content == "!info" {
		var server, err = s.Guild(m.GuildID)
		if err != nil {
			return
		}
		members, err := s.GuildMembers(m.GuildID, "", 1000)
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, serverInfoResponse.SetDescription(fmt.Sprintf(`Serveur: %s
Il y a %d membres`, server.Name, len(members))).ToMessageEmbed())
		if err != nil {
			fmt.Println(err)
		}
	}

}
