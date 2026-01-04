package route

import (
	"github.com/ilhamosaurus/HRIS/internal/modules/overtime/handler"
	"github.com/labstack/echo/v4"
)

type OvertimeRoute struct {
	overtimeHandler handler.OvertimeHandler
}

func NewOvertimeRoute(overtimeHandler handler.OvertimeHandler) *OvertimeRoute {
	return &OvertimeRoute{
		overtimeHandler: overtimeHandler,
	}
}

func (r *OvertimeRoute) RegisterRoute(group *echo.Group) {
	overtimeGroup := group.Group("/overtimes")
	{
		overtimeGroup.GET("", r.overtimeHandler.GetOvertimes)
		overtimeGroup.POST("", r.overtimeHandler.Create)
		overtimeGroup.GET("/:id", r.overtimeHandler.GetByID)
		overtimeGroup.PUT("/:id", r.overtimeHandler.Update)
		overtimeGroup.DELETE("/:id", r.overtimeHandler.Delete)
	}

	overtimeApprovalGroup := group.Group("/overtime-approvals")
	{
		overtimeApprovalGroup.PUT("/:id", r.overtimeHandler.Submit)
		overtimeApprovalGroup.POST("/:id", r.overtimeHandler.ProcessApproval)
	}
}
