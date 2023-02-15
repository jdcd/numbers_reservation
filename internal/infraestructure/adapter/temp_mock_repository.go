package adapter

import (
	"github.com/jdcd/numbers_reservation/internal/domain"
)

type TempHappyMock struct{}

func (r *TempHappyMock) MakeReservation(reservation domain.Reservation) error {
	return nil
}

func (r *TempHappyMock) GetAllReservations() ([]domain.Reservation, error) {
	var results []domain.Reservation
	cheems := domain.Reservation{ClientId: "cheems", Number: 5}
	doge := domain.Reservation{ClientId: "doge", Number: 2}
	results = append(results, cheems, doge)

	return results, nil
}
