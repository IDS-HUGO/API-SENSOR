package infrastructure

import (
	"go-hexagonal-api/core"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTSender struct {
	Client mqtt.Client
}

func NewMQTTSender(broker string, clientID string) (*MQTTSender, error) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MQTTSender{Client: client}, nil
}

func (s *MQTTSender) SendMessage(data *core.Data) error {
	token := s.Client.Publish("data/topic", 0, false, data.Message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	log.Printf("Sent message: %s", data.Message)
	return nil
}
