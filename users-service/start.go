package users

import (
	"net/http"

	middleware "github.com/droquedev/e-commerce/pkg/middlewares"
	"github.com/droquedev/e-commerce/users-service/internal/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Use(middleware.GlobalErrorHandler())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.InitializeRoutes(r)

	r.Run(":8080")
}
