package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"live_coding/internal/user/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}
func (u UserRepository) GetUserBy(c *gin.Context, ID uuid.UUID) (entity.User, error) {
	var user entity.User

	if err := u.db.Preload("Addresses").Find(&user, ID).Error; err != nil {
		return user, err
	}
	return user, nil
}
