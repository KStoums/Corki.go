package main

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"
var baseEmbed = embed.New().SetColor(embed.BLUE_LIGHT).SetDefaultFooter().SetAuthor("KBot.go", "", "")

//Ping Command
var pongResponse = baseEmbed.SetDescription("Pong!")

//ServerInfo Command
var serverInfoResponse = baseEmbed

//Latency Command
var latencyResponse = baseEmbed

//Clear Command
var clearSuccess = baseEmbed.SetDescription("Messages supprimés !")
