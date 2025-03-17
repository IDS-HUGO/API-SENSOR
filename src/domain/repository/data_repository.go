package repository

import "go-hexagonal-api/src/domain/entities"

type DataRepository interface {
    SaveData(data *entities.Data) error
}
