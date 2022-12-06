package graph

import (
	"context"
	"log"

	"github.com/oryx-systems/makao/pkg/makao/usecases"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	makao *usecases.Makao
}

// NewResolver initializes a new resolver
func NewResolver(ctx context.Context, makao usecases.Makao) (*Resolver, error) {
	return &Resolver{
		makao: &makao,
	}, nil
}

func (r Resolver) checkPreconditions() {
	if r.makao == nil {
		log.Panicf("expected makao usecases to be defined resolver")
	}
}
