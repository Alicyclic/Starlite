package modules

import (
	"fmt"

	"github.com/starlite/utils"
)

type PermissionUser struct {
	ID   string
	Rank int64
}

var users map[string]int64 = make(map[string]int64)

func NewPermissionUser(id string) *PermissionUser {
	users[id] = 0
	return &PermissionUser{ID: id, Rank: 0}
}

func (permission *PermissionUser) SetRank(rank int64) {
	users[permission.ID] = rank
}

func (permission *PermissionUser) Check(other string) bool {
	return users[permission.ID] >= users[other]
}

func (permission *PermissionUser) Print() *utils.Embed {
	return utils.NewEmbed().
		AddField("Permission", fmt.Sprintf("%v", users[permission.ID]), true)
}
