package handlers

import (
	"akademia-api/db"
	"akademia-api/entities"
	usecases "akademia-api/usecases/db"
	"errors"
	"log"
	"log/slog"

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
		usecaseDb = *usecases.NewDbUsecase(db)
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
	case "user.new_course":

		eventData, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.New("invalid event data format")
		}

		// Get user_id
		userIDRaw, ok := eventData["user_id"]
		if !ok {
			slog.Error("missing user_id in event data")
			return errors.New("missing user_id")
		}
		userID, ok := userIDRaw.(string)
		if !ok {
			return errors.New("user_id must be string")
		}

		var localProductsIds []interface{}
		if ids, ok := eventData["local_products_ids"].([]interface{}); ok {
			localProductsIds = ids
		} else if ids, ok := eventData["local_product_ids"].([]interface{}); ok {
			localProductsIds = ids
		} else {
			return errors.New("missing or invalid local_products_ids/local_product_ids")
		}

		coursesId := make([]string, 0, len(localProductsIds))
		for i, id := range localProductsIds {
			courseID, ok := id.(string)
			if !ok {
				return errors.New("course id at index " + string(rune(i+'0')) + " must be string")
			}
			coursesId = append(coursesId, courseID)
		}

		if err := h.DbUsecase.AddCoursesToUser(userID, coursesId); err != nil {
			log.Printf("Error creating user course info: %v", err)
			return err
		}

		log.Println("User new course event received")

	case "user.created":
		eventData, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.New("invalid event data format")
		}
		userID := eventData["user_id"].(string)
		userCourse := entities.UserCoursesInfo{
			UserID: userID,
		}

		if err := h.DbUsecase.Repository.UserProgress.CreateUserCourseInfo(userCourse); err != nil {
			log.Printf("Error creating user course info: %v", err)
			return err
		}

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
