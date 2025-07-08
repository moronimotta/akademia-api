package handlers

import (
	"akademia-api/db"
	usecases "akademia-api/usecases/db"
	"errors"
	"log"

	messageWorker "github.com/moronimotta/message-worker-module"
)

type RabbitMqHandler struct {
	DbUsecase usecases.DbUsecase
}

func NewRabbitMqHandler(db db.Database) *RabbitMqHandler {
	var usecaseDb usecases.DbUsecase

	switch db.GetDB().Dialector.Name() {
	case "postgres":
		usecaseDb = *usecases.NewPgUsecase(db)
	default:
		usecaseDb = usecases.DbUsecase{}
	}

	return &RabbitMqHandler{
		DbUsecase: usecaseDb,
	}
}

func (h *RabbitMqHandler) EventBus(event messageWorker.Event) error {

	switch event.Event {
	case "product_created":
		// TODO: Creates a course
		log.Println("Product created event received")
	default:
		return errors.New("event not found")
	}
	return nil
}

// func (h *RabbitMqHandler) PublishMessage(topicName, eventName string, data map[string]string) error {

// 	input := messageWorker.Publisher{}
// 	input.ConnectionURL = os.Getenv("RABBITMQ_URL")
// 	input.TopicName = topicName

// 	messageInput := messageWorker.Event{
// 		Event: eventName,
// 		Data:  data,
// 	}

//		messageWorker.SendMessage(input, messageInput)
//		return nil
//	}
