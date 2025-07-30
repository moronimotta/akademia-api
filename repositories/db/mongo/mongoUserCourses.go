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
	CreateUserCourseInfo(userCourse entities.UserCoursesInfo) error
	GetUserCourseByID(id string) (*entities.UserCoursesInfo, error)
	GetUserCourseInfoByUserID(userID string) (*entities.UserCoursesInfo, error)
	UpdateUserCourseProgress(userID, courseID string) error
	AddCourseToUser(userID, courseID string, classes []entities.Classes) error
	UpdateClassStatus(userID, courseID, classID string) error
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

func (r *UserCoursesRepository) UpdateUserCourseProgress(userID, courseID string) error {
	filter := map[string]interface{}{
		"user_id":           userID,
		"courses.course_id": courseID,
	}

	update := map[string]interface{}{
		"$inc": map[string]interface{}{
			"courses.$.completed_classes": 1,
		},
		"$set": map[string]interface{}{
			"courses.$.updated_at": time.Now().Format("2006-01-02 15:04:05"),
			"updated_at":           time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	if _, err := r.db.GetMongoDB().Collection("user_courses").UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

func (r *UserCoursesRepository) AddCourseToUser(userID, courseID string, classes []entities.Classes) error {
	course := entities.UserCourse{
		CourseID:         courseID,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		DeletedAt:        "",
		TotalClasses:     len(classes),
		CompletedClasses: 0,
	}
	course.Classes = make([]entities.UserCourseClasses, len(classes))
	for i, class := range classes {
		course.Classes[i] = entities.UserCourseClasses{
			ClassID:  class.ID,
			Finished: false,
		}
	}
	filter := map[string]interface{}{"user_id": userID}
	update := map[string]interface{}{
		"$push": map[string]interface{}{
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

func (r *UserCoursesRepository) UpdateClassStatus(userID, courseID, classID string) error {
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
					userCourseInfo.Courses[i].Classes[j].Finished = true
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
