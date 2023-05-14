package models

import (
	"post/internal/enums"
	"post/pkg/app"
	"time"
)

type Post struct {
	PostId      int              `gorm:"primary_key" json:"post_id"`
	Title       string           `json:"title"`
	Status      enums.PostStatus `json:"status"` // ["To Do", "In Profress", "Done"]
	Description string           `json:"description"`

	UserId int `json:"user_id,omitempty"`

	Comment []Comment `gorm:"constraint:OnDelete:CASCADE;" json:"comment"`
	User    *User     `gorm:"references:user_id" json:"user,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsArchive bool      `json:"is_archive"`
}

type QueryPost struct {
	app.Query
}

func FindPosts(query QueryPost) (*[]Post, error) {
	var (
		posts  []Post
		limit  = query.Take
		offset = (query.Page - 1) * query.Take
	)

	err := db.
		Table("posts").
		Order("created_at DESC").
		Find(&posts).
		Offset(offset).
		Limit(limit).
		Preload("Comments").
		// Preload("Comments.User").
		Preload("User").
		Error

	if err != nil {
		return nil, err
	}

	return &posts, nil
}

func FindPost(postId string) (*Post, error) {
	var post Post
	err := db.
		Table("posts").
		Where("posts.post_id = ?", postId).
		Preload("Comment").
		Preload("User").
		// Joins("left join order_products on order_products.order_id = orders.order_id").
		// Preload(clause.Associations).
		First(&post).
		Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func CreatePost(post Post) (*Post, error) {
	if err := db.Table("posts").Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func UpdatePost(postId string, post Post) (*Post, error) {
	if err := db.Table("posts").Where("post_id = ?", postId).Updates(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func DeleteOrder(id string) error {
	if err := db.Table("posts").Where("post_id = ?", id).Delete(&Post{}).Error; err != nil {
		return err
	}

	return nil
}
