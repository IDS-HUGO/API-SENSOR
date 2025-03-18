package adapters

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseService struct{}

func (f *FirebaseService) SendNotification(deviceToken, message string) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("airsafetech-firebase-adminsdk-fbsvc-a7ffb511db.json")
	conf := &firebase.Config{ProjectID: "airsafetech"}

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return fmt.Errorf("error inicializando Firebase: %v", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error obteniendo el cliente Messaging: %v", err)
	}

	msg := &messaging.Message{
		Token: deviceToken,
		Notification: &messaging.Notification{
			Title: "Alerta de Sensor",
			Body:  message,
		},
	}

	_, err = client.Send(ctx, msg)
	if err != nil {
		return fmt.Errorf("error enviando notificación: %v", err)
	}
	return nil
}