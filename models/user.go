package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique, not null" json:"username"`
	Password []byte `gorm:"not null" json:"-"`
}
