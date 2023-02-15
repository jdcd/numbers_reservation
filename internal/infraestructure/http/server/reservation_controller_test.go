package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jdcd/numbers_reservation/internal"
	"github.com/jdcd/numbers_reservation/internal/application"
	"github.com/jdcd/numbers_reservation/internal/domain"
	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
	"github.com/jdcd/numbers_reservation/internal/domain/mock"
	"github.com/jdcd/numbers_reservation/internal/domain/ports"
	"github.com/jdcd/numbers_reservation/internal/infraestructure/http/server"
)

func Test_WhenPostReservationRequestIsOkThenControllerShouldReturnSuccess201(t *testing.T) {
	repositoryMock := &mock.ReservationRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	expectedStruct := domain.Reservation{ClientId: "cheems", Number: 5}
	body := bytes.NewReader([]byte(`{"client_id": "cheems", "number": 5}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/v1/reservation", body)
	var responseStruck domain.Reservation
	repositoryMock.On("MakeReservation", expectedStruct).Return(nil).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, expectedStruct, responseStruck)
	repositoryMock.AssertExpectations(t)

}

func Test_WhenPostReservationRequestIsInvalidThenControllerShouldReturn400(t *testing.T) {
	router := getMockedRouter(nil)
	expectedStruct := api_error.ApiError{
		Code:         http.StatusBadRequest,
		ErrorMessage: "invalid request",
		Message:      "error reading your reservation",
	}
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/v1/reservation", nil)
	var responseStruck api_error.ApiError

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, expectedStruct, responseStruck)

}

func Test_WhenPostReservationRequestHasInvalidFormatThenControllerShouldReturn400(t *testing.T) {
	router := getMockedRouter(nil)
	expectedStruct := api_error.ApiError{
		Code:         http.StatusBadRequest,
		ErrorMessage: "the field clientID cannot be empty",
		Message:      "invalid clientID format",
	}
	body := bytes.NewReader([]byte(`{"client_id_xxx": "cheems", "number": 5}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/v1/reservation", body)
	var responseStruck api_error.ApiError

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, expectedStruct, responseStruck)

}

func Test_WhenApplicationCrashThenPostControllerShouldReturn500(t *testing.T) {
	repositoryMock := &mock.ReservationRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	errorReason := "storage not available"
	errorExpected := errors.New(errorReason)
	requestStruct := domain.Reservation{ClientId: "cheems", Number: 5}
	expectedStruct := api_error.ApiError{
		Code:         http.StatusInternalServerError,
		ErrorMessage: errorReason,
		Message:      "internal server error happens when try to process your request",
	}
	body := bytes.NewReader([]byte(`{"client_id": "cheems", "number": 5}`))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/v1/reservation", body)
	var responseStruck api_error.ApiError
	repositoryMock.On("MakeReservation", requestStruct).Return(errorExpected).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, expectedStruct, responseStruck)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	repositoryMock.AssertExpectations(t)

}

func Test_WhenApplicationCrashThenGetAllReservationControllerShouldReturn500(t *testing.T) {
	repositoryMock := &mock.ReservationRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	errorReason := "storage not available"
	errorExpected := errors.New(errorReason)
	expectedStruct := api_error.ApiError{
		Code:         http.StatusInternalServerError,
		ErrorMessage: errorReason,
		Message:      "internal server error happens when try to process your request",
	}
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/reservation", nil)
	var responseStruck api_error.ApiError
	repositoryMock.On("GetAllReservations").Return([]domain.Reservation{}, errorExpected).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, expectedStruct, responseStruck)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	repositoryMock.AssertExpectations(t)
}

func Test_WhenApplicationResponseIsOkThenGetAllReservationControllerShouldReturnSuccess200WithData(t *testing.T) {
	repositoryMock := &mock.ReservationRepositoryMock{}
	router := getMockedRouter(repositoryMock)
	expectedResponse := []domain.Reservation{domain.Reservation{ClientId: "cheems", Number: 5}}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/reservation", nil)
	var responseStruck []domain.Reservation
	repositoryMock.On("GetAllReservations").Return(expectedResponse, nil).Once()

	router.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseStruck)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, responseStruck)
	repositoryMock.AssertExpectations(t)
}

func getMockedRouter(repository ports.ReservationRepository) *gin.Engine {
	appMock := &application.ReservationApplication{Repository: repository}
	router := internal.SetupRouter(&internal.RouterDependencies{
		ReservationController: &server.ReservationController{ReservationApp: appMock},
	})
	return router
}
