package renderings

type UserResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id,omitempty"`
	Email   string `json:"email,omitempty"`
}
