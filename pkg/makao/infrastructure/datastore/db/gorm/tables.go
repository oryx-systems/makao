package gorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"gorm.io/gorm"
)

var (
	// DefaultResidence is the default residence for a user incase none is specified
	DefaultResidence = helpers.MustGetEnvVar("DEFAULT_RESIDENCE_ID")
)

// Base is the base table for all tables
type Base struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy *string   `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy *string   `gorm:"column:updated_by"`
}

// User models the system user
type User struct {
	Base

	ID               *string       `gorm:"column:id"`
	FirstName        string        `gorm:"column:first_name"`
	MiddleName       string        `gorm:"column:middle_name"`
	LastName         string        `gorm:"column:last_name"`
	Active           bool          `gorm:"column:active"`
	Flavour          enums.Flavour `gorm:"column:flavour"`
	UserName         string        `gorm:"column:username"`
	UserType         string        `gorm:"column:user_type"`
	DeviceToken      string        `gorm:"column:device_token"`
	Residence        string        `gorm:"column:residence"`
	UserContact      Contact       `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	UserIdentifier   Identifier    `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	CurrentResidence *string       `gorm:"column:current_residence"`
	CurrentHouse     *string       `gorm:"column:current_house"`
}

// BeforeCreate is a hook run before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	u.Base.CreatedAt = time.Now()
	u.Base.UpdatedAt = time.Now()
	u.ID = &id

	if u.Residence == "" {
		u.Residence = DefaultResidence
	}

	return
}

// TableName customizes how the table name is generated
func (User) TableName() string {
	return "makao_user"
}

// Contact is a contact model for a user
type Contact struct {
	Base

	ID           string        `gorm:"column:id"`
	Active       bool          `gorm:"column:active"`
	ContactType  string        `gorm:"column:contact_type"`
	ContactValue string        `gorm:"column:contact_value"`
	Flavour      enums.Flavour `gorm:"column:flavour"`
	UserID       *string       `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating user contact
func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	c.CreatedAt = time.Now()
	c.ID = id

	return
}

// TableName customizes how the table name is generated
func (Contact) TableName() string {
	return "makao_contact"
}

// UserPIN models the user's PIN table
type UserPIN struct {
	Base

	ID        string        `gorm:"column:id"`
	Active    bool          `gorm:"column:active"`
	Flavour   enums.Flavour `gorm:"column:flavour"`
	ValidFrom time.Time     `gorm:"column:valid_from"`
	ValidTo   time.Time     `gorm:"column:valid_to"`
	HashedPIN string        `gorm:"column:hashed_pin"`
	Salt      string        `gorm:"column:salt"`
	UserID    string        `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating user PIN
func (u *UserPIN) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	u.CreatedAt = time.Now()
	u.ID = id

	return
}

// TableName customizes how the table name is generated
func (UserPIN) TableName() string {
	return "makao_user_pin"
}

// OTP is model for one time password
type OTP struct {
	Base

	ID          string        `gorm:"column:id"`
	IsValid     bool          `gorm:"column:is_valid"`
	ValidUntil  time.Time     `gorm:"column:valid_until"`
	PhoneNumber string        `gorm:"column:phone_number"`
	OTP         string        `gorm:"column:otp"`
	Flavour     enums.Flavour `gorm:"column:flavour"`
	Medium      string        `gorm:"column:medium"`
	UserID      string        `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating an OTP
func (o *OTP) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (OTP) TableName() string {
	return "makao_user_otp"
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

// BeforeCreate is a hook run before creating a residence
func (r *Residence) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreatedAt = time.Now()
	r.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (Residence) TableName() string {
	return "makao_residence"
}

// Identifiers models the identifier that may be used in the system
type Identifier struct {
	Base

	ID              string               `gorm:"column:id"`
	Active          bool                 `gorm:"column:active"`
	IdentifierType  enums.IdentifierType `gorm:"column:identifier_type"`
	IdentifierValue string               `gorm:"column:identifier_value"`
	UserID          *string              `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating an identifier
func (i *Identifier) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	i.CreatedAt = time.Now()
	i.ID = id
	return
}

// TableName customizes how the table name is generated
func (Identifier) TableName() string {
	return "makao_identifier"
}

// House models the relationship between a tenant and the living house
type HouseClient struct {
	ID       string `gorm:"column:id"`
	HouseID  string `gorm:"column:house_id"`
	TenantID string `gorm:"column:tenant_id"`
}

// BeforeCreate is a hook run before creating a house client
func (h *HouseClient) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (HouseClient) TableName() string {
	return "makao_house_client"
}

// House models the datastore entity for a house
type House struct {
	Base

	ID          string  `gorm:"column:id;primary_key"`
	Active      bool    `gorm:"column:active"`
	Number      string  `gorm:"column:number"`
	Category    string  `gorm:"column:category"`
	Class       string  `gorm:"column:class"` // applicable where houses maybe charged differently due to size
	RentValue   float64 `gorm:"column:rent_value"`
	State       string  `gorm:"column:state"`
	ResidenceID string  `gorm:"column:residence_id"`
}

// BeforeCreate is a hook run before creating a house
func (h *House) BeforeCreate(tx *gorm.DB) (err error) {
	h.CreatedAt = time.Now()
	h.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (House) TableName() string {
	return "makao_house"
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

// BeforeCreate is a hook run before creating a bill
func (b *Bill) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (Bill) TableName() string {
	return "makao_bill"
}

// UserResidence models the relationship between a user and the residence they live in
type UserResidence struct {
	ID          string `gorm:"column:id"`
	UserID      string `gorm:"column:user_id"`
	ResidenceID string `gorm:"column:residence_id"`
}

// BeforeCreate is a hook run before creating a house client
func (ur *UserResidence) BeforeCreate(tx *gorm.DB) (err error) {
	ur.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (UserResidence) TableName() string {
	return "makao_user_residence"
}
