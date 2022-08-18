package routes

import (
	"securityMS/pkg/repository"
	"securityMS/pkg/service"

	"github.com/gin-gonic/gin"
)

func InitializeAuth(router gin.Engine) {

	jwtR := repository.NewJWT()
	authS := service.NewAuthService(&jwtR)

	router.POST("/verifytoken", authS.VerifyAuth)

}
