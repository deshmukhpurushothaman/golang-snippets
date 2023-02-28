package routes

import (
	controller "github.com/deshmukhpurushothaman/golang-snippets/authentication/controllers"
	"github.com/deshmukhpurushothaman/golang-snippets/authentication/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
