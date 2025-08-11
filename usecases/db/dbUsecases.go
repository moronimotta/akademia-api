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

func (u *DbUsecase) AddCoursesToUser(userID string, coursesID []string) error {
	for _, courseID := range coursesID {
		allClasses, err := u.Repository.Content.GetAllClassesByCourseID(courseID)
		if err != nil {
			return err
		}

		if err := u.Repository.UserProgress.AddCourseToUser(userID, courseID, allClasses); err != nil {
			return err
		}
	}

	return nil
}
