package domain

import (
	"errors"
	"fmt"
	"github.com/jdcd/numbers_reservation/pkg"
	"os"
	"strconv"

	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
)

const (
	errorFormatClientID                  = "invalid clientID format"
	errorDetailEmptyClientID             = "the field clientID cannot be empty"
	errorDetailClientIDMaxLen            = "the maximum allowed size of the clientID is %d"
	loggerDetailClientIDMaxLen           = "cannot store clientID %s %s \n"
	errorFormatNumber                    = "invalid number format"
	errorInvalidValueNumber              = "the available numbers are the range [1-%d]"
	loggerInvalidValueNumber             = "%d is out of range, %s \n"
	errorDetailNonPositiveInvalidNumber  = "the number must be positive"
	loggerDetailNonPositiveInvalidNumber = "the number %d must be positive \n"
	errorDetailZeroOrNotFoundNumber      = "the field number cannot be empty or 0"
	defaultMaxClientIDLen                = 128
	defaultMaxNumberValue                = 1000
)

type Reservation struct {
	ClientId string `bson:"_id"  json:"client_id,omitempty"`
	Number   int    `bson:"reservation_number" json:"number,omitempty"`
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
		pkg.ErrorLogger().Println(errorDetailEmptyClientID)
		return errors.New(formattedError)
	}

	if len(r.ClientId) > r.clientIDMaxLength() {
		lenError := fmt.Sprintf(errorDetailClientIDMaxLen, r.clientIDMaxLength())
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatClientID,
			lenError)
		pkg.ErrorLogger().Printf(loggerDetailClientIDMaxLen, r.ClientId, lenError)
		return errors.New(formattedError)
	}

	return nil

}

func (r *Reservation) checkNumber() error {
	if r.Number == 0 {
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
			errorDetailZeroOrNotFoundNumber)
		pkg.ErrorLogger().Println(errorDetailZeroOrNotFoundNumber)
		return errors.New(formattedError)
	}

	if r.Number < 0 {
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
			errorDetailNonPositiveInvalidNumber)
		pkg.ErrorLogger().Printf(loggerDetailNonPositiveInvalidNumber, r.Number)
		return errors.New(formattedError)
	}

	if r.Number > r.numberMaxValue() {
		rangeError := fmt.Sprintf(errorInvalidValueNumber, r.numberMaxValue())
		formattedError := api_error.CreateFormatError(api_error.DataValidation, errorFormatNumber,
			rangeError)
		pkg.ErrorLogger().Printf(loggerInvalidValueNumber, r.Number, rangeError)
		return errors.New(formattedError)
	}

	return nil
}

func (r *Reservation) clientIDMaxLength() int {
	maxLen := defaultMaxClientIDLen
	envMaxLen, err := strconv.Atoi(os.Getenv("CLIENT_ID_MAX_LEN"))
	if err == nil {
		maxLen = envMaxLen
	}
	return maxLen
}

func (r *Reservation) numberMaxValue() int {
	maxValue := defaultMaxNumberValue
	intMaxValue, err := strconv.Atoi(os.Getenv("NUMBER_MAX_VALUE"))
	if err == nil {
		maxValue = intMaxValue
	}
	return maxValue
}
