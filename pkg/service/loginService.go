package service

import (
	"log"
	"securityMS/pkg/models"
	"securityMS/pkg/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
)

type LoginService interface {
	Login(c *gin.Context) *fiber.Cookie
}

type loginService struct {
	jwtSrv   repository.JWTInterface
	loginSrv repository.LoginInfo
}

func NewLoginService(jwtSrv repository.JWTInterface, loginSrv repository.LoginInfo) LoginService {
	return &loginService{
		jwtSrv:   jwtSrv,
		loginSrv: loginSrv,
	}
}

func (l *loginService) Login(c *gin.Context) *fiber.Cookie {
	var credentials models.User
	err := c.ShouldBind(&credentials)
	if err != nil {
		return nil
	}
	isAuthenticated, tenant := l.loginSrv.VerfiyCredentials(credentials.Email, credentials.Password)
	log.Println("auth")
	log.Println(isAuthenticated)
	if isAuthenticated {
		cookie := fiber.Cookie{
			Name:    "Authorization",
			Value:   l.jwtSrv.GenerateToken(credentials.Email, tenant),
			Expires: time.Now().Add(time.Minute * 30),
			Secure:  false,
			Domain:  "SecurityMicroS",
		}
		c.Cookie("Authorization")
		log.Println(cookie)
		return &cookie
	}

	return nil
}
