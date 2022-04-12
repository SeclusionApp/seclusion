package util

import (
	"os"

	"github.com/seclusionapp/seclusion/config"
)

func InitEnv() {
	// Set the port os.Env from the config file
	os.Setenv("PORT", config.Port)

	// Set the JWT secret os.Env from the config file
	os.Setenv("JWT_SECRET", config.JWT_SECRET)
}
