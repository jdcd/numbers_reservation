package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jdcd/numbers_reservation/internal/infraestructure/http/server"
)

type RouterDependencies struct {
	ReservationController *server.ReservationController
}

func SetupRouter(dependencies *RouterDependencies) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1Reservation := v1.Group("/reservation")
	v1Reservation.POST("", dependencies.ReservationController.PostReservation)

	return router
}