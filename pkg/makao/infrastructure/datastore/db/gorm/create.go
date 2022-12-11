package gorm

import (
	"context"
)

// Create holds all the database record creation methods
type Create interface {
	RegisterUser(ctx context.Context, user *User, contact *Contact, identifier *Identifier) error
	SaveOTP(ctx context.Context, otp *OTP) error
}

// RegisterUser creates a new user record.
// The user can be a resident or a staff member
func (db *PGInstance) RegisterUser(ctx context.Context, user *User, contact *Contact, identifier *Identifier) error {
	tx := db.DB.WithContext(ctx).Begin()

	// create user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create contact
	contact.UserID = user.ID
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create identifier
	identifier.UserID = user.ID
	if err := tx.Create(&identifier).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// SaveOTP saves an OTP in the database
func (db *PGInstance) SaveOTP(ctx context.Context, otp *OTP) error {
	if err := db.DB.WithContext(ctx).Create(&otp).Error; err != nil {
		return err
	}

	return nil
}
