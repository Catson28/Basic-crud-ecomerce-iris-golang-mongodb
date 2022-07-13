package controllers

import (
	"strings"

	iris "github.com/kataras/iris/v12"
	permissions "github.com/xyproto/permissionbolt"
)

type UserController struct {
	Ctx         iris.Context
	Permissions permissions.UserState
}

func (c *UserController) Init() {
}

func (c *UserController) GetAdmin() {
	c.Ctx.WriteString("super secret information that only logged in administrators must see!\n\n")
	if usernames, err := c.Permissions.AllUsernames(); err == nil {
		c.Ctx.Writef("lista de todos usuarios: %s" + strings.Join(usernames, ", "))
	}
}
