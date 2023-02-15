package mock

import (
	"github.com/jdcd/numbers_reservation/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ReservationRepositoryMock struct {
	mock.Mock
}

func (m *ReservationRepositoryMock) MakeReservation(reservation domain.Reservation) error {
	args := m.Called(reservation)
	return args.Error(0)
}

func (m *ReservationRepositoryMock) GetAllReservations() ([]domain.Reservation, error) {
	args := m.Called()
	return args.Get(0).([]domain.Reservation), args.Error(1)
}
