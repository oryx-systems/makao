package residence

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/application/dto"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore"
)

// UsecasesResidence holds all the usecase methods for the residence business logic
type UsecasesResidence interface {
	CreateResidence(ctx context.Context, input dto.ResidenceInput) (*domain.Residence, error)
}

// UsecasesResidenceImpl is the implementation of the Residence interface
type UsecasesResidenceImpl struct {
	Create datastore.Create
}

// NewResidence creates a new instance of the Residence usecase
func NewResidence(
	create datastore.Create,
) UsecasesResidence {
	return &UsecasesResidenceImpl{
		Create: create,
	}
}

// CreateResidence creates a new residence
func (r *UsecasesResidenceImpl) CreateResidence(ctx context.Context, input dto.ResidenceInput) (*domain.Residence, error) {
	residence := domain.Residence{
		Active:             true,
		Name:               input.Name,
		RegistrationNumber: input.RegistrationNumber,
		Location:           input.Location,
		LivingRoomsCount:   input.LivingRoomsCount,
		Owner:              input.Owner,
	}

	return r.Create.CreateResidence(ctx, residence)
}
