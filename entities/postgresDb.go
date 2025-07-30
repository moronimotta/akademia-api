package entities

import (
	"time"

	"github.com/google/uuid"
)

type Courses struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"` // e.g. "active", "draft", "archived"
	ProductID   string `json:"product_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type Classes struct {
	ID          string  `json:"id"`
	Order       int     `json:"order"`     // Order of the class in the course
	CourseID    string  `json:"course_id"` // ID of the course this class belongs tox
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Content     Content `json:"content"` // Content of the class, can be video, pdf, etc.
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   string  `json:"deleted_at"`
}

type Content struct {
	RawText    *string `json:"rawText,omitempty"`
	S3Key      *string `json:"s3Key,omitempty"`
	Type       *string `json:"type,omitempty"` // "video", "pdf", etc.
	Duration   *int    `json:"duration,omitempty"`
	Format     *string `json:"format,omitempty"`
	Thumbnail  *string `json:"thumbnail,omitempty"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  string  `json:"deleted_at"`
	ForeignKey string  `json:"foreign_key,omitempty"` // e.g. ID of the post or class this content belongs to
}

type Posts struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	AuthorID string  `json:"author_id"`
	Banner   Content `json:"banner"`    // Banner image or video for the post
	Category string  `json:"category"`  // e.g. "biblical studies", "religion"...
	ReadTime int     `json:"read_time"` // Estimated read time in minutes. Calculated based on content length
	Featured bool    `json:"featured"`  // Whether the post is featured on the homepage
	// TODO: READ TIME CALCULATED BASED ON CONTENT LENGTH
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (p *Posts) BeforeCreate() {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	now := time.Now().UTC().Format(time.RFC3339)
	p.CreatedAt = now
	p.UpdatedAt = now
}

func (p *Posts) BeforeUpdate() {
	p.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

func (c *Courses) BeforeCreate() {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
}

func (c *Courses) BeforeUpdate() {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

func (c *Classes) BeforeCreate() {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
}

func (c *Classes) BeforeUpdate() {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

func (c *Content) BeforeCreate() {
	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
}

func (c *Content) BeforeUpdate() {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}
