package models

type Comment struct {
	CommentId int `gorm:"primary_key" json:"comment_id"`
	PostId    int `json:"post_id"`
	UserId    int `json:"user_id"`

	Text string `json:"text"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
