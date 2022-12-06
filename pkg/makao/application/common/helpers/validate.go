package helpers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/oryx-systems/makao/pkg/makao/application/common"
	"github.com/oryx-systems/makao/pkg/makao/application/dto"
	"github.com/ttacon/libphonenumber"
)

func ValidateRegistrationInput(payload *dto.RegisterUserInput) error {
	if payload.IdentifierDocumentType == "" {
		return fmt.Errorf("identifier type is required")
	} else if payload.FirstName == "" {
		return fmt.Errorf("first name is required")
	} else if payload.LastName == "" {
		return fmt.Errorf("last name is required")
	} else if payload.Flavour == "" {
		return fmt.Errorf("flavour is required")
	} else if payload.UserName == "" {
		return fmt.Errorf("username is required")
	} else if payload.DeviceToken == "" {
		return fmt.Errorf("device token is required")
	}

	return nil
}

// IsMSISDNValid uses regular expression to validate the a phone number
func IsMSISDNValid(msisdn string) bool {
	if len(msisdn) < 10 {
		return false
	}
	reKen := regexp.MustCompile(`^(?:254|\+254|0)?((7|1)(?:(?:[129][0-9])|(?:0[0-8])|(4[0-1]))[0-9]{6})$`)
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !reKen.MatchString(msisdn) {
		return re.MatchString(msisdn)
	}
	return reKen.MatchString(msisdn)
}

// NormalizeMSISDN validates the input phone number.
// For valid phone numbers, it normalizes them to international format
// e.g +2547........
func NormalizeMSISDN(msisdn string) (*string, error) {
	if !IsMSISDNValid(msisdn) {
		return nil, fmt.Errorf("invalid phone number: %s", msisdn)
	}
	num, err := libphonenumber.Parse(msisdn, common.DefaultRegion)
	if err != nil {
		return nil, err
	}
	formatted := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	cleaned := strings.ReplaceAll(formatted, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	return &cleaned, nil
}
