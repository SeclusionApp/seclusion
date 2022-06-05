package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null" json:"email"`
	Password []byte `gorm:"not null" json:"-"` // Security MAX!!!
}

type Favorited_User struct {
	UserID           int `gorm:"primary_key" json:"user_id"`
	Favorited_UserID int `gorm:"primary_key" json:"favorited_user_id"`
}
