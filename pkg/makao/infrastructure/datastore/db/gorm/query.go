package gorm

import (
	"context"
	"fmt"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"gorm.io/gorm/clause"
)

// Query holds all the database record query methods
type Query interface {
	GetUserProfileByUserID(ctx context.Context, userID *string) (*User, error)
	GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*User, error)
	GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*UserPIN, error)
	GetUserResidencesByUserID(ctx context.Context, userID string) ([]*UserResidence, error)
	GetResidenceByID(ctx context.Context, residenceID string) (*Residence, error)
	SearchUser(ctx context.Context, searchTerm string) ([]*User, error)
	GetHouseByNumber(ctx context.Context, houseNumber string) (*House, error)
	ListHousesInResidence(ctx context.Context, residenceID string) ([]*House, error)
}

// GetUserProfileByUserID fetches a user profile using the user ID
func (db *PGInstance) GetUserProfileByUserID(ctx context.Context, userID *string) (*User, error) {
	var user User
	if err := db.DB.Where(&User{ID: userID, Active: true}).Preload(clause.Associations).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by user ID %v: %v", userID, err)
	}
	return &user, nil
}

// GetUserProfileByPhoneNumber fetches a user profile using the phone number
func (db *PGInstance) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*User, error) {
	var user User

	if err := db.DB.Joins("JOIN makao_contact on makao_user.id = makao_contact.user_id").Where("makao_contact.contact_value = ? AND makao_contact.flavour = ?", phoneNumber, flavour).Preload(clause.Associations).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by phonenumber %v: %v", phoneNumber, err)
	}

	return &user, nil
}

// GetUserPINByUserID fetches a user's pin using the user ID and Flavour
func (db *PGInstance) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*UserPIN, error) {
	if !flavour.IsValid() {
		return nil, fmt.Errorf("flavour is not valid")
	}
	var pin UserPIN
	if err := db.DB.Where(&UserPIN{UserID: userID, Active: true, Flavour: flavour}).First(&pin).Error; err != nil {
		return nil, fmt.Errorf("failed to get pin: %v", err)
	}

	return &pin, nil
}

// GetUserResidencesByUserID fetches a user's residences using the user ID
func (db *PGInstance) GetUserResidencesByUserID(ctx context.Context, userID string) ([]*UserResidence, error) {
	var residence []*UserResidence
	if err := db.DB.Where(&UserResidence{UserID: userID}).Find(&residence).Error; err != nil {
		return nil, fmt.Errorf("failed to get residence: %v", err)
	}

	return residence, nil
}

// GetResidenceByID fetches a residence using the residence ID
func (db *PGInstance) GetResidenceByID(ctx context.Context, residenceID string) (*Residence, error) {
	var residence Residence
	if err := db.DB.Where(&Residence{ID: residenceID}).First(&residence).Error; err != nil {
		return nil, fmt.Errorf("failed to get residence: %v", err)
	}

	return &residence, nil
}

// SearchUser searches for a user using the search term
func (db *PGInstance) SearchUser(ctx context.Context, searchTerm string) ([]*User, error) {
	var users []*User
	if err := db.DB.Joins("JOIN makao_contact on makao_user.id = makao_contact.user_id").
		Joins("JOIN makao_identifier on makao_user.id = makao_identifier.user_id").
		Where("makao_contact.contact_value ILIKE ? OR makao_user.first_name ILIKE ? "+
			"OR makao_user.last_name ILIKE ? OR makao_user.username ILIKE ? OR makao_identifier.identifier_value ILIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").
		Where("makao_user.active = ?", true).
		Preload(clause.Associations).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to search user: %v", err)
	}

	return users, nil
}

// GetHouseByNumber fetches a house using the house number
func (db *PGInstance) GetHouseByNumber(ctx context.Context, houseNumber string) (*House, error) {
	var house House
	if err := db.DB.Where(&House{Number: houseNumber}).First(&house).Error; err != nil {
		return nil, fmt.Errorf("failed to get house: %v", err)
	}

	return &house, nil
}

// ListHousesInResidence lists all the houses in a residence
func (db *PGInstance) ListHousesInResidence(ctx context.Context, residenceID string) ([]*House, error) {
	var houses []*House

	if err := db.DB.Where(&House{ResidenceID: residenceID}).Find(&houses).Error; err != nil {
		return nil, fmt.Errorf("failed to get houses: %v", err)
	}

	return houses, nil
}
