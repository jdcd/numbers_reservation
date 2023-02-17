package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"

	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
)

func TestWhenLenOfClientIDIsBiggerThanMaxLimitClientIDThenValidateShouldReturnError(t *testing.T) {
	t.Setenv("CLIENT_ID_MAX_LEN", "3")
	detailedError := fmt.Sprintf(errorDetailClientIDMaxLen, 3)
	fError := api_error.CreateFormatError(api_error.DataValidation, errorFormatClientID, detailedError)
	r := &Reservation{
		ClientId: "four",
		Number:   21,
	}
	expectedError := errors.New(fError)

	err := r.Validate()

	assert.Equal(t, expectedError, err)

}

// TestValidateInvalidNumber
func TestWhenNumberIsBiggerThanMaxNumValueThenValidShouldReturnError(t *testing.T) {
	t.Setenv("NUMBER_MAX_VALUE", "80")
	detailedError := fmt.Sprintf(errorInvalidValueNumber, 80)
	fError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber, detailedError)
	expectedError := errors.New(fError)
	r := &Reservation{ClientId: "four", Number: 81}

	err := r.Validate()

	assert.Equal(t, expectedError, err)

}

func TestWhenReservationIsFineThenValidShouldReturnNilError(t *testing.T) {
	var expectedError error = nil
	r := &Reservation{ClientId: "cheems", Number: 5}

	err := r.Validate()

	assert.Equal(t, expectedError, err)
}

func TestWhenNumberIsNotPositiveThenValidShouldReturnError(t *testing.T) {
	formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
		errorDetailNonPositiveInvalidNumber)
	expectedError := errors.New(formattedError)
	r := &Reservation{ClientId: "cheems", Number: -8}

	err := r.Validate()

	assert.Equal(t, expectedError, err)
}

func TestWhenClientIDIsEmptyThenValidShouldReturnError(t *testing.T) {
	r := &Reservation{Number: -8}
	formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatClientID,
		errorDetailEmptyClientID)
	expectedError := errors.New(formattedError)

	err := r.Validate()

	assert.Equal(t, expectedError, err)
}

func TestWhenNumberIsZeroOrNotFoundThenValidShouldReturnError(t *testing.T) {
	formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
		errorDetailZeroOrNotFoundNumber)
	expectedError := errors.New(formattedError)
	r := &Reservation{ClientId: "cheems"}

	err := r.Validate()

	assert.Equal(t, expectedError, err)
}
