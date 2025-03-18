package controller

import (
	"bytes"
	"encoding/json"
	"go-hexagonal-api/src/FCM/application/useCases"
	"go-hexagonal-api/src/FCM/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SensorController struct {
	Service *useCases.SensorService
}

// Manejar los datos de los sensores
func (sc *SensorController) HandleSensorData(c *gin.Context) {
	var data entities.SensorData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	// Aquí procesarías los datos recibidos, tal vez usando `sc.Service` si es necesario
	if err := sc.Service.ProcessSensorData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando datos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos de sensores procesados"})
}

// Función para suscribirse al tópico de Firebase
func (sc *SensorController) SuscribeToTopic(c *gin.Context) {
	var request struct {
		Token string `json:"token"`
		Topic string `json:"topic"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	// Obtener el token de servidor de Firebase
	fcmServerKey := os.Getenv("FIREBASE_SERVER_KEY") // Debes configurar tu clave del servidor en una variable de entorno

	// Construir la solicitud de suscripción
	url := "https://iid.googleapis.com/iid/v1:batchAdd"
	body := map[string]interface{}{
		"to": "/topics/" + request.Topic,
		"registration_tokens": []string{request.Token},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar los datos"})
		return
	}

	// Hacer la solicitud POST a FCM
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
		return
	}

	// Configurar los encabezados de la solicitud
	req.Header.Set("Authorization", "key="+fcmServerKey)
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con FCM"})
		return
	}
	defer resp.Body.Close()

	// Procesar la respuesta
	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"message": "Suscripción exitosa"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al suscribirse al tema"})
	}
}
