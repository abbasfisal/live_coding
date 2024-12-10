package routes

import (
	"github.com/gin-gonic/gin"
	"live_coding/internal/database"
	"live_coding/internal/user/repository"
	"live_coding/internal/user/service"
)

func SetUserRoutes(r *gin.Engine) {

	userRepo := repository.NewUserRepository(database.Get())
	userSrv := service.NewUserService(userRepo)

	//-------------
	// ROUTES
	//-------------
	r.GET("/users/:id", userSrv.GetUserByID)

}
