package usecases

import (
	"akademia-api/db"
	"akademia-api/repositories"
	mongoRepository "akademia-api/repositories/db/mongo"
	postgresRepository "akademia-api/repositories/db/postgres"
)

type DbUsecase struct {
	Repository repositories.AkademiaRepository
}

func NewDbUsecase(db db.Database) *DbUsecase {

	pgRepository := postgresRepository.NewPostgresRepository(db)
	mongoRepository := mongoRepository.NewMongoRepository(db)

	return &DbUsecase{
		Repository: repositories.AkademiaRepository{
			Content:      pgRepository,
			UserProgress: mongoRepository,
		},
	}
}

func (d *DbUsecase) MarkClassAsCompleted(userID, courseID, classID string) error {
	if err := d.Repository.UserProgress.UpdateClassStatus(userID, courseID, classID); err != nil {
		return err
	}

	if err := d.Repository.UserProgress.UpdateUserCourseProgress(userID, courseID); err != nil {
		return err
	}
	return nil
}
