package enums

import (
	"fmt"
	"io"
	"strconv"
)

// IdentifierType is a list of application types.
type IdentifierType string

const (
	//IdentifierTypeNationalID represents the admin app
	IdentifierTypeNationalID IdentifierType = "NATIONAL_ID"
	// IdentifierTypePassport represents the tenant app user
	IdentifierTypePassport IdentifierType = "PASSPORT"
)

// IsValid returns true if a IdentifierType type is valid
func (f IdentifierType) IsValid() bool {
	switch f {
	case IdentifierTypeNationalID, IdentifierTypePassport:
		return true
	}
	return false
}

func (f IdentifierType) String() string {
	return string(f)
}

// UnmarshalGQL converts the supplied value to a IdentifierType type.
func (f *IdentifierType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*f = IdentifierType(str)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid IdentifierType", str)
	}
	return nil
}

// MarshalGQL writes the IdentifierType type to the supplied
func (f IdentifierType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(f.String()))
}
