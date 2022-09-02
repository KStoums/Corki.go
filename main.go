package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

	args := strings.Split(m.Content, " ")
	command := args[0]
	args = args[1:]

	if command == "!ping" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		embedResponse, err := s.ChannelMessageSendEmbed(m.ChannelID, pongResponse.ToMessageEmbed())
		if err != nil {
			return
		}

		time.Sleep(5 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, embedResponse.ID)
	}

	if command == "!informations" || command == "!information" || command == "!info" {
		var server, err = s.Guild(m.GuildID)
		if err != nil {
			return
		}
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		members, err := s.GuildMembers(m.GuildID, "", 1000)
		embedResponse, err := s.ChannelMessageSendEmbed(m.ChannelID, serverInfoResponse.SetDescription(fmt.Sprintf(`Serveur: %s
Il y a %d membres`, server.Name, len(members))).ToMessageEmbed())
		_, err = embedResponse, err
		if err != nil {
			return
		}

		time.Sleep(5 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, embedResponse.ID)
	}

	if command == "!latence" {
		var latencyValue = s.HeartbeatLatency()
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		//Round latencyValue
		embedResponse, err := s.ChannelMessageSendEmbed(m.ChannelID, latencyResponse.SetDescription(fmt.Sprintf("Latence: %d ms", latencyValue.Milliseconds())).ToMessageEmbed())
		if err != nil {
			return
		}

		time.Sleep(5 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, embedResponse.ID)
	}

	if command == "!clear" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		if len(args) == 0 {
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

		messages, err := s.ChannelMessages(m.ChannelID, messageToDelete, "", "", "")
		if err != nil {
			return
		}

		var messagesId []string

		for _, s2 := range messages {
			messagesId = append(messagesId, s2.ID)
		}

		s.ChannelMessagesBulkDelete(m.ChannelID, messagesId)
		embedClearSuccess, err := s.ChannelMessageSendEmbed(m.ChannelID, clearSuccess.ToMessageEmbed())
		if err != nil {
			return
		}

		time.Sleep(5 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, embedClearSuccess.ID)
	}

}
