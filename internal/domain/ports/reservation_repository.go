package ports

import "github.com/jdcd/numbers_reservation/internal/domain"

type ReservationRepository interface {
	MakeReservation(reservation domain.Reservation) error
}
