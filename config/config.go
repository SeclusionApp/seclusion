package config

const (
	// Port is the port to run the server on
	PORT = "8080"

	// JWT_SECRET is the secret used to sign the JWT
	JWT_SECRET = "sakdjnasmcavkadjvnackjaw!IU!HJ!U@EIKNJ"

	// JWT_EXPIRY is the expiry time for the JWT (seconds)
	JWT_EXPIRY = 86400

	//Logger Config:
	LOGGER_FORMAT      = "[${time}] ${status} ${pid} ${latency} ${locals:requestid} ${ip}:${port} - ${method} ${path}\n"
	LOGGER_TIME_FORMAT = "15:04:05"
	LOGGER_TIME_ZONE   = "Local"
	LOGGER_OUTPUT      = "./log.txt"
)
