package infrastructure

import (
	"github.com/oryx-systems/makao/pkg/makao/application/extension"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore"
)

// ServiceInfrastructure is the interface that groups all the service's infrastructure
type ServiceInfrastructure interface {
	datastore.Repository
	extension.Extension
}

// Interactor ...
type Interactor struct {
	datastore.DbServiceImpl
	extension.ExtImpl
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor() Interactor {
	db := datastore.NewDbService()
	ext := extension.NewExtension()

	return Interactor{
		*db,
		*ext,
	}
}
