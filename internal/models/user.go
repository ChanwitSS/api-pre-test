package models

type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// ProfileImg string `json:"profile_img"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
