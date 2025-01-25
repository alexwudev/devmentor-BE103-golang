package repository

import (
	"devmentor-BE103-golang/model/database"
)

type PostRepositoryInterface interface {
	Create(postModel database.Post) error
	FindAll() (*database.Posts, error)
	FindOne(*database.Post) (*database.Post, error)
	UpdateOne(string, *database.Post) (*database.Post, error)
}

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// find one by id
func (post *PostRepository) FindOne(postd *database.Post) (*database.Post, error) {
	var postModel database.Post
	err := postModel.Model().Where(postd.Id).First(&postModel).Error
	if err != nil {
		return nil, err
	}
	return &postModel, nil
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

func (post *PostRepository) UpdateOne(id string, postd *database.Post) (*database.Post, error) {
	var postModel database.Post
	err := postModel.Model().Where(id).UpdateColumns(&postModel).Error
	if err != nil {
		return nil, err
	}
	return &postModel, nil
}
