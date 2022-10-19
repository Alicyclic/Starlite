package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/starlite/utils"
)

type Command struct {
	*discordgo.Session
	*discordgo.Interaction
	Options map[string]*discordgo.ApplicationCommandInteractionDataOption
}

func (c *Command) GetOption(option string) (value *discordgo.ApplicationCommandInteractionDataOption, ok bool) {
	value, ok = c.Options[option]
	return
}

func (c *Command) ParseOptions() {
	options := c.Interaction.ApplicationCommandData().Options
	c.Options = make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		c.Options[opt.Name] = opt
	}
}

func (c *Command) AddOption(name string, t discordgo.ApplicationCommandOptionType) *Command {
	c.Options[name] = &discordgo.ApplicationCommandInteractionDataOption{Name: name, Type: t}
	return c
}

func (c *Command) CheckIfNSFW() bool {
	channel, err := c.State.Channel(c.ChannelID)
	return (err == nil) && channel.NSFW
}

func (c *Command) Bot() *discordgo.Member {
	member, err := c.Session.GuildMember(c.GuildID, c.State.User.ID)
	if err != nil {
		panic(err)
	}
	return member
}

func (c *Command) Send(embed ...*utils.Embed) {
	NewResponse(c.Interaction, c.Session).
		AddEmbeds(embed...).
		SendResponse()
}

func (c *Command) SendNSFWMessage(embed ...*utils.Embed) {
	NewResponse(c.Interaction, c.Session).
		SetEphemeral(!c.CheckIfNSFW()).
		AddEmbeds(embed...).
		SendResponse()
}
