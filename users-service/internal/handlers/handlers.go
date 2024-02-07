package handlers

import (
	"net/http"

	"github.com/droquedev/e-commerce/users-service/internal/dto"
	"github.com/droquedev/e-commerce/users-service/internal/entities"
	"github.com/droquedev/e-commerce/users-service/internal/use_cases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase *use_cases.UserUseCases
}

func NewUserHandler(userUseCase *use_cases.UserUseCases) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var createUserDTO dto.UserCreateDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entities.User{
		ID:       uuid.NewString(),
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

	err2 := h.userUseCase.UserCreator(user)

	if err2 != nil {
		if err2.Error() == "username or email already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"result":  user,
	})
}
