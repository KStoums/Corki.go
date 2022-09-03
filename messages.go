package main

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"

//General
var commandSyntaxe = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Erreur: Syntaxe incorrect. Tapez !help pour plus d'informations.")

//Ping Command
var pongResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Pong!")

//ServerInfo Command
var serverInfoResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter()

//Latency Command
var latencyResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter()

//Clear Command
var clearSuccess = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Messages supprimés !")
var noIntDefineToClear = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Erreur: Vous devez définir un nombre de message à supprimer.")

//Kick Command
var kickNoUserDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Erreur: Vous n'avez pas notifié de membre.")
var memberKicked = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Membre exclus !")

//Note Command
var noNoteDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Erreur: Vous devez entrer une note.")
var noteResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter()
