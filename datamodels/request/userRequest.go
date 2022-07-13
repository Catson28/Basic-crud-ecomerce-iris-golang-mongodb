package request

import "tentativa/datamodels"

type UserRequest struct {
	Firstname string            `json:"firstname" form:"firstname"  validate:"required"`
	Username  string            `json:"username" form:"username"  validate:"required"`
	Password  string            `json:"password" form:"password"  validate:"required"`
	Roles     []datamodels.Role `json:"roles" form:"roles"  validate:"required"`
}
