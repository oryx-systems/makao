package mock

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
)

// GormDatastoreMock is a mocks the database layer
type GormDatastoreMock struct {
	MockRegisterUserFn                func(ctx context.Context, user *gorm.User, contact *gorm.Contact, identifier *gorm.Identifier) error
	MockSaveOTPFn                     func(ctx context.Context, otp *gorm.OTP) error
	MockSavePINFn                     func(ctx context.Context, pinData *gorm.UserPIN) (bool, error)
	MockGetUserProfileByUserIDFn      func(ctx context.Context, userID *string) (*gorm.User, error)
	MockGetUserProfileByPhoneNumberFn func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*gorm.User, error)
	MockGetUserPINByUserIDFn          func(ctx context.Context, userID string, flavour enums.Flavour) (*gorm.UserPIN, error)
	MockGetUserResidencesByUserIDFn   func(ctx context.Context, userID string) ([]*gorm.UserResidence, error)
	MockGetResidenceByIDFn            func(ctx context.Context, residenceID string) (*gorm.Residence, error)
	MockInvalidatePINFn               func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error)
	MockCreateResidenceFn             func(ctx context.Context, payload gorm.Residence) (*gorm.Residence, error)
	MockSearchUserFn                  func(ctx context.Context, searchTerm string) ([]*gorm.User, error)
	MockCreateHouseFn                 func(ctx context.Context, house *gorm.House) error
	MockGetHouseByNumberFn            func(ctx context.Context, houseNumber string) (*gorm.House, error)
	MockUpdateUserFn                  func(ctx context.Context, user *gorm.User, updateData map[string]interface{}) (bool, error)
	MockListHousesInResidenceFn       func(ctx context.Context, residenceID string) ([]*gorm.House, error)
}

// NewGormDatastoreMock initializes a new GormDatastoreMock
func NewGormDatastoreMock() *GormDatastoreMock {
	UUID := uuid.New().String()

	residence := &gorm.Residence{
		ID:                 UUID,
		Active:             true,
		Name:               gofakeit.Name(),
		RegistrationNumber: gofakeit.Name(),
		Location:           gofakeit.Name(),
		LivingRoomsCount:   10,
		Owner:              gofakeit.Name(),
	}

	contact := &gorm.Contact{
		ID:           UUID,
		Active:       true,
		ContactType:  "PHONE",
		ContactValue: gofakeit.Phone(),
		Flavour:      enums.FlavourPro,
		UserID:       &UUID,
	}

	identifier := &gorm.Identifier{
		ID:              UUID,
		Active:          true,
		IdentifierType:  "NATIONAL_ID",
		IdentifierValue: "1234566789",
		UserID:          &UUID,
	}

	user := &gorm.User{
		ID:             &UUID,
		FirstName:      gofakeit.BeerAlcohol(),
		MiddleName:     gofakeit.BeerAlcohol(),
		LastName:       gofakeit.BeerAlcohol(),
		Active:         true,
		Flavour:        enums.FlavourPro,
		UserName:       gofakeit.BeerAlcohol(),
		UserType:       "STAFF",
		DeviceToken:    gofakeit.BeerAlcohol(),
		Residence:      uuid.New().String(),
		UserContact:    *contact,
		UserIdentifier: *identifier,
	}

	userResidence := &gorm.UserResidence{
		ID:          UUID,
		UserID:      UUID,
		ResidenceID: UUID,
	}

	return &GormDatastoreMock{
		MockRegisterUserFn: func(ctx context.Context, user *gorm.User, contact *gorm.Contact, identifier *gorm.Identifier) error {
			return nil
		},
		MockGetResidenceByIDFn: func(ctx context.Context, residenceID string) (*gorm.Residence, error) {
			return residence, nil
		},
		MockSaveOTPFn: func(ctx context.Context, otp *gorm.OTP) error {
			return nil
		},
		MockSavePINFn: func(ctx context.Context, pinData *gorm.UserPIN) (bool, error) {
			return true, nil
		},
		MockGetUserProfileByUserIDFn: func(ctx context.Context, userID *string) (*gorm.User, error) {
			return user, nil
		},
		MockGetUserProfileByPhoneNumberFn: func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*gorm.User, error) {
			return user, nil
		},
		MockSearchUserFn: func(ctx context.Context, searchTerm string) ([]*gorm.User, error) {
			return []*gorm.User{user}, nil
		},
		MockGetUserPINByUserIDFn: func(ctx context.Context, userID string, flavour enums.Flavour) (*gorm.UserPIN, error) {
			return &gorm.UserPIN{
				ID:        uuid.New().String(),
				Active:    true,
				Flavour:   flavour,
				ValidFrom: time.Now(),
				ValidTo:   time.Now().Add(time.Hour * 24 * 30),
				HashedPIN: "hashedPIN",
				Salt:      "salt",
				UserID:    userID,
			}, nil
		},
		MockGetUserResidencesByUserIDFn: func(ctx context.Context, userID string) ([]*gorm.UserResidence, error) {
			return []*gorm.UserResidence{
				userResidence,
			}, nil
		},
		MockInvalidatePINFn: func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
			return true, nil
		},
		MockCreateHouseFn: func(ctx context.Context, house *gorm.House) error {
			return nil
		},
		MockCreateResidenceFn: func(ctx context.Context, payload gorm.Residence) (*gorm.Residence, error) {
			return residence, nil
		},
		MockGetHouseByNumberFn: func(ctx context.Context, houseNumber string) (*gorm.House, error) {
			return &gorm.House{
				ID:        uuid.New().String(),
				Active:    true,
				Number:    houseNumber,
				Category:  "BEDSITTER",
				Class:     "CLASS_A",
				RentValue: 10,
			}, nil
		},
		MockUpdateUserFn: func(ctx context.Context, user *gorm.User, updateData map[string]interface{}) (bool, error) {
			return true, nil
		},
		MockListHousesInResidenceFn: func(ctx context.Context, residenceID string) ([]*gorm.House, error) {
			return []*gorm.House{
				{
					ID:        uuid.New().String(),
					Active:    true,
					Number:    "1",
					Category:  "BEDSITTER",
					Class:     "CLASS_A",
					RentValue: 10,
				},
			}, nil
		},
	}
}

// RegisterUser mocks the RegisterUser method
func (m *GormDatastoreMock) RegisterUser(ctx context.Context, user *gorm.User, contact *gorm.Contact, identifier *gorm.Identifier) error {
	return m.MockRegisterUserFn(ctx, user, contact, identifier)
}

// GetResidenceByID mocks the GetResidenceByID method
func (m *GormDatastoreMock) GetResidenceByID(ctx context.Context, residenceID string) (*gorm.Residence, error) {
	return m.MockGetResidenceByIDFn(ctx, residenceID)
}

// SaveOTP mocks the SaveOTP method
func (m *GormDatastoreMock) SaveOTP(ctx context.Context, otp *gorm.OTP) error {
	return m.MockSaveOTPFn(ctx, otp)
}

// SavePIN mocks the SavePIN method
func (m *GormDatastoreMock) SavePIN(ctx context.Context, pinData *gorm.UserPIN) (bool, error) {
	return m.MockSavePINFn(ctx, pinData)
}

// GetUserProfileByUserID mocks the GetUserProfileByUserID method
func (m *GormDatastoreMock) GetUserProfileByUserID(ctx context.Context, userID *string) (*gorm.User, error) {
	return m.MockGetUserProfileByUserIDFn(ctx, userID)
}

// GetUserProfileByPhoneNumber mocks the GetUserProfileByPhoneNumber method
func (m *GormDatastoreMock) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*gorm.User, error) {
	return m.MockGetUserProfileByPhoneNumberFn(ctx, phoneNumber, flavour)
}

// GetUserPINByUserID mocks the GetUserPINByUserID method
func (m *GormDatastoreMock) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*gorm.UserPIN, error) {
	return m.MockGetUserPINByUserIDFn(ctx, userID, flavour)
}

// GetUserResidencesByUserID mocks the GetUserResidencesByUserID method
func (m *GormDatastoreMock) GetUserResidencesByUserID(ctx context.Context, userID string) ([]*gorm.UserResidence, error) {
	return m.MockGetUserResidencesByUserIDFn(ctx, userID)
}

// InvalidatePIN mocks the InvalidatePIN method
func (m *GormDatastoreMock) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	return m.MockInvalidatePINFn(ctx, userID, flavour)
}

// CreateResidence mocks the CreateResidence method
func (m *GormDatastoreMock) CreateResidence(ctx context.Context, payload gorm.Residence) (*gorm.Residence, error) {
	return m.MockCreateResidenceFn(ctx, payload)
}

// SearchUser mocks the SearchUser method
func (m *GormDatastoreMock) SearchUser(ctx context.Context, searchTerm string) ([]*gorm.User, error) {
	return m.MockSearchUserFn(ctx, searchTerm)
}

// CreateHouse mocks the action of creating a house
func (m *GormDatastoreMock) CreateHouse(ctx context.Context, house *gorm.House) error {
	return m.MockCreateHouseFn(ctx, house)
}

// GetHouseByNumber mocks the action of getting a house by number
func (m *GormDatastoreMock) GetHouseByNumber(ctx context.Context, houseNumber string) (*gorm.House, error) {
	return m.MockGetHouseByNumberFn(ctx, houseNumber)
}

// UpdateUser mocks the action of updating a user
func (m *GormDatastoreMock) UpdateUser(ctx context.Context, user *gorm.User, updateData map[string]interface{}) (bool, error) {
	return m.MockUpdateUserFn(ctx, user, updateData)
}

// ListHousesInResidence mocks the action of listing houses in a residence
func (m *GormDatastoreMock) ListHousesInResidence(ctx context.Context, residenceID string) ([]*gorm.House, error) {
	return m.MockListHousesInResidenceFn(ctx, residenceID)
}
