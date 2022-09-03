package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func ChannelCreate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.BeforeUpdate == nil || v.BeforeUpdate.ChannelID != v.ChannelID {
		if v.ChannelID == "1015669797175971921" {
			channelToCreate, _ := s.GuildChannelCreateComplex(v.GuildID, discordgo.GuildChannelCreateData{
				Name:     fmt.Sprintf("🔈 Discussion"),
				Type:     discordgo.ChannelTypeGuildVoice,
				ParentID: "1010347750254202920",
			})

			s.GuildMemberMove(v.GuildID, v.UserID, &channelToCreate.ID)
			return
		}

		if v.BeforeUpdate == nil {
			return
		}

		var voiceState []*discordgo.VoiceState
		guild, _ := s.State.Guild(v.GuildID)
		for _, state := range guild.VoiceStates {
			if state.ChannelID == v.BeforeUpdate.ChannelID {
				voiceState = append(voiceState, state)
			}
		}

		fmt.Println(len(voiceState))
		if len(voiceState) == 0 {
			oldState := v.BeforeUpdate.ChannelID
			oldChannel, _ := s.Channel(oldState)

			if oldChannel.ParentID == "1010347750254202920" && oldChannel.ID != "1015669797175971921" {
				s.ChannelDelete(oldState)
			}
		}
	}
}
