package database

import (
	"github.com/seclusionapp/seclusion/config"
	models "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// We are using a global variable to store the database connection
var DB *gorm.DB

func Connect() {

	conn, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	util.HandleError(err, "Failed to connect to database")

	DB = conn

	conn.AutoMigrate(&models.User{}, &models.Channel{}, &models.Message{}, &models.Channel_User{})
}
