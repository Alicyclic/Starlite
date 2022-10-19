package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

func NewEmbed() *Embed {
	return &Embed{
		MessageEmbed: &discordgo.MessageEmbed{},
	}
}

func (e *Embed) ConvertFieldsToDescription() *Embed {
	if e.Description == "" || len(e.Fields) > 0 {
		for _, field := range e.Fields {
			if e.Description != "" {
				e.Description += "\n"
			}
			e.MessageEmbed.Description += fmt.Sprintf("**%s**: %s", field.Name, field.Value)
		}
	}
	return e
}

func (e *Embed) SetAuthor(Name, URL, IconURL string) *Embed {
	e.Author = &discordgo.MessageEmbedAuthor{
		Name:    Name,
		URL:     URL,
		IconURL: IconURL,
	}
	e.MessageEmbed.Author = e.Author
	return e
}

func (e *Embed) SetTitle(title string) *Embed {
	e.MessageEmbed.Title = e.TextLimiter(title, 256)
	return e
}

func (e *Embed) TextLimiter(str string, limiter int) string {
	if len(str) > limiter {
		str = str[:limiter]
	}
	return str
}

func (e *Embed) SetDescription(msg string) *Embed {
	e.MessageEmbed.Description = e.TextLimiter(msg, 2048)
	return e
}

func (e *Embed) SetColor(color string) *Embed {
	color = strings.Replace(color, "0x", "", -1)
	color = strings.Replace(color, "0X", "", -1)
	color = strings.Replace(color, "#", "", -1)
	colorInt, err := strconv.ParseInt(color, 16, 64)
	if err != nil {
		panic(err)
	}
	e.MessageEmbed.Color = int(colorInt)
	return e
}

func (e *Embed) SetURL(URL string) *Embed {
	e.MessageEmbed.URL = URL
	return e
}

func (e *Embed) SetThumbnail(URL string) *Embed {
	e.MessageEmbed.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: URL}
	return e
}

func (e *Embed) SetImage(URL string) *Embed {
	e.MessageEmbed.Image = &discordgo.MessageEmbedImage{URL: URL}
	return e
}

func (e *Embed) SetFooter(Text, IconURL string) *Embed {
	Text = e.TextLimiter(Text, 2048)
	e.MessageEmbed.Footer = &discordgo.MessageEmbedFooter{Text: Text, IconURL: IconURL}
	return e
}

func (e *Embed) AddField(Name, Value string, Inline bool) *Embed {
	Name = e.TextLimiter(Name, 256)
	Value = e.TextLimiter(Value, 1024)
	e.MessageEmbed.Fields = append(e.MessageEmbed.Fields, &discordgo.MessageEmbedField{Name: Name, Value: Value, Inline: Inline})
	return e
}

func (e *Embed) CreateMessageEmbed() *discordgo.MessageEmbed {
	return e.MessageEmbed
}
