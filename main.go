package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"live_coding/internal/database"
	"live_coding/internal/user/routes"
	"log"
)

func main() {
	startServer()
}
func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("env file not found :", envErr)
		return
	}
}

func startServer() {

	go func() {
		database.GenerateData()
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.SetUserRoutes(r)

	log.Println("http://localhost:8083")
	err := r.Run(":8083")
	if err != nil {
		log.Fatal("[server err]:", err)
		return
	}
}
