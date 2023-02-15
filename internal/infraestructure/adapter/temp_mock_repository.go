package adapter

import (
	"github.com/jdcd/numbers_reservation/internal/domain"
)

type TempHappyMock struct{}

func (r *TempHappyMock) MakeReservation(reservation domain.Reservation) error {
	return nil
}
