package events

import "github.com/bwmarrin/discordgo"

func RawEvent(s *discordgo.Session, e *discordgo.Event) {
	reactionAdd, ok := e.Struct.(*discordgo.MessageReactionAdd)
	if !ok {
		return
	}

	if reactionAdd.Emoji.Name == "📤" && reactionAdd.ChannelID == "1014890741853589537" && s.State.User.ID != reactionAdd.UserID {
		s.ChannelMessageDelete(reactionAdd.ChannelID, reactionAdd.MessageID)
	}
}
