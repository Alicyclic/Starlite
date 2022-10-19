package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var actions map[string]func(*discordgo.Session, *discordgo.Interaction) = make(map[string]func(*discordgo.Session, *discordgo.Interaction))

func AddAction(id string, action func(*discordgo.Session, *discordgo.Interaction)) {
	actions[id] = action
}

func GetAction(id string) (act func(*discordgo.Session, *discordgo.Interaction), ok bool) {
	act, ok = actions[id]
	log.Printf("Get MessageCom -> `%s` exists %v", id, ok)
	return
}
