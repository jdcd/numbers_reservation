package domain

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
)

const (
	errorFormatClientID           = "invalid clientID format"
	errorDetailEmptyClientID      = "the field clientID cannot be empty"
	errorDetailClientIDMaxLen     = "the maximum allowed size of the clientID is %d"
	errorFormatNumber             = "invalid number format"
	errorInvalidNumber            = "the available numbers are the range [1-%d]"
	errorNonPositiveInvalidNumber = "the number must be positive"
	defaultMaxClientIDLen         = 128
	defaultMaxNumberValue         = 1000
)

type Reservation struct {
	ClientId string `json:"client_id,omitempty"`
	Number   int    `json:"number,omitempty"`
}

func (r *Reservation) Validate() error {
	if err := r.checkClientID(); err != nil {
		return err
	}

	if err := r.checkNumber(); err != nil {
		return err
	}

	return nil
}

func (r *Reservation) checkClientID() error {
	if r.ClientId == "" {
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatClientID,
			errorDetailEmptyClientID)
		//ToDo add Logger
		return errors.New(formattedError)
	}

	if len(r.ClientId) > r.clientIDMaxLength() {
		lenError := fmt.Sprintf(errorDetailClientIDMaxLen, r.clientIDMaxLength())
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatClientID,
			lenError)
		//ToDo add Logger
		return errors.New(formattedError)
	}

	return nil

}

func (r *Reservation) checkNumber() error {
	if r.Number < 1 {
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
			errorNonPositiveInvalidNumber)
		//ToDo add Logger
		return errors.New(formattedError)
	}

	if r.Number > r.numberMaxValue() {
		rangeError := fmt.Sprintf(errorInvalidNumber, r.numberMaxValue())
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
			rangeError)
		//ToDo add Logger
		return errors.New(formattedError)
	}

	return nil
}

func (r *Reservation) clientIDMaxLength() int {
	maxLen := defaultMaxClientIDLen
	envMaxLen, err := strconv.Atoi(os.Getenv("CLIENT_ID_MAX_LEN"))
	if err == nil {
		//ToDo add Logger
		maxLen = envMaxLen
	}
	return maxLen
}

func (r *Reservation) numberMaxValue() int {
	maxValue := defaultMaxNumberValue
	intMaxValue, err := strconv.Atoi(os.Getenv("NUMBER_MAX_VALUE"))
	if err == nil {
		//ToDo add Logger
		maxValue = intMaxValue
	}
	return maxValue
}
