package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password []byte `gorm:"not null" json:"-"`
}
