package mongoRepository

import (
	"akademia-api/db"
	"akademia-api/repositories"
)

type MongoRepository struct {
	UserCoursesInfo
}

func NewMongoRepository(db db.Database) repositories.UserProgressRepository {
	return &MongoRepository{
		UserCoursesInfo: NewUserCoursesRepository(db),
	}
}
