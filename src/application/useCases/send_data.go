package usecases

import (
    "go-hexagonal-api/src/domain/entities"
    "go-hexagonal-api/src/domain/repository"
    "go-hexagonal-api/src/infrastructure/adapters"
)

type SendDataUseCase struct {
    Repo   repository.DataRepository
    Sender *adapters.MQTTSender
}

func NewSendDataUseCase(repo repository.DataRepository, sender *adapters.MQTTSender) *SendDataUseCase {
    return &SendDataUseCase{
        Repo:   repo,
        Sender: sender,
    }
}

func (uc *SendDataUseCase) Execute(data *entities.Data) error {
    if err := uc.Repo.SaveData(data); err != nil {
        return err
    }

    if err := uc.Sender.SendMessage(data); err != nil {
        return err
    }

    return nil
}