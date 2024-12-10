package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"live_coding/internal/database"
	"live_coding/internal/user/routes"
	"log"
	"os"
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

	url := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Println(url)
	err := r.Run(url)
	if err != nil {
		log.Fatal("[server err]:", err)
		return
	}
}
