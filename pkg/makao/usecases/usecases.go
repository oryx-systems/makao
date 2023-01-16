package usecases

import (
	"github.com/oryx-systems/makao/pkg/makao/usecases/house"
	"github.com/oryx-systems/makao/pkg/makao/usecases/otp"
	"github.com/oryx-systems/makao/pkg/makao/usecases/residence"
	"github.com/oryx-systems/makao/pkg/makao/usecases/user"
)

// Makao manages the usecases intrefaces
type Makao struct {
	User      user.UseCasesUser
	OTP       otp.UseCasesOTP
	Residence residence.UsecasesResidence
	House     house.UsecaseHouse
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewMakaoUsecase(
	user user.UseCasesUser,
	otp otp.UseCasesOTP,
	residence residence.UsecasesResidence,
	house house.UsecaseHouse,
) *Makao {
	m := &Makao{
		User:      user,
		OTP:       otp,
		Residence: residence,
		House:     house,
	}

	return m
}
