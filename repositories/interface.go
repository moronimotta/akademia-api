package repositories

import "akademia-api/entities"

type AkademiaRepository struct {
	Content      ContentRepository
	UserProgress UserProgressRepository
}

type ContentRepository interface {
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
	DeleteClassesByCourseID(courseID string) error
}

type UserProgressRepository interface {
	CreateUserCourseInfo(userCourse entities.UserCoursesInfo) error
	GetUserCourseByID(id string) (*entities.UserCoursesInfo, error)
	GetUserCourseInfoByUserID(userID string) (*entities.UserCoursesInfo, error)
	UpdateUserCourseProgress(userID, courseID string) error
	AddCourseToUser(userID, courseID string, classes []entities.Classes) error
	UpdateClassStatus(userID, courseID, classID string) error
	DeleteUserCourseInfo(id string) error
}
