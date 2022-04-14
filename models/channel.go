package models

type Channel struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"unique, not null" json:"name"`
}

type Channel_User struct {
	ChannelID int `gorm:"primary_key" json:"channel_id"`
	UserID    int `gorm:"primary_key" json:"user_id"`
}
