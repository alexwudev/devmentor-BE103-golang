package service

import (
	"devmentor-BE103-golang/model/database"
	"devmentor-BE103-golang/repository"
)

type PostServiceInterface interface {
	Create(postModel database.Post) error
	FindAll() (*database.Posts, error)
	FindOne(post *database.Post) (*database.Post, error)
}

type PostService struct {
	postRepository repository.PostRepositoryInterface
}

func NewPostService(postRepo repository.PostRepositoryInterface) *PostService {
	res := &PostService{}
	if postRepo == nil {
		postRepo = repository.NewPostRepository()
	}
	res.postRepository = postRepo
	return res
}

func (s *PostService) FindOne(post *database.Post) (postModels *database.Post, err error) {
	return s.postRepository.FindOne(post)
}

func (s *PostService) Create(postModel database.Post) error {
	return s.postRepository.Create(postModel)
}

func (s *PostService) FindAll() (postModels *database.Posts, err error) {
	postModels, err = s.postRepository.FindAll()
	return
}
