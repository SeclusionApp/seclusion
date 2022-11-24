package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)



const (

	// Database Information.
	DB_SERVER = "localhost"
	DB_PORT   = 3306
	DB_USER   = "postgres"
	DB_PASS   = "postgres"
	DB_NAME   = "seclusion"

	DSN = "api:password@/seclusion"

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
	AllowOrigins:     "http://localhost:3000,http://192.168.1.123/",
	AllowMethods:     "GET,POST,PUT,DELETE,HEAD,PATCH",
	AllowHeaders:     "Accept, Accept-Encoding, Authorization, Cookie, Set-Cookie, Content-Length, Content-Type, Content-Type, Host, Origin, Referer, User-Agent,Set-Cookie",
	AllowCredentials: true,
}
