package api_error

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestWhenErrorIsNilThenMapApiErrorShouldReturnStandardError(t *testing.T) {
	expectedApiError := getStandardApiError(errors.New(unknownErrorDetail))

	apiErr := MapApiError(nil)

	assert.Equal(t, expectedApiError, apiErr)
}

func TestWhenErrorIsDataValidationTypeThenMapApiErrorShouldReturnApiErrorFormatted(t *testing.T) {
	errorMessage := "invalid format"
	errorDetail := "invalid len for value"
	formattedError := CreateFormatError(DataValidation, errorMessage, errorDetail)
	expectedApiError := &ApiError{
		Code:         http.StatusBadRequest,
		ErrorMessage: errorDetail,
		Message:      errorMessage,
	}

	apiError := MapApiError(errors.New(formattedError))

	assert.Equal(t, expectedApiError, apiError)
}

func TestWhenErrorIsDataNotFoundTypeThenMapApiErrorShouldReturnApiErrorFormatted(t *testing.T) {
	errorMessage := "empty repository"
	errorDetail := "results not found"
	formattedError := CreateFormatError(DataNotFound, errorMessage, errorDetail)
	expectedApiError := &ApiError{
		Code:         http.StatusOK,
		ErrorMessage: errorDetail,
		Message:      errorMessage,
	}

	apiError := MapApiError(errors.New(formattedError))

	assert.Equal(t, expectedApiError, apiError)
}

func TestWhenErrorIsThirdPartTypeThenMapApiErrorShouldReturnApiErrorFormatted(t *testing.T) {
	errorMessage := "error with repository"
	errorDetail := "connection refused, we don't trust in you identity"
	formattedError := CreateFormatError(ThirdPart, errorMessage, errorDetail)
	expectedApiError := &ApiError{
		Code:         http.StatusInternalServerError,
		ErrorMessage: errorDetail,
		Message:      errorMessage,
	}

	apiError := MapApiError(errors.New(formattedError))

	assert.Equal(t, expectedApiError, apiError)
}

func TestWhenErrorIsUnknownTypeThenMapApiErrorShouldReturnApiErrorFormatted(t *testing.T) {
	errorMessage := genericInternalServerErrorDetail
	errorDetail := "error with the internal processor"
	formattedError := CreateFormatError("new", errorMessage, errorDetail)
	expectedApiError := getStandardApiError(errors.New(errorDetail))

	apiError := MapApiError(errors.New(formattedError))

	assert.Equal(t, expectedApiError, apiError)
}

func TestWhenErrorDoesNotHaveFormatThenMapApiErrorShouldReturnStandardApiError(t *testing.T) {
	unFormattedError := errors.New(unknownErrorDetail)
	expectedApiError := getStandardApiError(unFormattedError)

	apiErr := MapApiError(unFormattedError)

	assert.Equal(t, expectedApiError, apiErr)
}
