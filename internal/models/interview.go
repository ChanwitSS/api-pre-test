package models

import (
	"post/pkg/app"
)

type Post struct {
	PostId int `gorm:"primary_key" json:"post_id"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
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
