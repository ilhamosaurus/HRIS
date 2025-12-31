package route

import (
	"github.com/ilhamosaurus/HRIS/internal/modules/user/handler"
	"github.com/labstack/echo/v4"
)

type UserRoute struct {
	userHandler handler.UserHandler
}

func (r UserRoute) RegisterRoutes(group *echo.Group) {
	userGroup := group.Group("/users")
	userGroup.POST("", r.userHandler.CreateUser)
	userGroup.GET("", r.userHandler.ListUsers)
	userGroup.GET("/:id", r.userHandler.GetUserByID)
	userGroup.PUT("/:id", r.userHandler.UpdateUser)
	userGroup.DELETE("/:id", r.userHandler.DeleteUser)
	userGroup.GET("/:username", r.userHandler.GetUserByUsername)
}

func NewUserRoute(userHandler handler.UserHandler) *UserRoute {
	return &UserRoute{userHandler: userHandler}
}
