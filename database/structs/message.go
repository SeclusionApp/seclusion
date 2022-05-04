package models

type Message struct {
	ID        int    `gorm:"primary_key" json:"id"`
	UserID    int    `json:"user_id"`
	ChannelID int    `json:"channel_id"`
	Content   string `gorm:"not null" json:"content"`
	Time      int64  `gorm:"not null" json:"time"`
}
