package repositories

import "akademia-api/entities"

type AkademiaRepository interface {
	CreatePost(post entities.Posts) error
	GetPostByID(id string) (*entities.Posts, error)
	GetAllPosts() ([]entities.Posts, error)
	UpdatePost(post *entities.Posts) error
	DeletePost(id string) error

	CreateCourse(course entities.Courses) error
	GetCourseByID(id string) (*entities.Courses, error)
	GetAllCourses() ([]entities.Courses, error)
	UpdateCourse(course *entities.Courses) error
	DeleteCourse(id string) error

	CreateClass(class entities.Classes) error
	CreateClasses(classes []entities.Classes) error
	GetClassByID(id string) (*entities.Classes, error)
	GetAllClasses() ([]entities.Classes, error)
	GetAllClassesByCourseID(courseID string) ([]entities.Classes, error)
	UpdateClass(class *entities.Classes) error
	DeleteClass(id string) error
}
