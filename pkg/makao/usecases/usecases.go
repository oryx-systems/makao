package usecases

import (
	"github.com/oryx-systems/makao/pkg/makao/usecases/otp"
	"github.com/oryx-systems/makao/pkg/makao/usecases/residence"
	"github.com/oryx-systems/makao/pkg/makao/usecases/user"
)

// Makao manages the usecases intrefaces
type Makao struct {
	User      user.UseCasesUser
	OTP       otp.UseCasesOTP
	Residence residence.UsecasesResidence
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewMakaoUsecase(
	user user.UseCasesUser,
	otp otp.UseCasesOTP,
	residence residence.UsecasesResidence,
) *Makao {
	m := &Makao{
		User:      user,
		OTP:       otp,
		Residence: residence,
	}

	return m
}
