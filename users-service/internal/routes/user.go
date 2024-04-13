package routes

import (
	"github.com/droquedev/e-commerce/users-service/internal/handlers"
	"github.com/droquedev/e-commerce/users-service/internal/repository"
	"github.com/droquedev/e-commerce/users-service/internal/use_cases"
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine) {

	userRepository := repository.NewUserPostgresRepository()

	userUseCases := use_cases.NewUserUseCases(userRepository)

	userHandler := handlers.NewUserHandler(userUseCases)

	productGroup := router.Group("/api/users")
	{
		productGroup.POST("/", userHandler.CreateUserHandler)
	}
}
