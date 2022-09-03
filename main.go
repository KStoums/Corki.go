package main

import (
	"KBot/commands"
	"KBot/events"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	forever := make(chan bool, 1)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, os.Signal(syscall.SIGTERM))
	go func() {
		<-signalChan
		fmt.Println("> KBot.go down")
		forever <- true
	}()

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

	defer discord.Close()

	discord.AddHandler(events.MessageCreate)
	discord.AddHandler(events.MemberJoin)
	discord.AddHandler(events.MemberQuit)
	discord.AddHandler(events.ChannelCreate)

	commands.AddCommand(commands.LatencyCommand{}, commands.PingCommand{}, commands.ServerInfoCommand{},
		commands.ClearCommand{}, commands.KickCommand{}, commands.BanCommand{}, commands.NoteCommand{})

	discord.Identify.Intents = discordgo.IntentsAll
	discord.Open()
	fmt.Println("> KBot.go started !")
	<-forever
}
