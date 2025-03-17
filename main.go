package main

import (
	"go-hexagonal-api/src/MQTT/application/repository"
	"go-hexagonal-api/src/MQTT/application/usecases"
	"go-hexagonal-api/src/MQTT/infrastructure/adapters"
	"go-hexagonal-api/src/MQTT/infrastructure/controllers"
	"go-hexagonal-api/src/MQTT/infrastructure/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo, err := repository.NewMySQLRepository("root:root@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sender, err := adapters.NewMQTTSender("tcp://184.72.149.16:1883", "clientID")
	if err != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", err)
	}

	useCase := usecases.NewSendDataUseCase(repo, sender)

	controller := controllers.NewDataController(useCase)

	router := mux.NewRouter()

	routes.InitRoutes(router, controller)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
