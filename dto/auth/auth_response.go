package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

