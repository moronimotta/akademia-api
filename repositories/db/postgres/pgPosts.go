package postgresRepository

import (
	"akademia-api/db"
	"akademia-api/entities"
)

type PgPosts interface {
	CreatePost(post entities.Posts) error
	GetPostByID(id string) (*entities.Posts, error)
	GetAllPosts() ([]entities.Posts, error)
	UpdatePost(post *entities.Posts) error
	DeletePost(id string) error
}

type pgPostsRepository struct {
	db db.Database
}

func NewPgPostsRepository(db db.Database) PgPosts {
	return &pgPostsRepository{
		db: db,
	}
}
func (r *pgPostsRepository) CreatePost(post entities.Posts) error {

	if err := r.db.GetSQLDB().Create(&post).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgPostsRepository) GetPostByID(id string) (*entities.Posts, error) {
	post := &entities.Posts{}
	if err := r.db.GetSQLDB().Where("id = ?", id).First(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
func (r *pgPostsRepository) GetAllPosts() ([]entities.Posts, error) {
	posts := []entities.Posts{}
	if err := r.db.GetSQLDB().Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *pgPostsRepository) UpdatePost(post *entities.Posts) error {
	if err := r.db.GetSQLDB().Model(&entities.Posts{}).Where("id = ?", post.ID).Updates(post).Error; err != nil {
		return err
	}
	return nil
}
func (r *pgPostsRepository) DeletePost(id string) error {
	if err := r.db.GetSQLDB().Where("id = ?", id).Delete(&entities.Posts{}).Error; err != nil {
		return err
	}
	return nil
}
