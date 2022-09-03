package main

import (
	"KBot/commands"
	"KBot/events"
	"fmt"
	"github.com/bwmarrin/discordgo"
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

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	defer discord.Close()

	discord.AddHandler(events.MessageCreate)
	discord.AddHandler(events.MemberJoin)
	discord.AddHandler(events.MemberQuit)
	discord.AddHandler(events.ChannelCreate)
	discord.AddHandler(events.RawEvent)
	discord.AddHandler(events.Ready)

	commands.AddCommand(commands.LatencyCommand{}, commands.PingCommand{}, commands.ServerInfoCommand{},
		commands.ClearCommand{}, commands.KickCommand{}, commands.BanCommand{}, commands.NoteCommand{})

	discord.Identify.Intents = discordgo.IntentsAll
	err = discord.Open()
	fmt.Println("> KBot.go started !")
	<-forever
}
