package models

import (
	"post/pkg/app"
)

type Post struct {
	PostId      int    `gorm:"primary_key" json:"post_id"`
	Status      string `json:"status"` // ["To Do", "In Profress", "Done"]
	Description string `json:"description"`
	// CreateUser

	CreatedAt int  `json:"created_at"`
	UpdatedAt int  `json:"updated_at"`
	IsArchive bool `json:"is_archive"`
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
		Order("create_date DESC").
		Find(&posts).
		Offset(offset).
		Limit(limit).
		// Preload("Comments").
		// Preload("Comments.User").
		// Preload("CreateUser").
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
		// Joins("left join order_products on order_products.order_id = orders.order_id").
		Preload("Comment").
		// Preload(clause.Associations).
		First(&post).
		Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func InsertOrder(post Post) (*Post, error) {
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
