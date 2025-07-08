package postgresRepository

import (
	"akademia-api/db"
	"akademia-api/entities"
)

type PgClasses interface {
	CreateClass(class entities.Classes) error
	GetClassByID(id string) (*entities.Classes, error)
	GetAllClasses() ([]entities.Classes, error)
	GetAllClassesByCourseID(courseID string) ([]entities.Classes, error)
	UpdateClass(class *entities.Classes) error
	DeleteClass(id string) error
}

type pgClassesRepository struct {
	db db.Database
}

func NewPgClassesRepository(db db.Database) PgClasses {
	return &pgClassesRepository{
		db: db,
	}
}
func (r *pgClassesRepository) CreateClass(class entities.Classes) error {

	if err := r.db.GetDB().Create(&class).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgClassesRepository) GetClassByID(id string) (*entities.Classes, error) {
	class := &entities.Classes{}
	if err := r.db.GetDB().Where("id = ?", id).First(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}
func (r *pgClassesRepository) GetAllClasses() ([]entities.Classes, error) {
	posts := []entities.Classes{}
	if err := r.db.GetDB().Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *pgClassesRepository) GetAllClassesByCourseID(courseID string) ([]entities.Classes, error) {
	posts := []entities.Classes{}
	if err := r.db.GetDB().Where("course_id = ?", courseID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *pgClassesRepository) UpdateClass(class *entities.Classes) error {
	if err := r.db.GetDB().Model(&entities.Classes{}).Where("id = ?", class.ID).Updates(class).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgClassesRepository) DeleteClass(id string) error {
	if err := r.db.GetDB().Where("id = ?", id).Delete(&entities.Classes{}).Error; err != nil {
		return err
	}
	return nil
}
