package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
)

// RegisterUser registers a new user in the database
func (d *DbServiceImpl) RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact, identifier *domain.Identifier) error {
	usr := &gorm.User{
		FirstName:   user.FirstName,
		MiddleName:  user.MiddleName,
		LastName:    user.LastName,
		Active:      true,
		Flavour:     user.Flavour,
		UserName:    user.UserName,
		UserType:    user.UserType,
		DeviceToken: user.DeviceToken,
		Residence:   user.Residence,
	}

	ct := &gorm.Contact{
		Active:       true,
		ContactType:  contact.ContactType,
		ContactValue: contact.ContactValue,
		Flavour:      contact.Flavour,
		UserID:       usr.ID,
	}

	id := &gorm.Identifier{
		Active:          true,
		IdentifierType:  identifier.IdentifierType,
		IdentifierValue: identifier.IdentifierValue,
		UserID:          usr.ID,
	}

	return d.create.RegisterUser(ctx, usr, ct, id)
}

// SaveOTP saves an OTP in the database
func (d *DbServiceImpl) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	otpData := &gorm.OTP{
		IsValid:     otp.IsValid,
		ValidUntil:  otp.ValidUntil,
		PhoneNumber: otp.PhoneNumber,
		OTP:         otp.OTP,
		Flavour:     otp.Flavour,
		Medium:      otp.Medium,
		UserID:      otp.UserID,
	}

	return d.create.SaveOTP(ctx, otpData)
}

// SavePIN saves a PIN in the database
func (d *DbServiceImpl) SavePIN(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
	pinObj := &gorm.UserPIN{
		UserID:    pinInput.UserID,
		HashedPIN: pinInput.HashedPIN,
		ValidFrom: pinInput.ValidFrom,
		ValidTo:   pinInput.ValidTo,
		Active:    pinInput.Active,
		Flavour:   pinInput.Flavour,
		Salt:      pinInput.Salt,
	}

	_, err := d.create.SavePIN(ctx, pinObj)
	if err != nil {
		return false, fmt.Errorf("failed to save user pin: %v", err)
	}

	return true, nil
}

// CreateHouse is used to create a house record in the database
func (d *DbServiceImpl) CreateHouse(ctx context.Context, house *domain.House) (bool, error) {
	h := &gorm.House{
		Active:      house.Active,
		Number:      house.Number,
		Category:    house.Category,
		Class:       house.Class,
		RentValue:   house.RentValue,
		State:       house.State,
		ResidenceID: house.ResidenceID,
	}

	err := d.create.CreateHouse(ctx, h)
	if err != nil {
		return false, fmt.Errorf("failed to create house: %v", err)
	}

	return true, nil
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
