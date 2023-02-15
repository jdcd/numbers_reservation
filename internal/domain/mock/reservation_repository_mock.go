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
