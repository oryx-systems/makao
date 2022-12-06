package domain

import "time"

type User struct {
	ID          string `gorm:"column:id:primary_key"`
	FirstName   string `gorm:"column:first_name"`
	MiddleName  string `gorm:"column:middle_name"`
	LastName    string `gorm:"column:last_name"`
	Active      bool   `gorm:"column:active"`
	Flavour     string `gorm:"column:flavour"`
	UserName    string `gorm:"column:username"`
	UserType    string `gorm:"column:user_type"`
	DeviceToken string `gorm:"column:device_token"`
	Residence   string `gorm:"column:residence"`
}

type Contact struct {
	ID           string `gorm:"column:id"`
	Active       bool   `gorm:"column:active"`
	ContactType  string `gorm:"column:contact_type"`
	ContactValue string `gorm:"column:contact_value"`
	Flavour      string `gorm:"column:flavour"`
	UserID       string `gorm:"column:user_id"`
}

type UserPIN struct {
	ID        string    `gorm:"column:id"`
	Active    bool      `gorm:"column:active"`
	Flavour   string    `gorm:"column:flavour"`
	ValidFrom time.Time `gorm:"column:valid_from"`
	ValidTo   time.Time `gorm:"column:valid_to"`
	HashedPIN string    `gorm:"column:hashed_pin"`
	Salt      string    `gorm:"column:salt"`
	UserID    string    `gorm:"column:user_id"`
}

type Identifier struct {
	ID              string `gorm:"column:id"`
	Active          bool   `gorm:"column:active"`
	IdentifierType  string `gorm:"column:identifier_type"`
	IdentifierValue string `gorm:"column:identifier_value"`
	UserID          string `gorm:"column:user_id"`
}

type Login struct {
	Phone string `json:"phone"`
	PIN   string `json:"pin"`
}

type LoginResponse struct {
	ID                string       `json:"id"`
	Username          string       `json:"username"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	UserContact       Contact      `json:"contact"`
	ManagedResidences []*Residence `json:"managed_residences"`
	AuthToken         string       `json:"auth_token"`
	RefreshToken      string       `json:"refresh_token"`
	CurrentResidence  Residence    `json:"current_residence"`
}
