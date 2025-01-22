package service

import (
	"devmentor-BE103-golang/model/database"
	"devmentor-BE103-golang/repository"
)

type PostServiceInterface interface {
	FindOne(string) (*database.Post, error)
	Create(postModel database.Post) error
	FindAll() (*database.Posts, error)
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

func (s *PostService) FindOne(id string) (postModel *database.Post, err error) {
	postModel, err = s.postRepository.FindOne(id)
	return
}

func (s *PostService) Create(postModel database.Post) error {
	return s.postRepository.Create(postModel)
}

func (s *PostService) FindAll() (postModels *database.Posts, err error) {
	postModels, err = s.postRepository.FindAll()
	return
}
