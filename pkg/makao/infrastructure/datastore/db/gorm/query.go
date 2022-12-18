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
