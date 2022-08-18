package routes

import (
	"securityMS/pkg/repository"
	"securityMS/pkg/service"
	"securityMS/projectLib/cqlWrapper"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeLogin(router gin.Engine) {

	session := cqlWrapper.CreateCassandraSession("sec")

	jwtR := repository.NewJWT()
	loginR := repository.NewloginCrd(session)

	loginS := service.NewLoginService(jwtR, loginR)

	router.POST("/login", func(c *gin.Context) {
		cookie := loginS.Login(c)
		if cookie != nil {
			c.SetCookie(cookie.Name, cookie.Value, 2000, cookie.Path, "SecurityMicroS", false, false)
			c.JSON(http.StatusOK, "success")
		} else {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

	})
}
