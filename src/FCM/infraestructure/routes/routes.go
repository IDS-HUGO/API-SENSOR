package routes

import (
	"go-hexagonal-api/src/FCM/infraestructure/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, sensorController *controller.SensorController) {
	
	r.POST("/sensor-data", sensorController.HandleSensorData)
	
	// Ruta para suscribirse al tópico de Firebase
	r.POST("/suscribe", sensorController.SuscribeToTopic)
}
