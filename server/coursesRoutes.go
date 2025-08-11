package server

import (
	"akademia-api/entities"

	"github.com/gin-gonic/gin"
)

func (s *Server) initCoursesRoutes() {
	s.app.GET("/courses", func(c *gin.Context) {
		output, err := s.dbHandler.Repository.Content.GetAllCourses()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, output)
	})

	s.app.GET("/courses/product/:productID", func(c *gin.Context) {
		productID := c.Param("productID")
		course, err := s.dbHandler.Repository.Content.GetCourseByProductID(productID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, course)
	})

	s.app.GET("/courses/:id", func(c *gin.Context) {
		id := c.Param("id")
		course, err := s.dbHandler.Repository.Content.GetCourseByID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if course == nil {
			c.JSON(404, gin.H{"error": "Course not found"})
			return
		}
		c.JSON(200, course)
	})

	s.app.GET("/courses/drafts", func(c *gin.Context) {
		drafts, err := s.dbHandler.GetAllDrafts()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, drafts)
	})

	s.app.POST("/courses", func(c *gin.Context) {
		var input entities.CourseInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if err := s.dbHandler.CreateFullCourse(input); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Course created successfully"})
	})

	s.app.PUT("/courses", func(c *gin.Context) {
		var course entities.CourseInput
		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		if err := s.dbHandler.UpdateCourse(&course); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Course updated successfully"})
	})

	s.app.DELETE("/courses/:id", func(c *gin.Context) {
		id := c.Param("id")

		if err := s.dbHandler.Repository.Content.DeleteClassesByCourseID(id); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := s.dbHandler.Repository.Content.DeleteCourse(id); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Course deleted successfully"})
	})
	s.app.GET("/courses/:id/classes", func(c *gin.Context) {
		courseID := c.Param("id") // Changed from "courseID" to "id"
		classes, err := s.dbHandler.Repository.Content.GetAllClassesByCourseID(courseID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, classes)
	})

	s.app.GET("/courses/:id/full-info", func(c *gin.Context) {
		courseID := c.Param("id")
		output, err := s.dbHandler.GetFullCourseInfo(courseID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if output == nil {
			c.JSON(404, gin.H{"error": "Course not found"})
			return
		}
		c.JSON(200, output)
	})

	s.app.GET("/courses/all/full-info", func(c *gin.Context) {
		output, err := s.dbHandler.GetAllFullCoursesInfo()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, output)
	})

}
