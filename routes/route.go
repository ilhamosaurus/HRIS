package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamosaurus/HRIS/handler"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	h := handler.NewHandler(util.NewHasher(setting.Server.Secret))

	validator := util.NewCustomValidator()
	e.Validator = validator
	apiRoute := e.Group("/api")

	apiRoute.POST("/login", h.Login)

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JWTCustomClaim)
		},
		SigningKey: []byte(setting.Server.JWTSecret),
	})
	userRoute := apiRoute.Group("/user", jwtMiddleware)
	{
		userRoute.GET("/getUserInfo", h.GetUserInfo)
	}
}
