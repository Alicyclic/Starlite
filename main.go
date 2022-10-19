package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	b, _ := discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN")))
	if err := b.Open(); err != nil {
		panic(err)
	}
	defer func() {
		b.Close()
	}()
	sc := make(chan os.Signal, 1)
	log.Println("Bot is now running. Press CTRL-C to exit.")
	signal.Notify(sc, os.Interrupt)
	<-sc
	log.Println("Bot is now closing...")
}
