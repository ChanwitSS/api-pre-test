package models

type User struct {
	UserId int `gorm:"primary_key" json:"user_id"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
