package postgresRepository

import (
	"akademia-api/db"
	"akademia-api/entities"
)

type PgCourses interface {
	CreateCourse(course *entities.Courses) error
	GetCourseByID(id string) (*entities.Courses, error)
	GetAllCourses() ([]entities.Courses, error)
	UpdateCourse(course *entities.Courses) error
	DeleteCourse(id string) error
	GetDraftCourses() ([]entities.Courses, error)
}

type pgCoursesRepository struct {
	db db.Database
}

func NewPgCoursesRepository(db db.Database) PgCourses {
	return &pgCoursesRepository{
		db: db,
	}
}
func (r *pgCoursesRepository) CreateCourse(course *entities.Courses) error {

	if err := r.db.GetSQLDB().Create(course).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgCoursesRepository) GetCourseByID(id string) (*entities.Courses, error) {
	course := &entities.Courses{}
	if err := r.db.GetSQLDB().Where("id = ?", id).First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}
func (r *pgCoursesRepository) GetAllCourses() ([]entities.Courses, error) {
	posts := []entities.Courses{}
	if err := r.db.GetSQLDB().Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *pgCoursesRepository) UpdateCourse(course *entities.Courses) error {
	if err := r.db.GetSQLDB().Model(&entities.Courses{}).Where("id = ?", course.ID).Updates(course).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgCoursesRepository) DeleteCourse(id string) error {
	if err := r.db.GetSQLDB().Where("id = ?", id).Delete(&entities.Courses{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *pgCoursesRepository) GetDraftCourses() ([]entities.Courses, error) {
	var drafts []entities.Courses
	if err := r.db.GetSQLDB().Where("status = ?", "draft").Find(&drafts).Error; err != nil {
		return nil, err
	}
	return drafts, nil
}
