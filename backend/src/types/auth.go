package types

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Handle   string `json:"handle"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
