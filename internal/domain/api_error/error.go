package api_error

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ErrorType string

const (
	DataNotFound    ErrorType = "dataNotFound"
	ConnectionError           = "connectionError"
	DataValidation            = "dataValidation"
	ThirdPart                 = "thirdPart"
)

const (
	unknownErrorDetail               = "unknown error"
	genericInternalServerErrorDetail = "internal server error happens when try to process your request"
)

type ApiError struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_detail"`
	Message      string `json:"message"`
}

func CreateFormatError(errorType ErrorType, message string, errDetail string) string {
	return fmt.Sprintf("%s|%s|%s", errorType, message, errDetail)
}

func MapApiError(err error) ApiError {
	if err == nil {
		return getStandardApiError(errors.New(unknownErrorDetail))
	}

	parts := strings.Split(err.Error(), "|")
	if len(parts) != 3 {
		return getStandardApiError(err)
	}

	errorType := ErrorType(parts[0])
	errorMessage := parts[1]
	errorDetail := parts[2]

	switch errorType {
	case ThirdPart, ConnectionError, DataValidation, DataNotFound:
		return ApiError{
			Code:         selectStatusCode(errorType),
			Message:      errorMessage,
			ErrorMessage: errorDetail,
		}
	default:
		return getStandardApiError(errors.New(errorDetail))
	}
}

func selectStatusCode(errorType ErrorType) int {
	switch errorType {
	case ThirdPart, ConnectionError:
		return http.StatusInternalServerError
	case DataValidation:
		return http.StatusBadRequest
	case DataNotFound:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

func getStandardApiError(err error) ApiError {
	return ApiError{
		Code:         http.StatusInternalServerError,
		Message:      genericInternalServerErrorDetail,
		ErrorMessage: err.Error(),
	}
}
