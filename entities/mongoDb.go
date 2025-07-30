package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCoursesInfo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Courses   []UserCourse       `json:"courses" bson:"courses"` // List of courses the user is enrolled in
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
	DeletedAt string             `json:"deleted_at" bson:"deleted_at"`
}

type UserCourse struct {
	CourseID         string              `json:"course_id" bson:"course_id"`
	Classes          []UserCourseClasses `json:"classes" bson:"classes"` // List of classes in the course
	TotalClasses     int                 `json:"total_classes" bson:"total_classes"`
	CompletedClasses int                 `json:"completed_classes" bson:"completed_classes"`
	CreatedAt        string              `json:"created_at" bson:"created_at"`
	UpdatedAt        string              `json:"updated_at" bson:"updated_at"`
	DeletedAt        string              `json:"deleted_at" bson:"deleted_at"`
}

type UserCourseClasses struct {
	ClassID  string `json:"class_id" bson:"class_id"`
	Finished bool   `json:"finished" bson:"finished"`
}
