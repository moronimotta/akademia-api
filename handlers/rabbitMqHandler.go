package handlers

import (
	"akademia-api/db"
	"akademia-api/entities"
	usecases "akademia-api/usecases/db"
	"errors"
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
		slog.Info("Processing user.new_course event")

		eventData, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.New("Event: user.new_course - invalid event data format")
		}

		userIDRaw, ok := eventData["user_id"]
		if !ok {
			return errors.New("Event: user.new_course - missing user_id")
		}
		userID, ok := userIDRaw.(string)
		if !ok {
			return errors.New("Event: user.new_course - user_id must be string")
		}

		var localProductsIds []interface{}
		if ids, ok := eventData["local_products_ids"].([]interface{}); ok {
			localProductsIds = ids
		} else if ids, ok := eventData["local_product_ids"].([]interface{}); ok {
			localProductsIds = ids
		} else {
			return errors.New("Event: user.new_course - missing or invalid local_products_ids/local_product_ids")
		}

		coursesId := make([]string, 0, len(localProductsIds))
		for i, id := range localProductsIds {
			courseID, ok := id.(string)
			if !ok {
				return errors.New("Event: user.new_course - course id at index " + string(rune(i+'0')) + " must be string")
			}
			coursesId = append(coursesId, courseID)
		}

		if err := h.DbUsecase.AddCoursesToUser(userID, coursesId); err != nil {
			return errors.New("Event: user.new_course - error creating user course info: " + err.Error())
		}

		slog.Info("Event: user.new_course - User new course event processed successfully")

	case "user.created":
		slog.Info("Event: user.created - Processing user.created event")
		eventData, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.New("Event: user.created - invalid event data format")
		}
		userID := eventData["user_id"].(string)
		userCourse := entities.UserCoursesInfo{
			UserID: userID,
		}

		if err := h.DbUsecase.Repository.UserProgress.CreateUserCourseInfo(userCourse); err != nil {
			return errors.New("Event: user.created - error creating user course info: " + err.Error())
		}

	default:
		return errors.New("event not found")
	}
	return nil
}
