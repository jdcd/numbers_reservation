package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jdcd/numbers_reservation/internal/application"
	"github.com/jdcd/numbers_reservation/internal/domain"
	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
	"github.com/jdcd/numbers_reservation/pkg"
	"net/http"
)

const (
	encodingJsonError        = "error reading your reservation"
	successfulCreationLogger = "clientId: %s with reservation %d creates successfully \n"
)

type ReservationController struct {
	ReservationApp application.ReservationApp
}

func (r ReservationController) PostReservation(c *gin.Context) {
	var newReservation domain.Reservation
	if err := c.BindJSON(&newReservation); err != nil {
		formattedError := api_error.CreateFormatError(api_error.DataValidation, encodingJsonError, err.Error())
		apiError := api_error.MapApiError(errors.New(formattedError))
		c.IndentedJSON(apiError.Code, apiError)
		pkg.ErrorLogger().Printf("%s, %s \n", encodingJsonError, err.Error())

		return
	}

	if err := newReservation.Validate(); err != nil {
		apiError := api_error.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	if err := r.ReservationApp.CreateReservation(newReservation); err != nil {
		apiError := api_error.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	pkg.InfoLogger().Printf(successfulCreationLogger, newReservation.ClientId, newReservation.Number)
	c.IndentedJSON(http.StatusCreated, newReservation)
}

func (r ReservationController) GetAllReservation(c *gin.Context) {
	reservations, err := r.ReservationApp.ConsultReservationAll()
	if err != nil {
		apiError := api_error.MapApiError(err)
		c.IndentedJSON(apiError.Code, apiError)
		return
	}

	c.IndentedJSON(http.StatusOK, reservations)
}
