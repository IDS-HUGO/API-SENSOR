package usecases

import (
    "go-hexagonal-api/src/domain/entities"
    "go-hexagonal-api/src/domain/repository"
    "go-hexagonal-api/src/infrastructure/adapters"
)

// SendDataUseCase es el caso de uso que maneja el flujo de guardar datos y enviar mensajes.
type SendDataUseCase struct {
    Repo   repository.DataRepository
    Sender *adapters.MQTTSender  // Cambiar a puntero
}

// NewSendDataUseCase crea un nuevo caso de uso.
func NewSendDataUseCase(repo repository.DataRepository, sender *adapters.MQTTSender) *SendDataUseCase {
    return &SendDataUseCase{
        Repo:   repo,
        Sender: sender,  // Acepta puntero
    }
}

// Execute guarda los datos en el repositorio y envía el mensaje.
func (uc *SendDataUseCase) Execute(data *entities.Data) error {
    // Guardar datos en el repositorio
    if err := uc.Repo.SaveData(data); err != nil {
        return err
    }

    // Enviar mensaje a través de MQTT
    if err := uc.Sender.SendMessage(data); err != nil {
        return err
    }

    return nil
}