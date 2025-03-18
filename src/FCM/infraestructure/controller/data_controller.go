package controller

import (
	"go-hexagonal-api/src/FCM/application/useCases"
	"go-hexagonal-api/src/FCM/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SensorController struct {
	Service *useCases.SensorService
}

func NewSensorController(service *useCases.SensorService) *SensorController {
	return &SensorController{
		Service: service,
	}
}

func (sc *SensorController) HandleSensorData(c *gin.Context) {
	var data entities.SensorData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	if err := sc.Service.ProcessSensorData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando datos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos de sensores procesados"})
}