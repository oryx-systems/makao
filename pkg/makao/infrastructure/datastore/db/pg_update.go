package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
)

// InvalidatePIN invalidates a pin that is linked to the user profile.
// This is done by toggling the IsValid field to false
func (d *DbServiceImpl) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	if userID == "" {
		return false, fmt.Errorf("userID cannot be empty")
	}

	return d.Repository.InvalidatePIN(ctx, userID, flavour)
}
