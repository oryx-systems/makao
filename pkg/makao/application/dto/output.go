package dto

import (
	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// LoginResponse represents the login response
type LoginResponse struct {
	UserProfile *domain.User `json:"user_profile"`
}
