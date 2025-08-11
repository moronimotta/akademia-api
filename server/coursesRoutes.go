package server

import (
	"akademia-api/entities"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (s *Server) initCoursesRoutes() {
	s.app.GET("/courses", func(c *gin.Context) {
		output, err := s.dbHandler.Repository.Content.GetAllCourses()
		if err != nil {
			slog.Error("Failed to get all courses: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all courses")
		c.JSON(200, output)
	})

	s.app.GET("/courses/product/:productID", func(c *gin.Context) {
		productID := c.Param("productID")
		course, err := s.dbHandler.Repository.Content.GetCourseByProductID(productID)
		if err != nil {
			slog.Error("Failed to get course by product ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved course by product ID", "productID", productID)
		c.JSON(200, course)
	})

	s.app.GET("/courses/:id", func(c *gin.Context) {
		id := c.Param("id")
		course, err := s.dbHandler.Repository.Content.GetCourseByID(id)
		if err != nil {
			slog.Error("Failed to get course by ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if course == nil {
			slog.Warn("Course not found", "id", id)
			c.JSON(404, gin.H{"error": "Course not found"})
			return
		}
		slog.Info("Successfully retrieved course", "id", id)
		c.JSON(200, course)
	})

	s.app.GET("/courses/drafts", func(c *gin.Context) {
		drafts, err := s.dbHandler.GetAllDrafts()
		if err != nil {
			slog.Error("Failed to get all drafts: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all drafts")
		c.JSON(200, drafts)
	})

	s.app.POST("/courses", func(c *gin.Context) {
		var input entities.CourseInput
		if err := c.ShouldBindJSON(&input); err != nil {
			slog.Error("Invalid input for course creation: %v", err)
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if err := s.dbHandler.CreateFullCourse(input); err != nil {
			slog.Error("Failed to create full course: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully created full course")
		c.JSON(201, gin.H{"message": "Course created successfully"})
	})

	s.app.PUT("/courses", func(c *gin.Context) {
		var course entities.CourseInput
		if err := c.ShouldBindJSON(&course); err != nil {
			slog.Error("Invalid input for course update: %v", err)
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if err := s.dbHandler.UpdateCourse(&course); err != nil {
			slog.Error("Failed to update course: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully updated course", "id", course.ID)
		c.JSON(200, gin.H{"message": "Course updated successfully"})
	})

	s.app.DELETE("/courses/:id", func(c *gin.Context) {
		id := c.Param("id")

		if err := s.dbHandler.Repository.Content.DeleteClassesByCourseID(id); err != nil {
			slog.Error("Failed to delete classes by course ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := s.dbHandler.Repository.Content.DeleteCourse(id); err != nil {
			slog.Error("Failed to delete course: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully deleted course", "id", id)
		c.JSON(200, gin.H{"message": "Course deleted successfully"})
	})
	s.app.GET("/courses/:id/classes", func(c *gin.Context) {
		courseID := c.Param("id") // Changed from "courseID" to "id"
		classes, err := s.dbHandler.Repository.Content.GetAllClassesByCourseID(courseID)
		if err != nil {
			slog.Error("Failed to get all classes by course ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all classes by course ID", "courseID", courseID)
		c.JSON(200, classes)
	})

	s.app.GET("/courses/:id/full-info", func(c *gin.Context) {
		courseID := c.Param("id")
		output, err := s.dbHandler.GetFullCourseInfo(courseID)
		if err != nil {
			slog.Error("Failed to get full course info: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if output == nil {
			slog.Warn("Course not found", "id", courseID)
			c.JSON(404, gin.H{"error": "Course not found"})
			return
		}
		slog.Info("Successfully retrieved full course info", "courseID", courseID)
		c.JSON(200, output)
	})

	s.app.GET("/courses/all/full-info", func(c *gin.Context) {
		output, err := s.dbHandler.GetAllFullCoursesInfo()
		if err != nil {
			slog.Error("Failed to get all full courses info: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all full courses info")
		c.JSON(200, output)
	})

}
