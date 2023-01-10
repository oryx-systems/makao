package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// UserMock struct implements mocks of user methods.
type UserMock struct {
	MockSearchUserFn func(ctx context.Context, searchTerm string) ([]*domain.User, error)
}

// NewUserMock initializes a new instance of user mock
func NewUserMock() *UserMock {
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

	return &UserMock{
		MockSearchUserFn: func(ctx context.Context, searchTerm string) ([]*domain.User, error) {
			return []*domain.User{
				user,
			}, nil
		},
	}
}

// SearchUser mocks the search user method
func (u *UserMock) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	return u.MockSearchUserFn(ctx, searchTerm)
}
