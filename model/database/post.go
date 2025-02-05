package database

import (
	"devmentor-BE103-golang/infrastructure"
	"time"

	"gorm.io/gorm"
)

// Post model
type Post struct {
	Id        int       `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"column:title"`
	Content   string    `json:"content" gorm:"column:content"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) Model() *gorm.DB { return infrastructure.Db.Model(post) }

type Posts []Post

func (posts *Posts) Model() *gorm.DB { return infrastructure.Db.Model(posts) }
