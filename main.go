package main

import (
	"log"

	"go-hexagonal-api/src/FCM/application/useCases"
	"go-hexagonal-api/src/FCM/infraestructure/controller"
	fcm "go-hexagonal-api/src/FCM/infraestructure/routes"
	"go-hexagonal-api/src/MQTT/application/repository"
	"go-hexagonal-api/src/MQTT/application/usecases"
	"go-hexagonal-api/src/MQTT/infrastructure/adapters"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo, err := repository.NewMySQLRepository("root:root@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	sender, err := adapters.NewMQTTSender("tcp://184.72.149.16:1883", "clientID")
	if err != nil {
		log.Fatalf("Error al conectar con el servidor MQTT: %v", err)
	}

	useCase := usecases.NewSendDataUseCase(repo, sender)
	controller := controller.NewSensorController(&useCases.SensorService{})

	fcm.SetupRouter(r, controller)

	log.Println("Servidor ejecutándose en http://localhost:3000")
	log.Fatal(r.Run(":3000"))
}