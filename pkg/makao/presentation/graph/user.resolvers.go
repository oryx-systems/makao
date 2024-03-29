package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/oryx-systems/makao/pkg/makao/domain"
)

// ProceedWithResidence is the resolver for the proceedWithResidence field.
func (r *mutationResolver) ProceedWithResidence(ctx context.Context, residenceID string) (bool, error) {
	r.checkPreconditions()

	return r.makao.User.ProceedWithResidence(ctx, residenceID)
}

// AssignHouseToAUser is the resolver for the assignHouseToAUser field.
func (r *mutationResolver) AssignHouseToAUser(ctx context.Context, userID string, houseNumber string) (bool, error) {
	r.checkPreconditions()

	return r.makao.User.AssignHouseToAUser(ctx, userID, houseNumber)
}

// FreezeUser is the resolver for the freezeUser field.
func (r *mutationResolver) FreezeUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: FreezeUser - freezeUser"))
}

// UnfreezeUser is the resolver for the unfreezeUser field.
func (r *mutationResolver) UnfreezeUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnfreezeUser - unfreezeUser"))
}

// GetUserResidences is the resolver for the getUserResidences field.
func (r *queryResolver) GetUserResidences(ctx context.Context) ([]*domain.Residence, error) {
	r.checkPreconditions()

	return r.makao.User.GetUserResidences(ctx)
}

// SearchUser is the resolver for the searchUser field.
func (r *queryResolver) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	r.checkPreconditions()

	return r.makao.User.SearchUser(ctx, searchTerm)
}
