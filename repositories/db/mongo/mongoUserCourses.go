package mongoRepository

import (
	"akademia-api/db"
	"akademia-api/entities"
	"context"
	"time"
)

type UserCoursesRepository struct {
	db db.Database
}

type UserCoursesInfo interface {
	// crud operations for user courses
	CreateUserCourseInfo(userCourse entities.UserCoursesInfo) error
	GetUserCourseByID(id string) (*entities.UserCoursesInfo, error)
	GetUserCourseInfoByUserID(userID string) (*entities.UserCoursesInfo, error)
	UpdateUserCourseProgress(userID, courseID string, completedClasses int) error
	AddCourseToUser(userID string, course entities.UserCourse, classes []string) error
	UpdateClassStatus(userID, courseID, classID string, finished bool) error
	DeleteUserCourseInfo(id string) error
}

func NewUserCoursesRepository(db db.Database) UserCoursesInfo {
	return &UserCoursesRepository{
		db: db,
	}
}

func (r *UserCoursesRepository) CreateUserCourseInfo(userCourse entities.UserCoursesInfo) error {

	userCourse.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	userCourse.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if _, err := r.db.GetMongoDB().Collection("user_courses").InsertOne(context.TODO(), userCourse); err != nil {
		return err
	}
	return nil
}

func (r *UserCoursesRepository) GetUserCourseInfoByUserID(userID string) (*entities.UserCoursesInfo, error) {
	userCourse := &entities.UserCoursesInfo{}
	if err := r.db.GetMongoDB().Collection("user_courses").FindOne(context.TODO(), map[string]interface{}{"user_id": userID}).Decode(userCourse); err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (r *UserCoursesRepository) UpdateUserCourseProgress(userID, courseID string, completedClasses int) error {
	filter := map[string]interface{}{
		"user_id":           userID,
		"courses.course_id": courseID,
	}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"courses.$.completed_classes": completedClasses,
			"courses.$.updated_at":        time.Now().Format("2006-01-02 15:04:05"),
			"updated_at":                  time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	if _, err := r.db.GetMongoDB().Collection("user_courses").UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

// When user buys a new course, it adds the course to the user's courses array
// TODO: Maybe send a request to postgres to get all the classes with the course id.
// classes {[class_id: "class_id", finished: false}, ...}]}
// Fill the classes with ids and finished status as false
// Get the len of []classes and create object UserCourseClasses based on that with ClassID and Finished as false by default
func (r *UserCoursesRepository) AddCourseToUser(userID string, course entities.UserCourse, classes []string) error {
	course.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	course.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	course.DeletedAt = ""
	course.TotalClasses = len(classes)
	course.CompletedClasses = 0
	course.Classes = make([]entities.UserCourseClasses, len(classes))
	for i, classID := range classes {
		course.Classes[i] = entities.UserCourseClasses{
			ClassID:  classID,
			Finished: false,
		}
	}
	filter := map[string]interface{}{"user_id": userID}
	update := map[string]interface{}{
		"$addToSet": map[string]interface{}{
			"courses": course,
		},
		"$set": map[string]interface{}{
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	if _, err := r.db.GetMongoDB().Collection("user_courses").UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

// Update the status of a specific class in a course
// This function updates a class status using a simpler approach
func (r *UserCoursesRepository) UpdateClassStatus(userID, courseID, classID string, finished bool) error {
	// First, get the user's course info
	userCourseInfo, err := r.GetUserCourseInfoByUserID(userID)
	if err != nil {
		return err
	}

	// Find and update the specific class in memory
	for i, course := range userCourseInfo.Courses {
		if course.CourseID == courseID {
			for j, class := range course.Classes {
				if class.ClassID == classID {
					userCourseInfo.Courses[i].Classes[j].Finished = finished
					userCourseInfo.Courses[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
					userCourseInfo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
					break
				}
			}
			break
		}
	}

	// Update the entire document
	filter := map[string]interface{}{"user_id": userID}
	update := map[string]interface{}{
		"$set": userCourseInfo,
	}

	if _, err := r.db.GetMongoDB().Collection("user_courses").UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

func (r *UserCoursesRepository) DeleteUserCourseInfo(id string) error {
	if _, err := r.db.GetMongoDB().Collection("user_courses").DeleteOne(context.TODO(), map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}

func (r *UserCoursesRepository) GetUserCourseByID(id string) (*entities.UserCoursesInfo, error) {
	userCourse := &entities.UserCoursesInfo{}
	if err := r.db.GetMongoDB().Collection("user_courses").FindOne(context.TODO(), map[string]interface{}{"id": id}).Decode(userCourse); err != nil {
		return nil, err
	}
	return userCourse, nil
}
