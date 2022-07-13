package request

// TodoRequest represents a Todo HTTP request.
type TodoRequest struct {
	Title string `json:"title" form:"title" url:"title" validate:"required"`
	Body  string `json:"body" form:"body" url:"body" validate:"required"`
}
