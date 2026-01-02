package route

import (
	"github.com/ilhamosaurus/HRIS/internal/modules/attendance/handler"
	"github.com/labstack/echo/v4"
)

type AttendanceRoute struct {
	attendanceHandler handler.AttendanceHandler
}

func (r *AttendanceRoute) RegisterRoutes(group *echo.Group) {
	attendanceGroup := group.Group("/attendances")
	{
		attendanceGroup.POST("/check-in", r.attendanceHandler.CheckIn)
		attendanceGroup.POST("/check-out", r.attendanceHandler.CheckOut)
		attendanceGroup.GET("", r.attendanceHandler.GetAttendances)
		attendanceGroup.DELETE("/:id", r.attendanceHandler.Delete)
	}
}

func NewAttendanceRoute(attendanceHandler handler.AttendanceHandler) *AttendanceRoute {
	return &AttendanceRoute{attendanceHandler: attendanceHandler}
}
