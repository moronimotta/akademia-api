package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Courses struct {
	ID          string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"not null"` // e.g. "active", "draft", "archived"
	ProductID   string `json:"product_id" gorm:"column:product_id"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   string `json:"deleted_at" gorm:"column:deleted_at"`
}

type CourseInput struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`     // e.g. "active", "draft", "archived"
	ProductID   string    `json:"product_id"` //TODO: Later will be required.
	Classes     []Classes `json:"classes"`    // List of classes in the course
}

type CoursesClassesOutput struct {
	Courses []Courses `json:"courses"`
	Classes []Classes `json:"classes"`
}

type CourseClassesOutput struct {
	Course  Courses   `json:"course"`
	Classes []Classes `json:"classes"`
}

type Classes struct {
	ID          string `json:"id"`
	Order       int    `json:"order"`     // Order of the class in the course
	CourseID    string `json:"course_id"` // ID of the course this class belongs tox
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"` // TESTING
	// Content     Content `json:"content"` // Content of the class, can be video, pdf, etc.
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
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
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"`
	Banner   string `json:"banner"` // TESTING
	// Banner   Content `json:"banner"`    // Banner image or video for the post
	Category string `json:"category"`  // e.g. "biblical studies", "religion"...
	ReadTime int    `json:"read_time"` // Estimated read time in minutes. Calculated based on content length
	Featured bool   `json:"featured"`  // Whether the post is featured on the homepage
	// TODO: READ TIME CALCULATED BASED ON CONTENT LENGTH
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (p *Posts) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()

	now := time.Now().UTC().Format(time.RFC3339)
	p.CreatedAt = now
	p.UpdatedAt = now
	return
}

func (p *Posts) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	return
}

func (c *Courses) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
	return
}

func (c *Courses) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	return
}

func (c *Classes) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
	return
}

func (c *Classes) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	return
}

func (c *Content) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	c.CreatedAt = now
	c.UpdatedAt = now
	return
}

func (c *Content) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	return
}
