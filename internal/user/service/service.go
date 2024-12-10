package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	repo "live_coding/internal/user/repository"
	"live_coding/internal/user/responses"
	"net/http"
)

type UserService struct {
	repo repo.UserRepositoryInterface
}

func NewUserService(repo repo.UserRepositoryInterface) UserService {
	return UserService{repo: repo}
}

func (s UserService) GetUserByID(c *gin.Context) {

	//check uuid
	UUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	//get user from db
	userData, err := s.repo.GetUserBy(c, UUID)
	if err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"error": "record not found",
			},
		)
		return
	}

	//pass user
	c.JSON(http.StatusOK, gin.H{
		"data": responses.ToUser(userData),
	})
}
