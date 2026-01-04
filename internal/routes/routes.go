package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamosaurus/HRIS/internal/container"
	customiddleware "github.com/ilhamosaurus/HRIS/internal/middleware"
	attendanceroute "github.com/ilhamosaurus/HRIS/internal/modules/attendance/route"
	overtimeroute "github.com/ilhamosaurus/HRIS/internal/modules/overtime/route"
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
	customMiddleware := customiddleware.NewCustomMiddleware(r.container.UserActivityDAO)
	e.Use(customMiddleware.ActivityMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method = ${method} | URL = \"${uri}\"| Status = ${status} | Latency = ${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Recover())

	apiRoute := e.Group("/api/v1")
	apiRoute.POST("/login", r.container.AuthHandler.Login)

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JWTCustomClaim)
		},
		SigningKey: []byte(setting.Server.JWTSecret),
	})

	apiRoute.Use(jwtMiddleware)
	apiRoute.Use(customMiddleware.AuthMiddeware)

	userRoute := userroute.NewUserRoute(r.container.UserHandler)
	userRoute.RegisterRoutes(apiRoute)

	attendanceRoute := attendanceroute.NewAttendanceRoute(r.container.AttendanceHandler)
	attendanceRoute.RegisterRoutes(apiRoute)

	overtimeRoute := overtimeroute.NewOvertimeRoute(r.container.OvertimeHandler)
	overtimeRoute.RegisterRoute(apiRoute)
}
