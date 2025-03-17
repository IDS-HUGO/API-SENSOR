package repository

import "go-hexagonal-api/src/MQTT/domain/entities"

type DataRepository interface {
	SaveData(data *entities.Data) error
}
