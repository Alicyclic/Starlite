package modules

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/starlite/utils"
)

type CommandManager struct {
	*discordgo.Session
}

func NewCommandManager(s *discordgo.Session) (b *CommandManager) {
	b = &CommandManager{s}
	b.Init()
	return
}

func (b *CommandManager) RegisterCommandWithin(command *CommandDispatcher, guildId string) (err error) {
	command.GuildId = guildId
	err = b.RegisterCommand(command)
	return
}

func (ds *CommandManager) RegisterCommandsWithin(guildId string, commands []*CommandDispatcher) {
	for _, command := range commands {
		go func(command *CommandDispatcher) {
			err := ds.RegisterCommandWithin(command, guildId)
			if err != nil {
				panic(err)
			}
		}(command)
	}
}

func (ds *CommandManager) RegisterCommands(commands []*CommandDispatcher) {
	for _, command := range commands {
		go func(command *CommandDispatcher) {
			err := ds.RegisterCommand(command)
			if err != nil {
				panic(err)
			}
		}(command)
	}
}

func (b *CommandManager) RegisterCommand(app *CommandDispatcher) (err error) {
	if CommandExists(app.Name()) {
		log.Printf("Command %s already exists", app.Name())
		return
	}
	if !app.Global() {
		log.Printf("Command %s is guild-specific", app.Name())
	}
	app.Specification, err = b.ApplicationCommandCreate(b.State.User.ID, app.GuildId, app.Specification)
	if err != nil {
		panic(err)
	}
	AddCommand(app)
	return nil
}

func (ds *CommandManager) UnregisterCommands() {
	log.Println("Unregistering commands")
	for _, command := range Commands {
		go func(command *CommandDispatcher) {
			err := ds.ApplicationCommandDelete(ds.State.User.ID, command.GuildId, command.Specification.ID)
			if err != nil {
				panic(err)
			}
		}(command)
	}
}

func (b *CommandManager) Init() {
	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Interaction.Type {
		case discordgo.InteractionApplicationCommand:
			if command, ok := DispatchCommand(i.ApplicationCommandData().Name); ok {
				go command.Invoke(&Command{Session: s, Interaction: i.Interaction})
			}
		case discordgo.InteractionMessageComponent:
			if act, ok := utils.GetAction(i.MessageComponentData().CustomID); ok {
				act(s, i.Interaction)
			} else {
				fmt.Println("Action doesn't exist?")
			}
		}

	})
}
