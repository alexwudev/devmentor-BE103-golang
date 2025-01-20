package repository

import (
	"devmentor-BE103-golang/model/database"
)

type PostRepositoryInterface interface {
	FindAll() (*database.Posts, error)
	FindOne() (*database.Post, error)
	Create(postModel database.Post) error
}

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (post *PostRepository) FindOne() (postModel *database.Post, err error) {
	postModel = &database.Post{}
	err = postModel.Model().First(postModel).Error
	return
}

// create
func (post *PostRepository) Create(postModel database.Post) error {
	return postModel.Model().Create(&postModel).Error
}

// find all
func (post *PostRepository) FindAll() (postModels *database.Posts, err error) {
	postModels = &database.Posts{}
	err = postModels.Model().Find(postModels).Error
	return
}
