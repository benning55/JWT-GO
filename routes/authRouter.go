package routes

import (
	"github.com/benning55/golang-jwt-mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("user/signup", controllers.Signup())
	incommingRoutes.POST("user/login", controllers.Login())
}
