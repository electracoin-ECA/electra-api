package server

import (
	"os"
)

// Start the server.
func Start() {
	router := Router()
	router.Run(":" + getPort())
}

func getPort() string {
	envPort := os.Getenv("PORT")

	if envPort == "" {
		return "8080"
	}

	return envPort
}
