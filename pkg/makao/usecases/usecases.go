package usecases

import (
	"github.com/oryx-systems/makao/pkg/makao/usecases/otp"
	"github.com/oryx-systems/makao/pkg/makao/usecases/user"
)

// Makao manages the usecases intrefaces
type Makao struct {
	User user.UseCasesUser
	OTP  otp.UseCasesOTP
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewMakaoUsecase(
	user user.UseCasesUser,
	otp otp.UseCasesOTP,
) *Makao {

	m := &Makao{
		User: user,
		OTP:  otp,
	}

	return m
}
