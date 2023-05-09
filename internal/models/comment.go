package models

type Comment struct {
	CommentId int `gorm:"primary_key" json:"comment_id"`
	PostId    int `gorm:"primary_key" json:"comment_id"`
	UserId    int `gorm:"primary_key" json:"comment_id"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
