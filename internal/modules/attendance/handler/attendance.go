package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/modules/attendance/service"
	"github.com/ilhamosaurus/HRIS/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AttendanceHandler interface {
	CheckIn(echo.Context) error
	CheckOut(echo.Context) error
	GetAttendances(echo.Context) error
	Delete(echo.Context) error
}

func NewAttendanceHandler(attendanceService service.AttendanceService) AttendanceHandler {
	return &attendanceHandler{attendanceService: attendanceService}
}

type attendanceHandler struct {
	attendanceService service.AttendanceService
}

func (h *attendanceHandler) CheckIn(c echo.Context) error {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if err := h.attendanceService.CheckIn(c.Request().Context(), date, now); err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, nil)
}

func (h *attendanceHandler) CheckOut(c echo.Context) error {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if err := h.attendanceService.CheckOut(c.Request().Context(), date, now); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.BadRequest(c, "attendance user have not checked in")
		}
		return response.InternalServerError(c, err.Error())
	}
	return response.SuccessReponse(c, nil)
}

func (h *attendanceHandler) GetAttendances(c echo.Context) error {
	var req dto.AttendancesQuery
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	attendances, total, err := h.attendanceService.GetAttendances(c.Request().Context(), &req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, &dto.AttendanceResponse{
		List:     attendances,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

func (h *attendanceHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid attendance ID")
	}

	if err := h.attendanceService.Delete(c.Request().Context(), id); err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, nil)
}
