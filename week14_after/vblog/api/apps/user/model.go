package user


func NewAuthRequest()*AuthRequest{
	return &AuthRequest{}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
