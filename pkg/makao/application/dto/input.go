package dto

import "github.com/oryx-systems/makao/pkg/makao/application/enums"

// LoginInput represents the login input
type LoginInput struct {
	Phone string `json:"phone"`
	PIN   string `json:"pin"`
}

// RegisterUserInput represents the register user input
type RegisterUserInput struct {
	IdentifierDocumentType       enums.IdentifierType `json:"identifier_type"`
	IdentificationDocumentNumber string               `json:"identification_number"`
	FirstName                    string               `json:"first_name"`
	MiddleName                   string               `json:"middle_name"`
	LastName                     string               `json:"last_name"`
	PhoneNumber                  string               `json:"phone_number"`
	Flavour                      enums.Flavour        `json:"flavour"`
	UserName                     string               `json:"username"`
	DeviceToken                  string               `json:"device_token"`
	Residence                    string               `json:"residence"`
}
