package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"live_coding/entity"
)

type UserRepositoryInterface interface {
	GetUserBy(c *gin.Context, ID uuid.UUID) (entity.User, error)
}