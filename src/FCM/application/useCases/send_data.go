package useCases

import (
	"fmt"
	"log"
	"go-hexagonal-api/src/FCM/domain/entities"
	"go-hexagonal-api/src/FCM/domain/repository"
)

const (
	HighGasThreshold       = 0.5
	LowAirQualityThreshold = 12.0
	HighTemperatureBMP     = 34.0
	HighTemperatureDHT     = 34.0
	LowHumidityThreshold   = 40.0
)

type SensorService struct {
	Repo repository.DataRepository
}

func (s *SensorService) ProcessSensorData(data entities.SensorData) error {
	if data.GasInflamable > HighGasThreshold {
		message := fmt.Sprintf("Alerta: Gas inflamable alto en sensor %s: %.2f", data.DeviceID, data.GasInflamable)
		s.notifyAll(message)
	}

	if data.CalidadAire < LowAirQualityThreshold {
		message := fmt.Sprintf("Alerta: Baja calidad del aire en sensor %s: %.2f", data.DeviceID, data.CalidadAire)
		s.notifyAll(message)
	}

	if data.TemperaturaBMP > HighTemperatureBMP || data.TemperaturaDHT > HighTemperatureDHT {
		message := fmt.Sprintf("Alerta: Alta temperatura en sensor %s: BMP: %.2f°C, DHT: %.2f°C", data.DeviceID, data.TemperaturaBMP, data.TemperaturaDHT)
		s.notifyAll(message)
	}

	if data.Humedad < LowHumidityThreshold {
		message := fmt.Sprintf("Alerta: Humedad baja en sensor %s: %.2f%%", data.DeviceID, data.Humedad)
		s.notifyAll(message)
	}

	return nil
}

func (s *SensorService) notifyAll(message string) {
	tokens := s.Repo.GetDeviceTokens()
	for _, token := range tokens {
		err := s.Repo.SendNotification(token, message)
		if err != nil {
			log.Printf("Error enviando notificación a %s: %v", token, err)
		}
	}
}