package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// GetUserProfileByUserID fetches and returns a userprofile using their user ID
func (d *DbServiceImpl) GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := d.query.GetUserProfileByUserID(ctx, &userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by user ID: %v", err)
	}

	contact := &domain.Contact{
		ID:           user.UserContact.ID,
		Active:       user.UserContact.Active,
		ContactType:  user.UserContact.ContactType,
		ContactValue: user.UserContact.ContactValue,
		Flavour:      user.UserContact.Flavour,
		UserID:       *user.ID,
	}

	identifier := &domain.Identifier{
		ID:              user.UserIdentifier.ID,
		Active:          user.UserIdentifier.Active,
		IdentifierType:  user.UserIdentifier.IdentifierType,
		IdentifierValue: user.UserIdentifier.IdentifierValue,
		UserID:          *user.ID,
	}

	return &domain.User{
		ID:             *user.ID,
		UserIdentifier: *identifier,
		UserContact:    *contact,
		FirstName:      user.FirstName,
		MiddleName:     user.MiddleName,
		LastName:       user.LastName,
		Active:         user.Active,
		Flavour:        user.Flavour,
		UserName:       user.UserName,
		UserType:       user.UserType,
		DeviceToken:    user.DeviceToken,
		Residence:      user.Residence,
	}, nil
}

// GetUserProfileByPhoneNumber fetches and returns a userprofile using their phone number
func (d *DbServiceImpl) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
	user, err := d.query.GetUserProfileByPhoneNumber(ctx, phoneNumber, flavour)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by phonenumber: %v", err)
	}

	identifier := &domain.Identifier{
		ID:              user.UserIdentifier.ID,
		Active:          user.UserIdentifier.Active,
		IdentifierType:  user.UserIdentifier.IdentifierType,
		IdentifierValue: user.UserIdentifier.IdentifierValue,
		UserID:          *user.ID,
	}

	contact := &domain.Contact{
		ID:           user.UserContact.ID,
		Active:       user.UserContact.Active,
		ContactType:  user.UserContact.ContactType,
		ContactValue: user.UserContact.ContactValue,
		Flavour:      user.UserContact.Flavour,
		UserID:       *user.ID,
	}

	return &domain.User{
		ID:             *user.ID,
		UserIdentifier: *identifier,
		UserContact:    *contact,
		FirstName:      user.FirstName,
		MiddleName:     user.MiddleName,
		LastName:       user.LastName,
		Active:         user.Active,
		Flavour:        user.Flavour,
		UserName:       user.UserName,
		UserType:       user.UserType,
		DeviceToken:    user.DeviceToken,
		Residence:      user.Residence,
	}, nil
}

// GetUserPINByUserID fetches and returns a user PIN using their user ID
func (d *DbServiceImpl) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
	if userID == "" {
		return nil, fmt.Errorf("user id cannot be empty")
	}
	pinData, err := d.query.GetUserPINByUserID(ctx, userID, flavour)
	if err != nil {
		return nil, fmt.Errorf("failed query and retrieve user PIN data: %s", err)
	}

	return &domain.UserPIN{
		UserID:    pinData.UserID,
		HashedPIN: pinData.HashedPIN,
		ValidFrom: pinData.ValidFrom,
		ValidTo:   pinData.ValidTo,
		Flavour:   pinData.Flavour,
		Active:    pinData.Active,
		Salt:      pinData.Salt,
	}, nil
}

// GetUserResidencesByUserID fetches and returns a user residence using their user ID
func (d *DbServiceImpl) GetUserResidencesByUserID(ctx context.Context, userID string) ([]*domain.Residence, error) {
	var userResidences []*domain.Residence

	records, err := d.query.GetUserResidencesByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user residences by user ID: %v", err)
	}

	for _, record := range records {
		residence, err := d.query.GetResidenceByID(ctx, record.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get residence by ID: %v", err)
		}

		userResidences = append(userResidences, &domain.Residence{
			ID:                 residence.ID,
			Active:             residence.Active,
			Name:               residence.Name,
			RegistrationNumber: residence.RegistrationNumber,
			Location:           residence.Location,
			LivingRoomsCount:   residence.LivingRoomsCount,
			Owner:              residence.Owner,
		})
	}

	return userResidences, nil
}

// SearchUser searches for users in the system using a search term
func (d *DbServiceImpl) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	var users []*domain.User

	records, err := d.query.SearchUser(ctx, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search user: %v", err)
	}

	for _, record := range records {
		contact := &domain.Contact{
			ID:           record.UserContact.ID,
			Active:       record.UserContact.Active,
			ContactType:  record.UserContact.ContactType,
			ContactValue: record.UserContact.ContactValue,
			Flavour:      record.UserContact.Flavour,
			UserID:       *record.ID,
		}

		identifier := &domain.Identifier{
			ID:              record.UserIdentifier.ID,
			Active:          record.UserIdentifier.Active,
			IdentifierType:  record.UserIdentifier.IdentifierType,
			IdentifierValue: record.UserIdentifier.IdentifierValue,
			UserID:          *record.ID,
		}

		users = append(users, &domain.User{
			ID:             *record.ID,
			UserIdentifier: *identifier,
			UserContact:    *contact,
			FirstName:      record.FirstName,
			MiddleName:     record.MiddleName,
			LastName:       record.LastName,
			Active:         record.Active,
			Flavour:        record.Flavour,
			UserName:       record.UserName,
			UserType:       record.UserType,
			DeviceToken:    record.DeviceToken,
			Residence:      record.Residence,
		})
	}

	return users, nil
}

// GetHouseByNumber returns a house by its number
func (d *DbServiceImpl) GetHouseByNumber(ctx context.Context, houseNumber string) (*domain.House, error) {
	house, err := d.query.GetHouseByNumber(ctx, houseNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get house by number: %v", err)
	}

	return &domain.House{
		ID:          house.ID,
		Active:      house.Active,
		Number:      house.Number,
		Category:    house.Category,
		Class:       house.Class,
		RentValue:   house.RentValue,
		State:       house.State,
		ResidenceID: house.ResidenceID,
	}, nil
}

// ListHousesInResidence returns a list of houses in a residence
func (d *DbServiceImpl) ListHousesInResidence(ctx context.Context, residenceID string) ([]*domain.House, error) {
	records, err := d.query.ListHousesInResidence(ctx, residenceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list houses in residence: %v", err)
	}

	var houses []*domain.House
	for _, record := range records {
		houses = append(houses, &domain.House{
			ID:          record.ID,
			Active:      record.Active,
			Number:      record.Number,
			Category:    record.Category,
			Class:       record.Class,
			RentValue:   record.RentValue,
			State:       record.State,
			ResidenceID: record.ResidenceID,
		})
	}

	return houses, nil
}
