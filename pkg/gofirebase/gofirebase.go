package gofirebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type service struct {
	App             *firebase.App
	DBClient        *db.Client
	MessageClient   *messaging.Client
	FirestoreClient *firestore.Client
}

type Config struct {
	FirebaseJsonPath string
	ProjectID        string
	DatabaseURL      string
	IsInitMsg        bool
	IsInitDatabase   bool
}

type FirebaseIntf interface {
	Setup(config Config)
	SendMessage(token string, msg string) error
}

func NewFirebaseService(config Config) FirebaseIntf {
	svc := service{}
	svc.Setup(config)
	return &svc
}

func (s *service) Setup(sconfig Config) {
	opt := option.WithCredentialsFile(sconfig.FirebaseJsonPath)
	config := &firebase.Config{ProjectID: sconfig.ProjectID, DatabaseURL: sconfig.DatabaseURL}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf(" (1) Failed to initialize Firebase app: %v", err)
	}
	s.App = app

	ctx := context.Background()

	if sconfig.IsInitDatabase {
		db, err := app.Database(ctx)
		if err != nil {
			log.Fatalf("(2) Failed to access database: %v", err)
		}
		s.DBClient = db
	}

	if sconfig.IsInitMsg {
		messenger, err := app.Messaging(context.Background())
		if err != nil {
			log.Fatalf("(3) Failed to access messaging: %v", err)
		}
		s.MessageClient = messenger
	}

}

func (s *service) SendMessage(token string, msg string) error {
	log.Println("SendMessage() invoked ...")

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Device Notification",
			Body:  msg,
		},
		Token: token, // Replace with the device token of the target Android device
	}

	data, err := s.MessageClient.Send(context.Background(), message)
	if err != nil {
		log.Printf("(4) Failed to send message: %v", err)
		return err
	}

	log.Println("Successfully Sent Message: ", data)

	return err
}
