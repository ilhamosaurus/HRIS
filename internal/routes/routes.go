package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamosaurus/HRIS/internal/container"
	userroute "github.com/ilhamosaurus/HRIS/internal/modules/user/route"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	container *container.Container
}

func NewRoutes(container *container.Container) *Routes {
	return &Routes{
		container: container,
	}
}

func (r *Routes) SetupRoutes(e *echo.Echo) {
	validator := util.NewCustomValidator()
	e.Validator = validator
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method = ${method} | URL = \"${uri}\"| Status = ${status} | Latency = ${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Recover())

	apiRoute := e.Group("/api")
	apiRoute.POST("/login", r.container.AuthHandler.Login)

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JWTCustomClaim)
		},
		SigningKey: []byte(setting.Server.JWTSecret),
	})

	apiRoute.Use(jwtMiddleware)

	userRoute := userroute.NewUserRoute(r.container.UserHandler)
	userRoute.RegisterRoutes(apiRoute)
}
