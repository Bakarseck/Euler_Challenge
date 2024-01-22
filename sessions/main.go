package main

import (
	"os"
	"sessions/internals/models/server"
	"sessions/internals/utils"
)

func main() {
	utils.LoadEnv(".env")

	PORT := os.Getenv("PORT")

	router := server.NewServer()

	router.StartServer(PORT)
}
