package server

import (
	"akademia-api/entities"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (s *Server) initClassesRoutes() {
	s.app.GET("/classes", func(c *gin.Context) {
		output, err := s.dbHandler.Repository.Content.GetAllClasses()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, output)
	})

	s.app.GET("/classes/:id", func(c *gin.Context) {
		id := c.Param("id")
		class, err := s.dbHandler.Repository.Content.GetClassByID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if class == nil {
			c.JSON(404, gin.H{"error": "Class not found"})
			return
		}
		c.JSON(200, class)
	})

	s.app.POST("/classes", func(c *gin.Context) {
		rawData, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{"error": "Could not read request body"})
			return
		}

		var classes []entities.Classes

		if err := json.Unmarshal(rawData, &classes); err != nil {
			var single entities.Classes
			if err := json.Unmarshal(rawData, &single); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			if err := s.dbHandler.Repository.Content.CreateClass(single); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if len(classes) > 0 {
			if err := s.dbHandler.Repository.Content.CreateClasses(classes); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(201, gin.H{"message": "Class(es) created successfully"})
	})

	s.app.PUT("/classes/:id", func(c *gin.Context) {
		id := c.Param("id")
		var class entities.Classes
		if err := c.ShouldBindJSON(&class); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		class.ID = id // Ensure the ID is set for the update
		if err := s.dbHandler.Repository.Content.UpdateClass(&class); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Class updated successfully"})
	})

	s.app.DELETE("/classes/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := s.dbHandler.Repository.Content.DeleteClass(id); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Class deleted successfully"})
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

}
