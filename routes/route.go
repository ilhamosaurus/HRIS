package routes

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamosaurus/HRIS/handler"
	"github.com/ilhamosaurus/HRIS/middleware"
	"github.com/ilhamosaurus/HRIS/model"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	m, err := model.NewModel()
	if err != nil {
		log.Fatalf("failed to initiate database: %+v", err)
	}
	h := handler.NewHandler(util.NewHasher(setting.Server.Secret), m)

	validator := util.NewCustomValidator()
	e.Validator = validator
	e.Use(middleware.ActivityMiddleware)
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
		userRoute.GET("/info", h.GetUserInfo)
		userRoute.POST("", h.SetUser)
		userRoute.PUT("/changePassword", h.ChangePassword)
	}

	attendanceRoute := apiRoute.Group("/attendance", jwtMiddleware)
	{
		attendanceRoute.POST("/checkIn", h.CheckIn)
		attendanceRoute.POST("/checkOut", h.CheckOut)
		attendanceRoute.POST("", h.SetAttendance)
		attendanceRoute.GET("", h.GetAttendances)
	}

	overtimeRoute := apiRoute.Group("/overtime", jwtMiddleware)
	{
		overtimeRoute.POST("", h.SetOvertime)
		overtimeRoute.GET("", h.GetOvertime)
	}
}
