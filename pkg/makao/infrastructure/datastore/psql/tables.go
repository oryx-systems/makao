package psql

import "time"

// Base is the base table for all tables
type Base struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

// User models the system user
type User struct {
	Base

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

// Contact is a contact model for a user
type Contact struct {
	Base

	ID           string `gorm:"column:id"`
	Active       bool   `gorm:"column:active"`
	ContactType  string `gorm:"column:contact_type"`
	ContactValue string `gorm:"column:contact_value"`
	Flavour      string `gorm:"column:flavour"`
	UserID       string `gorm:"column:user_id"`
}

// UserPIN models the user's PIN table
type UserPIN struct {
	Base

	ID        string    `gorm:"column:id"`
	Active    bool      `gorm:"column:active"`
	Flavour   string    `gorm:"column:flavour"`
	ValidFrom time.Time `gorm:"column:valid_from"`
	ValidTo   time.Time `gorm:"column:valid_to"`
	HashedPIN string    `gorm:"column:hashed_pin"`
	Salt      string    `gorm:"column:salt"`
	UserID    string    `gorm:"column:user_id"`
}

// OTP is model for one time password
type OTP struct {
	Base

	ID          string    `gorm:"column:id"`
	IsValid     bool      `gorm:"column:is_valid"`
	ValidUntil  time.Time `gorm:"column:valid_until"`
	PhoneNumber string    `gorm:"column:phone_number"`
	OTP         string    `gorm:"column:otp"`
	Flavour     string    `gorm:"column:flavour"`
	Medium      string    `gorm:"column:medium"`
	UserID      string    `gorm:"column:user_id"`
}

// Residence models the residence's table
type Residence struct {
	Base

	ID                 string `gorm:"column:id"`
	Active             bool   `gorm:"column:active"`
	Name               string `gorm:"column:name"`
	RegistrationNumber string `gorm:"column:registration_number"`
	Location           string `gorm:"column:location"`
	LivingRoomsCount   int    `gorm:"column:living_rooms_count"`
	Owner              string `gorm:"column:owner"`
}

// Identifiers models the identifier that may be used in the system
type Identifier struct {
	Base

	ID              string `gorm:"column:id"`
	Active          bool   `gorm:"column:active"`
	IdentifierType  string `gorm:"column:identifier_type"`
	IdentifierValue string `gorm:"column:identifier_value"`
	UserID          string `gorm:"column:user_id"`
}

// House models the relationship between a tenant and the living house
type HouseClient struct {
	ID       string `gorm:"column:id"`
	HouseID  string `gorm:"column:house_id"`
	TenantID string `gorm:"column:tenant_id"`
}

// House models the datastore entity for a house
type House struct {
	Base

	ID          string  `gorm:"column:id;primary_key"`
	Active      bool    `gorm:"column:active"`
	HouseNumber string  `gorm:"column:house_number"`
	Category    string  `gorm:"column:category"`
	Class       string  `gorm:"column:class"` // applicable where houses maybe charged differently due to size
	RentValue   float64 `gorm:"column:rent_value"`
}

// Bill represents a billing model
type Bill struct {
	Base

	ID      string  `gorm:"column:id"`
	Active  bool    `gorm:"column:active"`
	Type    string  `gorm:"column:type"`
	Amount  float64 `gorm:"column:amount"`
	Penalty float64 `gorm:"column:penalty"`
	UserID  string  `gorm:"column:user_id"`
}
