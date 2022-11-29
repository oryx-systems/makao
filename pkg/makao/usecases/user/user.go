package user

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
)

// UseCasesUser represents all the user business logic
type UseCasesUser interface {
}

// UseCasesUserImpl represents the user usecase implementation
type UseCasesUserImpl struct {
	infrastructure infrastructure.ServiceInfrastructure
}

// NewUseCasesUser initializes the new user implementation
func NewUseCasesUser(infra infrastructure.ServiceInfrastructure) UseCasesUser {
	return &UseCasesUserImpl{
		infrastructure: infra,
	}
}

// HandleIncomingMessages receives ang processes the incoming SMS data
func (u UseCasesUserImpl) HandleIncomingMessages(ctx context.Context, payload interface{}) error {
	return nil
}
