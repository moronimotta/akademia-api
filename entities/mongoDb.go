package entities

type UserCourses struct {
	ID        string       `json:"id" bson:"_id"`
	UserID    string       `json:"user_id" bson:"user_id"`
	Courses   []UserCourse `json:"courses" bson:"courses"` // List of courses the user is enrolled in
	CreatedAt string       `json:"created_at" bson:"created_at"`
	UpdatedAt string       `json:"updated_at" bson:"updated_at"`
	DeletedAt string       `json:"deleted_at" bson:"deleted_at"`
}

type UserCourse struct {
	CourseID         string  `json:"course_id" bson:"course_id"`
	TotalClasses     int     `json:"total_classes" bson:"total_classes"`
	CompletedClasses int     `json:"completed_classes" bson:"completed_classes"`
	Progress         float64 `json:"progress" bson:"progress"`
	CreatedAt        string  `json:"created_at" bson:"created_at"`
	UpdatedAt        string  `json:"updated_at" bson:"updated_at"`
	DeletedAt        string  `json:"deleted_at" bson:"deleted_at"`
}

type UserCourseClasses struct {
	ClassID  string `json:"class_id" bson:"class_id"`
	Finished bool   `json:"finished" bson:"finished"`
}
