package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"

	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/presentation/graph/generated"
)

// SendOtp is the resolver for the sendOTP field.
func (r *mutationResolver) SendOtp(ctx context.Context, phoneNumber string, flavour enums.Flavour) (string, error) {
	r.checkPreconditions()
	return r.makao.OTP.GenerateAndSendOTP(ctx, phoneNumber, flavour)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
