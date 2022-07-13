package request

type SignRequest struct {
	Username string `json:"username" form:"username"  validate:"required"`
	Password string `json:"password" form:"username"  validate:"required"`
}
