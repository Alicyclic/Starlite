package commands

import (
	"encoding/json"

	"github.com/starlite/modules"
	"github.com/starlite/utils"
)

func init() {
	modules.NewCommand("embed", "Send an embed to a channel").
		AddOption("json", "parse json string!", 3, true).
		SetHandler(func(c *modules.Command) {
			arg, _ := c.GetOption("json")
			var embed utils.Embed
			if err := json.Unmarshal([]byte(arg.StringValue()), &embed); err != nil {
				panic(err)
			}
			c.Send(&embed)
		})
}
