package login

import (
	"github.com/scuba13/AmacoonServices/internal/owner"
)


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Owner *owner.Owner
	Token string
}
