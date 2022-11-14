package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var server = "seclusion-us-east1.database.windows.net"
var port = 1433
var user = "seclusion"
var password = "<your_password>"
var database = "seclusion"

const (

	// Database Information.
	DB_SERVER = "seclusion-us-east1.database.windows.net"
	DB_PORT   = 1433
	DB_USER   = "seclusion"
	DB_PASS   = "L0laL0la123!?$"
	DB_NAME   = "seclusion"

	// Port is the port to run the server on
	PORT = ":8080"

	// JWT_SECRET is the secret used to sign the JWT
	JWT_SECRET = "sakdjnasmcavkadjvnackjaw!IU!HJ!U@EIKNJ"

	// JWT_EXPIRY is the expiry time for the JWT (seconds)
	JWT_EXPIRY = 86400

	//Logger Config:
	LOGGER_FORMAT      = "[${time}] ${status} ${pid} ${latency} ${locals:requestid} ${ip}:${port} - ${method} ${path}\n"
	LOGGER_TIME_FORMAT = "15:04:05"
	LOGGER_TIME_ZONE   = "Local"
	LOGGER_OUTPUT      = "./logs/seclusion.log"

	MAX_REQUESTS = 30
)

var LOGGER = &logger.Config{
	Format:     LOGGER_FORMAT,
	TimeFormat: LOGGER_TIME_FORMAT,
	TimeZone:   LOGGER_TIME_ZONE,
	Output:     os.Stdout,
}

var CORS = &cors.Config{
	AllowOrigins:     "*",
	AllowMethods:     "GET,POST,PUT,DELETE,HEAD,PATCH",
	AllowHeaders:     "Accept,Accept-Encoding,Authorization,Cookie,Content-Length,Content-Type,Content-Type,Host,Origin,Referer,User-Agent",
	AllowCredentials: false,
	MaxAge:           3600,
}
