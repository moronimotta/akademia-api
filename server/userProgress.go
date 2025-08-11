package server

import (
	"akademia-api/entities"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (s *Server) initUserProgressRoutes() {

	s.app.GET("/user-progress", func(c *gin.Context) {
		output, err := s.dbHandler.Repository.UserProgress.GetAllUserCourses()
		if err != nil {
			slog.Error("Failed to get all user courses: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all user courses")
		c.JSON(200, output)
	})

	// GET USER COURSE INFO BY USER ID
	s.app.GET("/user-progress/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		output, err := s.dbHandler.Repository.UserProgress.GetUserCourseInfoByUserID(userID)
		if err != nil {
			slog.Error("Failed to get user course info: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved user course info")
		c.JSON(200, output)
	})

	// CREATE USER COURSE INFO
	s.app.POST("/user-progress", func(c *gin.Context) {
		var userProgress entities.UserCoursesInfo
		if err := c.ShouldBindJSON(&userProgress); err != nil {
			slog.Error("Invalid input for user progress creation: %v", err)
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		if err := s.dbHandler.Repository.UserProgress.CreateUserCourseInfo(userProgress); err != nil {
			slog.Error("Failed to create user course info: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully created user progress")
		c.JSON(201, gin.H{"message": "User progress created successfully"})
	})

	s.app.PUT("/user-progress/completed-class/:userID/:courseID/:classID", func(c *gin.Context) {
		userID := c.Param("userID")
		courseID := c.Param("courseID")
		classID := c.Param("classID")
		if err := s.dbHandler.MarkClassAsCompleted(userID, courseID, classID); err != nil {
			slog.Error("Failed to mark class as completed: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully marked class as completed")
		c.JSON(200, gin.H{"message": "Class marked as completed successfully"})
	})

	// DELETE USER COURSE INFO
	s.app.DELETE("/user-progress/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := s.dbHandler.Repository.UserProgress.DeleteUserCourseInfo(id); err != nil {
			slog.Error("Failed to delete user course info: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully deleted user course info")
		c.JSON(200, gin.H{"message": "User course info deleted successfully"})
	})

	// ADD COURSE TO USER
	s.app.POST("/user-progress/add-course/:userID/:courseID", func(c *gin.Context) {
		userID := c.Param("userID")
		courseID := c.Param("courseID")
		classes, err := s.dbHandler.Repository.Content.GetAllClassesByCourseID(courseID)
		if err != nil {
			slog.Error("Failed to get classes by course ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = s.dbHandler.Repository.UserProgress.AddCourseToUser(userID, courseID, classes)
		if err != nil {
			slog.Error("Failed to add course to user: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully added course to user")
		c.JSON(200, gin.H{"message": "Course added to user successfully"})
	})

}
