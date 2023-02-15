package application

import (
	"github.com/jdcd/numbers_reservation/internal/domain"
	"github.com/jdcd/numbers_reservation/internal/domain/ports"
)

type ReservationApp interface {
	CreateReservation(reservation domain.Reservation) error
	ConsultReservationAll() ([]domain.Reservation, error)
}

type ReservationApplication struct {
	Repository ports.ReservationRepository
}

func (r *ReservationApplication) CreateReservation(reservation domain.Reservation) error {
	return r.Repository.MakeReservation(reservation)
}

func (r *ReservationApplication) ConsultReservationAll() ([]domain.Reservation, error) {
	return r.Repository.GetAllReservations()
}
