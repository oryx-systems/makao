package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// HouseMock struct implements mocks of user methods.
type HouseMock struct {
	MockListHousesInResidenceFn func(ctx context.Context, residenceID string) ([]*domain.House, error)
}

// NewHouseMock initializes a new instance of user mock
func NewHouseMock() *HouseMock {
	house := &domain.House{
		ID:          uuid.New().String(),
		Active:      true,
		Number:      "A3",
		Category:    "SINGLE",
		Class:       "A",
		RentValue:   3000,
		State:       "OCCUPIED",
		ResidenceID: uuid.New().String(),
	}

	return &HouseMock{
		MockListHousesInResidenceFn: func(ctx context.Context, residenceID string) ([]*domain.House, error) {
			return []*domain.House{
				house,
			}, nil
		},
	}
}

// ListHousesInResidence mocks the search user method
func (u *HouseMock) ListHousesInResidence(ctx context.Context, residenceID string) ([]*domain.House, error) {
	return u.MockListHousesInResidenceFn(ctx, residenceID)
}
