package repository

import (
	"devmentor-BE103-golang/model/database"
)

type PostRepositoryInterface interface {
	Create(postModel database.Post) error
	FindAll() (*database.Posts, error)
	FindOne(map[string]interface{}) (*database.Post, error)
}

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (post *PostRepository) FindOne(conditions map[string]interface{}) (*database.Post, error) {
	var postModel database.Post
	err := postModel.Model().Where(conditions).First(&postModel).Error
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
