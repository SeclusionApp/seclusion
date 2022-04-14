package models

type Message struct {
	ID        int    `gorm:"primary_key" json:"id"`
	ChannelID int    `gorm:"not null" json:"channel_id"`
	UserID    int    `gorm:"not null" json:"user_id"`
	Content   string `gorm:"not null" json:"content"`
	Time      int64  `gorm:"not null" json:"time"`
}
