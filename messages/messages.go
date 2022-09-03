package messages

import (
	"KBot/utils/embed"
)

var iconUrl = "https://kstars.me/public/images/kstars.png"

//General
var CommandSyntaxe = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Syntaxe incorrect. Tapez !help pour plus d'informations.").SetTitle("Erreur:")

//Ping Command
var PongResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Pong!").SetTitle("Requête Ping:")

//ServerInfo Command
var ServerInfoResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetTitle("Informations:").SetImage(iconUrl)

//Latency Command
var LatencyResponse = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetTitle("Latence:")

//Clear Command
var ClearSuccess = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Messages supprimés !").SetTitle("Suppression:")
var NoIntDefineToClear = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous devez définir un nombre de message à supprimer.").SetTitle("Erreur:")

//Kick Command
var KickNoUserDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous n'avez pas notifié de membre.").SetTitle("Erreur:")
var MemberKicked = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Membre exclus !").SetTitle("Exclusion:")

//Note Command
var NoNoteDefine = embed.New().SetColor(embed.CYAN).SetDefaultFooter().SetDescription("Vous devez entrer une note.").SetTitle("Erreur:")
var NoteResponse = embed.New().SetColor(embed.RED_DARK).SetDefaultFooter().SetTitle("Note:")

//Welcome Message
var WelcomeMessage = embed.New().SetColor(embed.CYAN).SetDefaultFooter()
