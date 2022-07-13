package request

type MoreRequest struct {
	Name     string   `json:"name" form:"name"  validate:"required"`
	Games    []string `json:"games" form:"games"  validate:"required"`
	Password string   `json:"password" form:"name"  validate:"required"`
}
