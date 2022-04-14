package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/seclusionapp/seclusion/models"
	"github.com/seclusionapp/seclusion/util"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("test:password@/seclusion"), &gorm.Config{})
	util.HandleError(err, "Failed to connect to database")

	DB = conn

	conn.AutoMigrate(&models.User{}, &models.Channel{}, &models.Message{}, &models.Channel_User{})
}
