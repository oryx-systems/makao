package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/application/enums"
	"github.com/oryx-systems/makao/pkg/makao/application/extension"
	"github.com/oryx-systems/makao/pkg/makao/application/utils"
	"github.com/oryx-systems/makao/pkg/makao/domain"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
	"github.com/sirupsen/logrus"
)

const (
	appName = "Makao"
)

// UseCasesOTP contain all the method required for OTP delivery
type UseCasesOTP interface {
	GenerateAndSendOTP(ctx context.Context, phoneNumber string, flavour enums.Flavour) (string, error)
}

// UseCasesOTPImpl represents the user otp usecase implementation
type UseCasesOTPImpl struct {
	infrastructure infrastructure.Datastore
	Ext            extension.Extension
}

// NewUseCaseOTP initializes the new otp implementation
func NewUseCaseOTP(infra infrastructure.Datastore) UseCasesOTP {
	ext := extension.NewExtension()
	return &UseCasesOTPImpl{
		infrastructure: infra,
		Ext:            ext,
	}
}

// GenerateAndSendOTP generates and sends an OTP to the user
func (o *UseCasesOTPImpl) GenerateAndSendOTP(ctx context.Context, phoneNumber string, flavour enums.Flavour) (string, error) {
	// Logged in user
	uid, err := o.Ext.GetLoggedInUserUID(ctx)
	if err != nil {
		return "", err
	}

	logrus.Print("THE LOGGED IN USER ID IS: ", uid)

	validatePhoneNumber, err := helpers.NormalizeMSISDN(phoneNumber)
	if err != nil {
		return "", err
	}

	// TODO: Implement get user profile by phone number

	if !flavour.IsValid() {
		return "", fmt.Errorf("invalid flavour")
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return "", fmt.Errorf("failed to generate an OTP")
	}

	var message string
	switch flavour {
	case enums.FlavourConsumer:
		message = fmt.Sprintf("Your %v verification code is %s", appName, otp)
	case enums.FlavourPro:
		message = fmt.Sprintf("Your %v verification code is %s", appName, otp)
	}

	// TODO: 1. Implement send sms logic here
	logrus.Print("OTP MESSAGE: ", message)

	otpData := &domain.OTP{
		IsValid:     true,
		ValidUntil:  time.Now().Add(time.Minute * 5),
		PhoneNumber: *validatePhoneNumber,
		OTP:         otp,
		Flavour:     flavour,
		Medium:      "SMS",
		UserID:      "ad2ba725-dd9b-4e1b-ae5d-b10d5d852ff4", // TODO: Get the user id from the user profile from 1 above
	}

	// Save the OTP to the database
	err = o.infrastructure.Create.SaveOTP(ctx, otpData)
	if err != nil {
		return "", err
	}

	return otp, nil
}
