package service

import (
	"log"
	"net/http"

	"securityMS/cmd/utils"
	"securityMS/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
)

type AuthService interface {
	VerifyAuth(c *gin.Context)
}

type AuthDBHandler struct {
	repo repository.JWTInterface
}

func NewAuthService(repo *repository.JWTInterface) AuthService {
	return &AuthDBHandler{repo: *repo}
}

func (a *AuthDBHandler) VerifyAuth(c *gin.Context) {

	cookieString, _ := c.Cookie("Authorization")
	cookie := fiber.Cookie{
		Name:  "Authorization",
		Value: cookieString,
	}
	token, err := a.repo.VerifyToken(&cookie)
	log.Println(token)
	if err == nil {
		c.JSON(http.StatusOK, "Valid Token")
	} else {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.BadRequestError("invalid token", err))
	}
}
