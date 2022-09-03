package main

import (
	"KBot/utils/embed"
	"KBot/utils/emoji"
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
		fmt.Println("ping", args)
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		sendEmbedAndDelete(pongResponse, s, m)
		return
	}

	if command == "!informations" || command == "!information" || command == "!info" {
		fmt.Println("informations", args)
		var server, err = s.Guild(m.GuildID)
		if err != nil {
			return
		}

		s.ChannelMessageDelete(m.ChannelID, m.ID)
		members, err := s.GuildMembers(m.GuildID, "", 1000)
		if err != nil {
			return
		}

		sendEmbedAndDelete(serverInfoResponse.SetDescription(fmt.Sprintf(`Serveur: %s
Il y a %d membres`, server.Name, len(members))), s, m)
		return
	}

	if command == "!latence" {
		fmt.Println("latence", args)
		var latencyValue = s.HeartbeatLatency()
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		sendEmbedAndDelete(latencyResponse.SetDescription(fmt.Sprintf("Latence: %d ms", latencyValue.Milliseconds())), s, m)
		return
	}

	if command == "!clear" {
		fmt.Println("clear", args)
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		if len(args) == 0 {
			sendEmbedAndDelete(noIntDefineToClear, s, m)
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
		sendEmbedAndDelete(clearSuccess, s, m)
		return
	}

	if command == "!kick" {
		fmt.Println("kick", args)
		if len(args) == 0 {
			s.ChannelMessageSendEmbed(m.ChannelID, kickNoUserDefine.ToMessageEmbed())
			return
		}

		if len(args) > 1 {
			sendEmbedAndDelete(commandSyntaxe, s, m)
			return
		}

		if len(m.Mentions) == 0 || len(m.Mentions) > 1 {
			sendEmbedAndDelete(commandSyntaxe, s, m)
			return
		}

		var memberToKick = m.Mentions[0]
		s.GuildMemberDelete(m.GuildID, memberToKick.ID)
		sendEmbedAndDelete(memberKicked, s, m)
		return
	}

	if command == "!note" {
		fmt.Println("note", args)
		if len(args) == 0 {
			sendEmbedAndDelete(noNoteDefine, s, m)
			return
		}

		channelId, err := s.Channel("1014890741853589537")
		if err != nil {
			return
		}

		noteEmbed, err := s.ChannelMessageSendEmbed(channelId.ID, noteResponse.SetDescription(strings.Join(args, " ")).SetTitle("Note:").SetColor(embed.RED_DARK).ToMessageEmbed())
		if err != nil {
			return
		}

		err = s.MessageReactionAdd(m.ChannelID, noteEmbed.ID, emoji.ConvertEmoji("📤"))
	}

}

func sendEmbedAndDelete(embed *embed.Embed, s *discordgo.Session, m *discordgo.MessageCreate) {
	sendEmbed, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.ToMessageEmbed())
	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)
	s.ChannelMessageDelete(m.ChannelID, sendEmbed.ID)
}
