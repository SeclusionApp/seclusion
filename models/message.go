package models

type Message struct {
	ID      int    `gorm:"primary_key" json:"id"`
	Content string `gorm:"not null" json:"content"`
	Time    int64  `gorm:"not null" json:"time"`
}

type Channel_Message struct {
	ChannelID int `gorm:"primary_key" json:"channel_id"`
	MessageID int `gorm:"primary_key" json:"message_id"`
}
