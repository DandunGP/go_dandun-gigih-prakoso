package usecase

import (
	"weekly3/entities"
	"weekly3/repository"

	"github.com/google/uuid"
)

type BlogUsecase interface {
	Find(blogData []entities.Blog) ([]entities.Blog, error)
	Create(blogData entities.Blog) error
	FindByID(id uuid.UUID) (entities.Blog, error)
	Update(id uuid.UUID, input entities.Blog) (entities.Blog, error)
	Delete(id uuid.UUID) error
	FindByCat(id int) ([]entities.Blog, error)
	FindByKey(key string) ([]entities.Blog, error)
}

type blogUsecase struct {
	BlogRepository repository.BlogRepository
}

func NewBlogUsecase(blogRepo repository.BlogRepository) *blogUsecase {
	return &blogUsecase{BlogRepository: blogRepo}
}

func (s *blogUsecase) Create(BlogData entities.Blog) error {

	if err := s.BlogRepository.CreateBlog(BlogData); err != nil {
		return err
	}

	return nil
}

func (s *blogUsecase) Find(BlogData []entities.Blog) ([]entities.Blog, error) {

	if data, err := s.BlogRepository.GetAllBlogs(BlogData); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (s *blogUsecase) FindByID(id uuid.UUID) (entities.Blog, error) {
	if data, err := s.BlogRepository.GetBlogByID(id); err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (s *blogUsecase) Update(id uuid.UUID, input entities.Blog) (entities.Blog, error) {
	if err := s.BlogRepository.UpdateBlog(id, input); err != nil {
		return input, err
	}

	return input, nil
}

func (s *blogUsecase) Delete(id uuid.UUID) error {
	if err := s.BlogRepository.DeleteBlog(id); err != nil {
		return err
	}

	return nil
}

func (s *blogUsecase) FindByCat(id int) ([]entities.Blog, error) {
	if data, err := s.BlogRepository.GetBlogByCat(id); err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (s *blogUsecase) FindByKey(key string) ([]entities.Blog, error) {
	if data, err := s.BlogRepository.GetBlogByKey(key); err != nil {
		return data, err
	} else {
		return data, nil
	}
}
