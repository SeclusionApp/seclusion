package util

import (
	"os"
	"time"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetTime() int64 {
	return time.Now().Unix()
}
