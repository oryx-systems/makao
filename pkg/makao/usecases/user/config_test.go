package user_test

// import (
// 	"github.com/oryx-systems/makao/pkg/makao/application/extension"
// 	extensionMock "github.com/oryx-systems/makao/pkg/makao/application/extension/mock"
// 	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
// 	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore"
// 	datastoreMock "github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/mock"
// 	"github.com/oryx-systems/makao/pkg/makao/usecases"
// )

// var fakeRepo datastoreMock.DataStoreMock
// var fakeExtension extensionMock.FakeExtensionMock

// func InitializeFakeInfrastructure() infrastructure.ServiceInfrastructure {
// 	var r datastore.Repository = &fakeRepo
// 	var e extension.Extension = &fakeExtension

// 	type InfrastructureMock struct {
// 		datastore.Repository
// 		extension.Extension
// 	}
// 	infra := func() infrastructure.ServiceInfrastructure {
// 		return &InfrastructureMock{r, e}
// 	}()

// 	return infra
// }

// // InitializeFakeShortCodeService represents a fake shortcode service interactor
// func InitializeFakeShortCodeService() (usecases.Usecases, error) {

// 	infra := InitializeFakeInfrastructure()

// 	i := usecases.NewUseCasesInteractor(infra)

// 	return i, nil
// }
