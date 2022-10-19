package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/starlite/modules"
	"github.com/starlite/utils"
)

// TODO: Add your code here!
type mojangSession struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// changed
func GetMojangSession(username string) (session mojangSession, err error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&session)
	return
}

func mcModel(id, usr string) (e *utils.Embed) {
	e = utils.NewEmbed()
	head := fmt.Sprintf("https://www.mc-heads.net/head/%s.png", usr)
	namemc := fmt.Sprintf("https://www.namemc.com/%s", usr)
	e.SetAuthor(usr, namemc, head)
	return
}

func init() {
	modules.NewCommand("minecraft", "Get the mojang session for a username").
		AddOption("username", "The username to get the session for", discordgo.ApplicationCommandOptionString, true).
		SetHandler(func(c *modules.Command) {
			arg, _ := c.GetOption("username")
			username := arg.StringValue()
			username = strings.Replace(username, " ", "_", -1)
			session, err := GetMojangSession(username)
			if err != nil {
				panic(err)
			}
			c.Send(mcModel(session.ID, session.Name))
		})
}
