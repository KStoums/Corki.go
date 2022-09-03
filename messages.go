package main

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"

//General
var commandSyntaxe = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Syntaxe incorrect. Tapez !help pour plus d'informations.").SetTitle("Erreur:")

//Ping Command
var pongResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Pong!").SetTitle("Requête Ping:")

//ServerInfo Command
var serverInfoResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetTitle("Informations:").SetImage(iconUrl)

//Latency Command
var latencyResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetTitle("Latence:")

//Clear Command
var clearSuccess = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Messages supprimés !").SetTitle("Suppression:")
var noIntDefineToClear = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous devez définir un nombre de message à supprimer.").SetTitle("Erreur:")

//Kick Command
var kickNoUserDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous n'avez pas notifié de membre.").SetTitle("Erreur:")
var memberKicked = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Membre exclus !").SetTitle("Exclusion:")

//Note Command
var noNoteDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous devez entrer une note.").SetTitle("Erreur:")
var noteResponse = embed.New().SetColor(embed.RED_DARK).SetDefaultFooter().SetTitle("Note:")
