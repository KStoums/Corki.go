package main

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"
var baseEmbed = embed.New().SetColor(embed.BLUE_LIGHT).SetDefaultFooter().SetAuthor("KBot.go", "", "")

//General
var commandSyntaxe = baseEmbed.SetDescription("Erreur: Syntaxe incorrect. Tapez !help pour plus d'informations.")

//Ping Command
var pongResponse = baseEmbed.SetDescription("Pong!")

//ServerInfo Command
var serverInfoResponse = baseEmbed

//Latency Command
var latencyResponse = baseEmbed

//Clear Command
var clearSuccess = baseEmbed.SetDescription("Messages supprimés !")

//Kick Command
var kickNoUserDefine = baseEmbed.SetDescription("Erreur: Vous n'avez pas notifié de membre.")
var memberKicked = baseEmbed.SetDescription("Utilisateur exclus !")
