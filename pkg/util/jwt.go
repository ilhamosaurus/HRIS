package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaim struct {
	Username string     `json:"username"`
	Role     types.Role `json:"role"`
	jwt.RegisteredClaims
}

func GeneratoeJWTToken(username, role string) (string, error) {
	claims := JWTCustomClaim{
		Username: username,
		Role:     types.StringToRole(role),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "HRIS",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(setting.Server.JWTSecret))
}

func GetUserAuth(c echo.Context) *JWTCustomClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaim)
	return claims
}
