package repository

import "go-hexagonal-api/src/FCM/domain/entities"

type DataRepository interface {
	ProcessSensorData(data entities.SensorData) error
	RegisterToken(token string) error
	SendNotification(deviceToken, message string) error
	GetDeviceTokens() []string
}