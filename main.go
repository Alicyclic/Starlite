package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/starlite/commands"
	"github.com/starlite/modules"
	"github.com/starlite/utils"
)

var commandManager *modules.CommandManager

func main() {
	commands.Hi()
	b, _ := discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN")))
	if err := b.Open(); err != nil {
		panic(err)
	}
	commandManager = modules.NewCommandManager(b)
	commandManager.RegisterCommandsWithin("1026149960489644224", modules.GetCommands())

	b.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !utils.IsDiscordInvite(m.Content) {
			return
		}
		if e := s.ChannelMessageDelete(m.ChannelID, m.Reference().MessageID); e != nil {
			panic(e.Error())
		}
	})

	defer func() {
		// commandManager.UnregisterCommands()
		b.Close()
	}()
	sc := make(chan os.Signal, 1)
	log.Println("Bot is now running. Press CTRL-C to exit.")
	signal.Notify(sc, os.Interrupt)
	<-sc
	log.Println("Bot is now closing...")
}
