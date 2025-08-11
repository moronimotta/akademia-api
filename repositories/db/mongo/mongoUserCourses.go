package mongoRepository

import (
	"akademia-api/db"
	"akademia-api/entities"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	GetAllUserCourses() ([]entities.UserCoursesInfo, error)
}

func NewUserCoursesRepository(db db.Database) UserCoursesInfo {
	return &UserCoursesRepository{
		db: db,
	}
}

func (r *UserCoursesRepository) GetAllUserCourses() ([]entities.UserCoursesInfo, error) {
	cursor, err := r.db.GetMongoDB().Collection("user_courses").Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var userCourses []entities.UserCoursesInfo
	if err := cursor.All(context.TODO(), &userCourses); err != nil {
		return nil, err
	}
	return userCourses, nil
}

func (r *UserCoursesRepository) CreateUserCourseInfo(userCourse entities.UserCoursesInfo) error {
	userCourse.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	userCourse.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// MongoDB will auto-generate ObjectID if ID is empty
	if userCourse.ID == primitive.NilObjectID {
		userCourse.ID = primitive.NewObjectID()
	}

	if userCourse.Courses == nil {
		userCourse.Courses = make([]entities.UserCourse, 0)
	}

	_, err := r.db.GetMongoDB().Collection("user_courses").InsertOne(context.TODO(), userCourse)
	if err != nil {
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

	result, err := r.db.GetMongoDB().Collection("user_courses").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	// If no document was modified, it means the user doesn't exist, so create a new document
	if result.ModifiedCount == 0 {
		userCourse := entities.UserCoursesInfo{
			UserID:    userID,
			Courses:   []entities.UserCourse{course},
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		if _, err := r.db.GetMongoDB().Collection("user_courses").InsertOne(context.TODO(), userCourse); err != nil {
			return err
		}
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
					if class.Finished {
						return errors.New("Class already marked as finished") // Class already marked as finished
					}
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err := r.db.GetMongoDB().Collection("user_courses").DeleteOne(context.TODO(), map[string]interface{}{"_id": objectID}); err != nil {
		return err
	}
	return nil
}

func (r *UserCoursesRepository) GetUserCourseByID(id string) (*entities.UserCoursesInfo, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	userCourse := &entities.UserCoursesInfo{}
	if err := r.db.GetMongoDB().Collection("user_courses").FindOne(context.TODO(), map[string]interface{}{"_id": objectID}).Decode(userCourse); err != nil {
		return nil, err
	}
	return userCourse, nil
}
