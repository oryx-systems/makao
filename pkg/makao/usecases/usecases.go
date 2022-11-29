package usecases

import (
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
	"github.com/oryx-systems/makao/pkg/makao/usecases/user"
)

// Usecases manages the usecases intrefaces
type Usecases interface {
	user.UseCasesUser
}

// Interactor is an implementation of the usecases interface
type Interactor struct {
	infrastructure.ServiceInfrastructure
	user.UseCasesUser
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewUseCasesInteractor(
	infrastructure infrastructure.ServiceInfrastructure,
) Usecases {
	user := user.NewUseCasesUser(infrastructure)

	impl := Interactor{
		infrastructure,
		user,
	}

	return impl
}
