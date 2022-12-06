package infrastructure

import (
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db"
)

// ServiceInfrastructure is the interface that groups all the service's infrastructure
type Datastore struct {
	Create datastore.Create
	Query  datastore.Query
	Update datastore.Update
}

// Interactor ...
type Interactor struct {
	Datastore
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor() *Datastore {
	db := db.NewDBService()
	// ext := extension.NewExtension()

	return &Datastore{
		Create: db,
		Query:  db,
		Update: db,
	}
}
