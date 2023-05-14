package models

import "time"

type Comment struct {
	CommentId int `gorm:"primary_key" json:"comment_id"`
	PostId    int `json:"post_id"`
	UserId    int `json:"user_id"`

	Text string `json:"text"`

	User *User `gorm:"references:user_id" json:"user,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateComment(comment Comment) (*Comment, error) {
	if err := db.Table("comments").Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}
