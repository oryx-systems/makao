package user

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/application/dto"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
)

// UseCasesUser represents all the user business logic
type UseCasesUser interface {
	Login(ctx context.Context, loginInput *dto.LoginInput) error
	RegisterUser(ctx context.Context, registerInput *dto.RegisterUserInput) error
}

// UseCasesUserImpl represents the user usecase implementation
type UseCasesUserImpl struct {
	infrastructure infrastructure.Datastore
}

// NewUseCasesUser initializes the new user implementation
func NewUseCasesUser(infra infrastructure.Datastore) UseCasesUser {
	return &UseCasesUserImpl{
		infrastructure: infra,
	}
}

// HandleIncomingMessages receives ang processes the incoming SMS data
func (u UseCasesUserImpl) Login(ctx context.Context, loginInput *dto.LoginInput) error {
	return nil
}

// HandleRegistration handles the user registration
func (u UseCasesUserImpl) RegisterUser(ctx context.Context, registerInput *dto.RegisterUserInput) error {
	user := &domain.User{
		FirstName:   registerInput.FirstName,
		MiddleName:  registerInput.MiddleName,
		LastName:    registerInput.LastName,
		Active:      true,
		Flavour:     registerInput.Flavour,
		UserName:    registerInput.UserName,
		DeviceToken: registerInput.DeviceToken,
		Residence:   registerInput.Residence,
	}

	if user.Flavour == enums.FlavourConsumer {
		user.UserType = "TENANT"
	} else {
		user.UserType = "MANAGER"
	}

	contact := &domain.Contact{
		Active:       true,
		ContactType:  "PHONE",
		ContactValue: registerInput.PhoneNumber,
		Flavour:      registerInput.Flavour,
	}

	identifier := &domain.Identifier{
		Active:          true,
		IdentifierType:  registerInput.IdentifierDocumentType,
		IdentifierValue: registerInput.IdentificationDocumentNumber,
	}

	return u.infrastructure.Create.RegisterUser(ctx, user, contact, identifier)
}
