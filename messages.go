package main

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"

//Ping Command
var pongResponse = embed.New().SetDescription("Pong!").SetAuthor("KBot.go", iconUrl, "").SetColor(embed.BLUE_LIGHT).SetDefaultFooter().SetImage(iconUrl).ToMessageEmbed()

var serverInfoResponse = embed.New().SetColor(embed.BLUE_LIGHT).SetDefaultFooter().SetAuthor("KBot.go", iconUrl, "").SetImage(iconUrl)
