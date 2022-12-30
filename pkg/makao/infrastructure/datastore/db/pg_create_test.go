package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	gormMock "github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm/mock"
)

func TestDbServiceImpl_RegisterUser(t *testing.T) {
	var fakeGorm = gormMock.NewGormDatastoreMock()
	d := NewDBService(fakeGorm, fakeGorm, fakeGorm)

	type args struct {
		ctx        context.Context
		user       *domain.User
		contact    *domain.Contact
		identifier *domain.Identifier
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: Register user",
			args: args{
				ctx: context.Background(),
				user: &domain.User{
					ID:         uuid.New().String(),
					FirstName:  gofakeit.BeerAlcohol(),
					MiddleName: gofakeit.BeerAlcohol(),
					LastName:   gofakeit.BeerAlcohol(),
					Active:     true,
					Flavour:    enums.FlavourPro,
					UserName:   gofakeit.BeerAlcohol(),
					UserType:   "TEST",
					UserIdentifier: domain.Identifier{
						ID:              uuid.New().String(),
						Active:          true,
						IdentifierType:  "NATIONAL_ID",
						IdentifierValue: "123456789",
						UserID:          uuid.New().String(),
					},
					UserContact:     domain.Contact{},
					DeviceToken:     "",
					Residence:       "",
					AuthCredentials: domain.AuthCredentials{},
				},
				contact:    &domain.Contact{},
				identifier: &domain.Identifier{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.RegisterUser(tt.args.ctx, tt.args.user, tt.args.contact, tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("DbServiceImpl.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
