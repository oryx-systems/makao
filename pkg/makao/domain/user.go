package domain

import (
	"time"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
)

type User struct {
	ID              string          `json:"id"`
	FirstName       string          `json:"firstName"`
	MiddleName      string          `json:"middleName"`
	LastName        string          `json:"lastName"`
	Active          bool            `json:"active"`
	Flavour         enums.Flavour   `json:"flavour"`
	UserName        string          `json:"username"`
	UserType        string          `json:"userType"`
	UserIdentifier  Identifier      `json:"userIdentifier"`
	UserContact     Contact         `json:"userContact"`
	DeviceToken     string          `json:"deviceToken"`
	Residence       string          `json:"residence"`
	AuthCredentials AuthCredentials `json:"authCredentials"`
}

type AuthCredentials struct {
	RefreshToken string    `json:"refreshToken"`
	IDToken      string    `json:"idToken"`
	ExpiresIn    time.Time `json:"expiresIn"`
}

type Contact struct {
	ID           string        `json:"id"`
	Active       bool          `json:"active"`
	ContactType  string        `json:"contact_type"`
	ContactValue string        `json:"contact_value"`
	Flavour      enums.Flavour `json:"flavour"`
	UserID       string        `json:"user_id"`
}

type UserPIN struct {
	ID        string        `json:"id"`
	Active    bool          `json:"active"`
	Flavour   enums.Flavour `json:"flavour"`
	ValidFrom time.Time     `json:"valid_from"`
	ValidTo   time.Time     `json:"valid_to"`
	HashedPIN string        `json:"hashed_pin"`
	Salt      string        `json:"salt"`
	UserID    string        `json:"user_id"`
}

type Identifier struct {
	ID              string               `json:"id"`
	Active          bool                 `json:"active"`
	IdentifierType  enums.IdentifierType `json:"identifierType"`
	IdentifierValue string               `json:"identifierValue"`
	UserID          string               `json:"userID"`
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
