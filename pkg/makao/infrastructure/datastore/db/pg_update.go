package db

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
)

// InvalidatePIN invalidates a pin that is linked to the user profile.
// This is done by toggling the IsValid field to false
func (d *DbServiceImpl) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	return d.update.InvalidatePIN(ctx, userID, flavour)
}

// CreateResidence creates a new residence
func (d *DbServiceImpl) CreateResidence(ctx context.Context, payload domain.Residence) (*domain.Residence, error) {
	data := &gorm.Residence{
		Active:             payload.Active,
		Name:               payload.Name,
		RegistrationNumber: payload.RegistrationNumber,
		Location:           payload.Location,
		LivingRoomsCount:   payload.LivingRoomsCount,
		Owner:              payload.Owner,
	}

	residence, err := d.create.CreateResidence(ctx, *data)
	if err != nil {
		return nil, err
	}

	return &domain.Residence{
		ID:                 residence.ID,
		Active:             residence.Active,
		Name:               residence.Name,
		RegistrationNumber: residence.RegistrationNumber,
		Location:           residence.Location,
		LivingRoomsCount:   residence.LivingRoomsCount,
		Owner:              residence.Owner,
	}, nil
}
