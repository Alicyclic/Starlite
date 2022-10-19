package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/starlite/utils"
)

type Response struct {
	Ephemeral bool
	*discordgo.Interaction
	*discordgo.Session
	*discordgo.InteractionResponseData
}

func NewResponse(i *discordgo.Interaction, s *discordgo.Session) *Response {
	return &Response{
		Session:                 s,
		Interaction:             i,
		Ephemeral:               false,
		InteractionResponseData: &discordgo.InteractionResponseData{},
	}
}

func (r *Response) AddMessageComponent(com ...discordgo.MessageComponent) {
	r.Components = append(r.Components, discordgo.ActionsRow{Components: com})
	fmt.Printf("r.Components: %v\n", r.Components)
}

func (r *Response) CheckPermission(userID string, channelID string, permission int64) bool {
	p, err := r.State.UserChannelPermissions(userID, channelID)
	return (err == nil) && (p&permission == permission)
}

func (r *Response) AddEmbeds(e ...*utils.Embed) *Response {
	if !r.CheckPermission(r.State.User.ID, r.ChannelID, discordgo.PermissionEmbedLinks) {
		r.Content = "Unable to send embed messages!"
		return r
	}
	if len(e) >= 4000 {
		r.Embeds = r.Embeds[:4000]
	}
	for _, v := range e {
		r.Embeds = append(r.Embeds, v.MessageEmbed)
	}
	return r
}

func (r *Response) SetEphemeral(eph bool) *Response {
	r.Ephemeral = eph
	return r
}

func (r *Response) ResponseId() string {
	id, _ := r.InteractionResponse(r.Interaction)
	return id.Reference().MessageID
}

func (r *Response) SendResponse() {
	if r.Ephemeral {
		r.Flags = discordgo.MessageFlagsEphemeral
	}
	r.InteractionRespond(r.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: r.InteractionResponseData,
	})
}
