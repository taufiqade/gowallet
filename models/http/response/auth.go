package response

// LoginResponse godoc
type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
