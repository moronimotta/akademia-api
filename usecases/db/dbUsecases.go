package usecases

import (
	"akademia-api/db"
	"akademia-api/entities"
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

func (d *DbUsecase) GetAllDrafts() ([]entities.CoursesClassesOutput, error) {
	courses, err := d.Repository.Content.GetDraftCourses()
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(courses))
	for i, course := range courses {
		ids[i] = course.ID
	}

	classes, err := d.Repository.Content.GetClassesByCoursesID(ids)
	if err != nil {
		return nil, err
	}

	var output []entities.CoursesClassesOutput

	output = append(output, entities.CoursesClassesOutput{
		Courses: courses,
		Classes: classes,
	})

	return output, nil
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
