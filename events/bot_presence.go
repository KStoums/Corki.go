package events

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	statusList := []string{
		"https://kstars.me",
		"apprendre GoLang",
	}

	current := 0

	for {
		if current == 0 {
			s.UpdateGameStatus(0, statusList[0])
			current += 1
		}

		if current == 1 {
			s.UpdateGameStatus(0, statusList[1])
			current -= 1
		}

		time.Sleep(5 * time.Second)
	}

}
