package repository

import (
	"os"
	"securityMS/projectLib/httpWrapper"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type JWTInterface interface {
	GenerateToken(name string, tenantid int) string
	VerifyToken(cookie *fiber.Cookie) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Email    string `json:"email"`
	TenantId int    `json:"tenant_id"`
	jwt.StandardClaims
}

type jwtStruct struct {
	secretKey string
	issuer    string
}

func NewJWT() JWTInterface {
	return &jwtStruct{
		secretKey: getSecretKey(),
		issuer:    "test",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (jwtR *jwtStruct) GenerateToken(email string, tenantid int) string {
	claims := &jwtCustomClaims{
		email,
		tenantid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    jwtR.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtR.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtR *jwtStruct) VerifyToken(cookie *fiber.Cookie) (*jwt.Token, error) {
	token, err := httpWrapper.VerifyAuthZ(cookie)
	return token, err
}
