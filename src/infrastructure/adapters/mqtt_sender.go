package adapters

import (
    "go-hexagonal-api/src/domain/entities"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "log"
)

// MQTTSender es el adaptador que envía mensajes a un broker MQTT.
type MQTTSender struct {
    Client mqtt.Client
}

// NewMQTTSender crea una nueva instancia de MQTTSender.
func NewMQTTSender(broker string, clientID string) (*MQTTSender, error) {
    opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }

    return &MQTTSender{Client: client}, nil
}

// SendMessage envía un mensaje a través de MQTT.
func (s *MQTTSender) SendMessage(data *entities.Data) error {
    token := s.Client.Publish("data/topic", 0, false, data.Message)
    token.Wait()
    if token.Error() != nil {
        return token.Error()
    }
    log.Printf("Sent message: %s", data.Message)
    return nil
}