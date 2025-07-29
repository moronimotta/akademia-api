package handlers

import (
	"akademia-api/db"
	usecases "akademia-api/usecases/db"
	"errors"
	"log"

	messageWorker "github.com/moronimotta/message-worker-module"
	"github.com/redis/go-redis/v9"
)

type RabbitMqHandler struct {
	DbUsecase   usecases.DbUsecase
	RedisClient *redis.Client
}

func NewRabbitMqHandler(db db.Database, redisClient *redis.Client) *RabbitMqHandler {
	var usecaseDb usecases.DbUsecase

	switch db.GetSQLDB().Dialector.Name() {
	case "postgres":
		usecaseDb = *usecases.NewPgUsecase(db)
	default:
		usecaseDb = usecases.DbUsecase{}
	}

	return &RabbitMqHandler{
		DbUsecase:   usecaseDb,
		RedisClient: redisClient,
	}
}

func (h *RabbitMqHandler) EventBus(event messageWorker.Event) error {

	switch event.Event {
	case "product_created":
		// TODO: Creates a course
		log.Println("Product created event received")
	case "courses.getAll":
		// TODO: Get all courses
		// TODO: Return to a queue
		log.Println("Get all courses event received")
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
