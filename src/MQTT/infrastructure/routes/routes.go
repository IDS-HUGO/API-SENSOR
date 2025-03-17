package routes

import (
	"go-hexagonal-api/src/MQTT/infrastructure/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, controller *controllers.DataController) {
	router.HandleFunc("/data", controller.HandlePostData).Methods("POST")
}
