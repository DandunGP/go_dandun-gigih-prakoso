package repository

import (
	"weekly3/entities"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BlogRepository interface {
	GetAllBlogs(data []entities.Blog) ([]entities.Blog, error)
	CreateBlog(data entities.Blog) error
	GetBlogByID(id uuid.UUID) (entities.Blog, error)
	UpdateBlog(id uuid.UUID, input entities.Blog) error
	DeleteBlog(id uuid.UUID) error
	GetBlogByCat(id int) ([]entities.Blog, error)
	GetBlogByKey(key string) ([]entities.Blog, error)
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

func (r *blogRepository) GetAllBlogs(data []entities.Blog) ([]entities.Blog, error) {
	if err := r.db.Preload("User").Preload("Category").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *blogRepository) GetBlogByID(id uuid.UUID) (entities.Blog, error) {
	var blog entities.Blog

	if err := r.db.Where("id = ?", id).First(&blog).Error; err != nil {
		return blog, err
	}

	return blog, nil
}

func (r *blogRepository) CreateBlog(data entities.Blog) error {
	return r.db.Create(&data).Error
}

func (r *blogRepository) UpdateBlog(id uuid.UUID, input entities.Blog) error {
	var blog entities.Blog
	r.db.Where("id = ?", id).First(&blog)

	return r.db.Model(&blog).Updates(&input).Error
}

func (r *blogRepository) DeleteBlog(id uuid.UUID) error {
	var blog entities.Blog
	if err := r.db.Where("id = ?", id).Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}

func (r *blogRepository) GetBlogByCat(id int) ([]entities.Blog, error) {
	var blog []entities.Blog

	if err := r.db.Where("category_id = ?", id).Find(&blog).Error; err != nil {
		return blog, err
	}

	return blog, nil
}

func (r *blogRepository) GetBlogByKey(key string) ([]entities.Blog, error) {
	var blog []entities.Blog

	if err := r.db.Where("title LIKE ?", "%"+key+"%").Find(&blog).Error; err != nil {
		return blog, err
	}

	return blog, nil
}
