package httpWrapper

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
)

//takes a cookie as a parameter, verify token validity and returns it
func VerifyAuthZ(cookie *fiber.Cookie) (*jwt.Token, error) {
	SecretKey := os.Getenv("JWT_KEY")
	if SecretKey == "" {
		SecretKey = "secret"
	}
	tokenString := cookie.Value
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
}

//sends and http request to the securityMS to verify the token
func CheckAuth(c *gin.Context) (bool, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	cookie, _ := c.Cookie("Authorization")
	cookieString := "Authorization=" + cookie
	log.Println(cookie)
	req, _ := http.NewRequest("POST", "http://SecurityMicroS:8080/verifytoken", nil)
	req.Header.Add("Cookie", cookieString)
	resp, err := client.Do(req)
	return resp.StatusCode == http.StatusOK, err
}

func VerifyCache(c *gin.Context, path string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", "http://CacheMS:8082/cache"+path, nil)
	cookie, _ := c.Cookie("Authorization")
	cookieString := "Authorization=" + cookie
	req.Header.Add("Cookie", cookieString)
	resp, err := client.Do(req)
	return resp, err
}
