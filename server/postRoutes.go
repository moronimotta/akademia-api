package server

import (
	"akademia-api/entities"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (s *Server) initPostRoutes() {
	s.app.GET("/posts", func(c *gin.Context) {
		output, err := s.dbHandler.Repository.Content.GetAllPosts()
		if err != nil {
			slog.Error("Failed to get all posts: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully retrieved all posts")
		c.JSON(200, output)
	})

	s.app.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		post, err := s.dbHandler.Repository.Content.GetPostByID(id)
		if err != nil {
			slog.Error("Failed to get post by ID: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if post == nil {
			slog.Warn("Post not found", "id", id)
			c.JSON(404, gin.H{"error": "Post not found"})
			return
		}
		slog.Info("Successfully retrieved post", "id", id)
		c.JSON(200, post)
	})
	s.app.POST("/posts", func(c *gin.Context) {
		var post entities.Posts
		if err := c.ShouldBindJSON(&post); err != nil {
			slog.Error("Invalid input for post creation: %v", err)
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		if err := s.dbHandler.Repository.Content.CreatePost(post); err != nil {
			slog.Error("Failed to create post: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully created post")
		c.JSON(201, gin.H{"message": "Post created successfully"})
	})
	s.app.PUT("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var post entities.Posts
		if err := c.ShouldBindJSON(&post); err != nil {
			slog.Error("Invalid input for post update: %v", err)
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		post.ID = id // Ensure the ID is set for the update
		if err := s.dbHandler.Repository.Content.UpdatePost(&post); err != nil {
			slog.Error("Failed to update post: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully updated post", "id", id)
		c.JSON(200, gin.H{"message": "Post updated successfully"})
	})
	s.app.DELETE("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := s.dbHandler.Repository.Content.DeletePost(id); err != nil {
			slog.Error("Failed to delete post: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		slog.Info("Successfully deleted post", "id", id)
		c.JSON(200, gin.H{"message": "Post deleted successfully"})
	})
}
