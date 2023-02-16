package mongo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jdcd/numbers_reservation/internal/domain"
	"github.com/jdcd/numbers_reservation/internal/domain/api_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	errorRepositoryIsEmptyMessage    = "no reservation has been made"
	errorRepositoryIsEmptyDetail     = "repository is empty"
	errorFindMessage                 = "error when try to search for reservations"
	genericErrorInsertMessage        = "error when try to save your reservation clientID: %s number: %d"
	duplicateClientIDErrorFlag       = "write exception: write errors: [E11000 duplicate key error collection: numbers_reservation.reservations index: _id_"
	duplicateClientIDErrorMessage    = "each client only has one reserved number."
	duplicateClientIDErrorDetail     = "clientID: %s already has an active reservation"
	duplicateReservationErrorFlag    = "write exception: write errors: [E11000 duplicate key error collection: numbers_reservation.reservations index: reservation_number_1"
	duplicateReservationErrorMessage = "each number only belongs to one client"
	duplicateReservationErrorDetail  = "reservation: %d is already taken"
)

type ReservationRepositoryMongo struct {
	Coll *mongo.Collection
}

func (r *ReservationRepositoryMongo) MakeReservation(reservation domain.Reservation) error {
	_, err := r.Coll.InsertOne(context.TODO(), reservation)
	if err != nil {
		return r.processError(err, reservation)
	}

	//ToDo add Logger
	return nil
}

func (r *ReservationRepositoryMongo) GetAllReservations() ([]domain.Reservation, error) {
	cur, err := r.Coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		formattedError := api_error.CreateFormatError(api_error.ConnectionError, errorFindMessage, err.Error())
		//ToDo add Logger

		return []domain.Reservation{}, errors.New(formattedError)
	}

	var results []domain.Reservation
	if err := cur.Err(); err != nil {
		formattedError := api_error.CreateFormatError(api_error.ConnectionError, errorFindMessage,
			err.Error())
		//ToDo add Logger

		return []domain.Reservation{}, errors.New(formattedError)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Reservation
		err := cur.Decode(&elem)
		if err != nil {
			formattedError := api_error.CreateFormatError(api_error.ConnectionError, errorFindMessage,
				err.Error())
			//ToDo add Logger

			return []domain.Reservation{}, errors.New(formattedError)
		}

		results = append(results, elem)
	}

	if len(results) == 0 {
		formattedError := api_error.CreateFormatError(api_error.DataNotFound, errorRepositoryIsEmptyMessage,
			errorRepositoryIsEmptyDetail)
		//ToDo add Logger

		return []domain.Reservation{}, errors.New(formattedError)
	}

	err = cur.Close(context.TODO())
	if err != nil {
		//Todo Warning
	}

	return results, nil
}

func (r *ReservationRepositoryMongo) processError(err error, reservation domain.Reservation) error {
	if strings.Contains(err.Error(), duplicateClientIDErrorFlag) {
		formattedError := api_error.CreateFormatError(api_error.BusinessRule,
			fmt.Sprintf(duplicateClientIDErrorMessage), fmt.Sprintf(duplicateClientIDErrorDetail, reservation.ClientId))
		//ToDo add Logger

		return errors.New(formattedError)

	} else if strings.Contains(err.Error(), duplicateReservationErrorFlag) {
		formattedError := api_error.CreateFormatError(api_error.BusinessRule, duplicateReservationErrorMessage,
			fmt.Sprintf(duplicateReservationErrorDetail, reservation.Number))
		//ToDo add Logger

		return errors.New(formattedError)

	} else {
		formattedError := api_error.CreateFormatError(api_error.ConnectionError,
			fmt.Sprintf(genericErrorInsertMessage, reservation.ClientId, reservation.Number),
			err.Error())
		//ToDo add Logger
		return errors.New(formattedError)
	}
}
