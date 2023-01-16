package gorm_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
)

func TestPGInstance_RegisterUser(t *testing.T) {
	invalidID := "invalid"

	type args struct {
		ctx        context.Context
		user       *gorm.User
		contact    *gorm.Contact
		identifier *gorm.Identifier
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: Successfully register user",
			args: args{
				ctx: context.Background(),
				user: &gorm.User{
					ID:          &userID,
					FirstName:   gofakeit.FirstName(),
					MiddleName:  gofakeit.BeerAlcohol(),
					LastName:    gofakeit.LastName(),
					Active:      true,
					Flavour:     enums.FlavourPro,
					UserName:    gofakeit.BeerHop(),
					UserType:    "TENANT",
					DeviceToken: gofakeit.UUID(),
				},
				contact: &gorm.Contact{
					ID:           uuid.New().String(),
					Active:       true,
					ContactType:  "PHONE",
					ContactValue: testPhone,
					Flavour:      enums.FlavourConsumer,
					UserID:       &userID,
				},
				identifier: &gorm.Identifier{
					ID:              uuid.New().String(),
					Active:          true,
					IdentifierType:  "NATIONAL_ID",
					IdentifierValue: "1234567893",
					UserID:          &userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to register user",
			args: args{
				ctx: context.Background(),
				user: &gorm.User{
					ID:          &userID,
					FirstName:   gofakeit.FirstName(),
					MiddleName:  gofakeit.BeerAlcohol(),
					LastName:    gofakeit.LastName(),
					Active:      true,
					Flavour:     enums.FlavourPro,
					UserName:    gofakeit.BeerHop(),
					UserType:    "TENANT",
					DeviceToken: gofakeit.UUID(),
				},
				contact: &gorm.Contact{
					ID:           uuid.New().String(),
					Active:       true,
					ContactType:  "PHONE",
					ContactValue: testPhone,
					Flavour:      enums.FlavourConsumer,
					UserID:       &invalidID,
				},
				identifier: &gorm.Identifier{
					ID:              uuid.New().String(),
					Active:          true,
					IdentifierType:  "NATIONAL_ID",
					IdentifierValue: "1234567893",
					UserID:          &userID,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testingDB.RegisterUser(tt.args.ctx, tt.args.user, tt.args.contact, tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
