package user

import (
	"context"
	"fmt"
	"time"

	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/application/dto"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/application/utils"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
)

// UseCasesUser represents all the user business logic
type UseCasesUser interface {
	Login(ctx context.Context, loginInput *dto.LoginInput) (*dto.LoginResponse, error)
	RegisterUser(ctx context.Context, registerInput *dto.RegisterUserInput) error
	SetUserPIN(ctx context.Context, input *dto.UserPINInput) (bool, error)
	GetUserResidences(ctx context.Context) ([]*domain.Residence, error)
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
func (u UseCasesUserImpl) Login(ctx context.Context, loginInput *dto.LoginInput) (*dto.LoginResponse, error) {
	user, err := u.infrastructure.Query.GetUserProfileByPhoneNumber(ctx, loginInput.PhoneNumber, loginInput.Flavour)
	if err != nil {
		return nil, err
	}

	userPIN, err := u.infrastructure.Query.GetUserPINByUserID(ctx, user.ID, user.Flavour)
	if err != nil {
		return nil, err
	}

	// If pin `ValidTo` field is in the past (expired). This means the user has to change their pin
	currentTime := time.Now()
	expired := currentTime.After(userPIN.ValidTo)
	if expired {
		return nil, fmt.Errorf("pin expired. Please change your pin")
	}

	matched := utils.ComparePIN(
		loginInput.PIN,
		userPIN.Salt,
		userPIN.HashedPIN,
		nil,
	)

	if !matched {
		return nil, fmt.Errorf("invalid pin")
	}

	tokenResponse, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		return nil, err
	}

	userToken, err := utils.ValidateJWTToken(tokenResponse.Token)
	if err != nil {
		return nil, err
	}

	user.AuthCredentials.IDToken = userToken.Token
	user.AuthCredentials.ExpiresIn = userToken.ExpiresIn

	return &dto.LoginResponse{
		UserProfile: user,
	}, nil
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

// SetUserPIN sets the user pin
func (u UseCasesUserImpl) SetUserPIN(ctx context.Context, input *dto.UserPINInput) (bool, error) {
	userProfile, err := u.infrastructure.Query.GetUserProfileByUserID(ctx, input.UserID)
	if err != nil {
		return false, fmt.Errorf("failed to get a user profile by user ID: %v", err)
	}

	err = utils.ValidatePIN(input.PIN)
	if err != nil {
		return false, err
	}

	salt, encryptedPIN := utils.EncryptPIN(input.PIN, nil)

	isMatch := utils.ComparePIN(input.ConfirmPIN, salt, encryptedPIN, nil)
	if !isMatch {
		return false, err
	}

	expiryDate, err := helpers.GetPinExpiryDate()
	if err != nil {
		return false, err
	}

	pinDataPayload := &domain.UserPIN{
		UserID:    userProfile.ID,
		HashedPIN: encryptedPIN,
		ValidFrom: time.Now(),
		ValidTo:   *expiryDate,
		Flavour:   input.Flavour,
		Active:    true,
		Salt:      salt,
	}

	_, err = u.infrastructure.Create.SavePIN(ctx, pinDataPayload)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUserResidences gets the user residences
func (u UseCasesUserImpl) GetUserResidences(ctx context.Context) ([]*domain.Residence, error) {
	uid, err := utils.GetLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}

	return u.infrastructure.Query.GetUserResidencesByUserID(ctx, uid)
}
