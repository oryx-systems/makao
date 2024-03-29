package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// DataStoreMock is a mock implementation of the datastore interface
type DataStoreMock struct {
	MockRegisterUserFn                func(ctx context.Context, user *domain.User, contact *domain.Contact, identifier *domain.Identifier) error
	MockSaveOTPFn                     func(ctx context.Context, otp *domain.OTP) error
	MockSavePINFn                     func(ctx context.Context, pinInput *domain.UserPIN) (bool, error)
	MockGetUserProfileByUserIDFn      func(ctx context.Context, userID string) (*domain.User, error)
	MockGetUserProfileByPhoneNumberFn func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error)
	MockGetUserPINByUserIDFn          func(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error)
	MockGetUserResidencesByUserIDFn   func(ctx context.Context, userID string) ([]*domain.Residence, error)
	MockInvalidatePINFn               func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error)
	MockCreateResidenceFn             func(ctx context.Context, payload domain.Residence) (*domain.Residence, error)
	MockSearchUserFn                  func(ctx context.Context, searchTerm string) ([]*domain.User, error)
	MockUpdateUserFn                  func(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error)
	MockListHousesInResidenceFn       func(ctx context.Context, residenceID string) ([]*domain.House, error)
}

// NewDataStoreMock returns a new instance of the mock datastore
func NewDataStoreMock() *DataStoreMock {
	user := &domain.User{
		ID:         uuid.New().String(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.BeerAlcohol(),
		LastName:   gofakeit.LastName(),
		Active:     true,
		Flavour:    enums.FlavourPro,
		UserName:   gofakeit.Username(),
		UserType:   "TENANT",
		UserIdentifier: domain.Identifier{
			ID: uuid.New().String(),
		},
		UserContact: domain.Contact{
			ID: uuid.New().String(),
		},
		DeviceToken: uuid.New().String(),
		Residence:   uuid.New().String(),
	}

	return &DataStoreMock{
		MockRegisterUserFn: func(ctx context.Context, user *domain.User, contact *domain.Contact, identifier *domain.Identifier) error {
			return nil
		},
		MockSaveOTPFn: func(ctx context.Context, otp *domain.OTP) error {
			return nil
		},
		MockSavePINFn: func(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
			return true, nil
		},
		MockGetUserProfileByUserIDFn: func(ctx context.Context, userID string) (*domain.User, error) {
			return nil, nil
		},
		MockGetUserProfileByPhoneNumberFn: func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
			return nil, nil
		},
		MockGetUserPINByUserIDFn: func(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
			return nil, nil
		},
		MockGetUserResidencesByUserIDFn: func(ctx context.Context, userID string) ([]*domain.Residence, error) {
			return nil, nil
		},
		MockInvalidatePINFn: func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
			return true, nil
		},
		MockCreateResidenceFn: func(ctx context.Context, payload domain.Residence) (*domain.Residence, error) {
			return &domain.Residence{
				ID:                 uuid.New().String(),
				Active:             true,
				Name:               gofakeit.BeerName(),
				RegistrationNumber: gofakeit.BeerName(),
				Location:           gofakeit.Address().City,
				LivingRoomsCount:   20,
				Owner:              uuid.New().String(),
			}, nil
		},
		MockSearchUserFn: func(ctx context.Context, searchTerm string) ([]*domain.User, error) {
			return []*domain.User{
				user,
			}, nil
		},
		MockUpdateUserFn: func(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error) {
			return true, nil
		},
		MockListHousesInResidenceFn: func(ctx context.Context, residenceID string) ([]*domain.House, error) {
			return []*domain.House{
				{
					ID:          uuid.New().String(),
					Active:      true,
					Number:      "A2",
					Category:    "BEDSITTER",
					Class:       "A",
					RentValue:   2000,
					State:       "OCCUPIED",
					ResidenceID: uuid.New().String(),
				},
			}, nil
		},
	}
}

// RegisterUser mocks the RegisterUser method
func (m *DataStoreMock) RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact, identifier *domain.Identifier) error {
	return m.MockRegisterUserFn(ctx, user, contact, identifier)
}

// SaveOTP mocks the SaveOTP method
func (m *DataStoreMock) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	return m.MockSaveOTPFn(ctx, otp)
}

// SavePIN mocks the SavePIN method
func (m *DataStoreMock) SavePIN(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
	return m.MockSavePINFn(ctx, pinInput)
}

// GetUserProfileByUserID mocks the GetUserProfileByUserID method
func (m *DataStoreMock) GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error) {
	return m.MockGetUserProfileByUserIDFn(ctx, userID)
}

// GetUserProfileByPhoneNumber mocks the GetUserProfileByPhoneNumber method
func (m *DataStoreMock) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
	return m.MockGetUserProfileByPhoneNumberFn(ctx, phoneNumber, flavour)
}

// GetUserPINByUserID mocks the GetUserPINByUserID method
func (m *DataStoreMock) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
	return m.MockGetUserPINByUserIDFn(ctx, userID, flavour)
}

// GetUserResidencesByUserID mocks the GetUserResidencesByUserID method
func (m *DataStoreMock) GetUserResidencesByUserID(ctx context.Context, userID string) ([]*domain.Residence, error) {
	return m.MockGetUserResidencesByUserIDFn(ctx, userID)
}

// InvalidatePIN mocks the InvalidatePIN method
func (m *DataStoreMock) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	return m.MockInvalidatePINFn(ctx, userID, flavour)
}

// CreateResidence mocks the CreateResidence method
func (m *DataStoreMock) CreateResidence(ctx context.Context, payload domain.Residence) (*domain.Residence, error) {
	return m.MockCreateResidenceFn(ctx, payload)
}

// SearchUser mocks the SearchUser method
func (m *DataStoreMock) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	return m.MockSearchUserFn(ctx, searchTerm)
}

// UpdateUser mocks the UpdateUser method
func (m *DataStoreMock) UpdateUser(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error) {
	return m.MockUpdateUserFn(ctx, user, updateData)
}

// ListHousesInResidence mocks the ListHousesInResidence method
func (m *DataStoreMock) ListHousesInResidence(ctx context.Context, residenceID string) ([]*domain.House, error) {
	return m.MockListHousesInResidenceFn(ctx, residenceID)
}
