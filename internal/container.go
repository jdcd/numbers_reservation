package internal

import (
	"github.com/jdcd/numbers_reservation/internal/application"
	"github.com/jdcd/numbers_reservation/internal/domain/ports"
	"github.com/jdcd/numbers_reservation/internal/infraestructure/adapter/mongo"
	"github.com/jdcd/numbers_reservation/internal/infraestructure/http/server"
)

func GetRouterDependencies() *RouterDependencies {
	return &RouterDependencies{
		ReservationController: getReservationController(),
	}
}

func getReservationController() *server.ReservationController {
	return &server.ReservationController{
		ReservationApp: getReservationApp(),
	}
}

func getReservationApp() application.ReservationApp {
	return &application.ReservationApplication{
		Repository: getReservationRepository(),
	}
}

func getReservationRepository() ports.ReservationRepository {
	return &mongo.ReservationRepositoryMongo{
		Coll: mongo.GetCollection(),
	}
}
