package house

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/application/dto"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore"
)

// UsecaseHouse hold all the house usecase methods
type UsecaseHouse interface {
	CreateHouse(ctx context.Context, input *dto.HouseInput) (bool, error)
	GetHouseByNumber(ctx context.Context, houseNumber string) (*domain.House, error)
	ListHousesInResidence(ctx context.Context, residenceID string) ([]*domain.House, error)
}

// UsecaseHouseImpl represents the house usecase implementation
type UsecaseHouseImpl struct {
	Create datastore.Create
	Query  datastore.Query
}

// NewUsecaseHouse initializes the new house usecase implementation
func NewUsecaseHouse(
	create datastore.Create,
	query datastore.Query,
) UsecaseHouse {
	return &UsecaseHouseImpl{
		Create: create,
		Query:  query,
	}
}

// CreateHouse creates a new house
func (h *UsecaseHouseImpl) CreateHouse(ctx context.Context, input *dto.HouseInput) (bool, error) {
	house := &domain.House{
		Active:      true,
		Number:      input.Number,
		Category:    input.Category,
		Class:       input.Class,
		RentValue:   input.RentValue,
		ResidenceID: input.ResidenceID,
	}

	return h.Create.CreateHouse(ctx, house)
}

// GetHouseByNumber gets a house by its number
func (h *UsecaseHouseImpl) GetHouseByNumber(ctx context.Context, houseNumber string) (*domain.House, error) {
	return h.Query.GetHouseByNumber(ctx, houseNumber)
}

// ListHousesInResidence lists all the houses in a residence
func (h *UsecaseHouseImpl) ListHousesInResidence(ctx context.Context, residenceID string) ([]*domain.House, error) {
	return h.Query.ListHousesInResidence(ctx, residenceID)
}
