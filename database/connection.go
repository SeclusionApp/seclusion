package database

import (
	"fmt"

	"github.com/seclusionapp/seclusion/config"
	models "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/util"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// We are using a global variable to store the database connection
var DB *gorm.DB

func Connect() {

	// dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.DB_USER, config.DB_PASS, config.DB_SERVER, config.DB_PORT, config.DB_NAME)

	// github.com/denisenkom/go-mssqldb

	conn, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	util.HandleError(err, "Failed to connect to database")

	DB = conn

	conn.AutoMigrate(&models.User{}, &models.Channel{}, &models.Message{}, &models.Channel_User{})
}
