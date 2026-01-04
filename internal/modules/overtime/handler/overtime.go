package handler

import (
	"strconv"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/modules/overtime/service"
	"github.com/ilhamosaurus/HRIS/pkg/response"
	"github.com/labstack/echo/v4"
)

type OvertimeHandler interface {
	Create(echo.Context) error
	GetOvertimes(echo.Context) error
	Update(echo.Context) error
	GetByID(echo.Context) error
	Delete(echo.Context) error
	Submit(echo.Context) error
	ProcessApproval(echo.Context) error
}

func NewOvertimeHandler(overtimeService service.OvertimeService) OvertimeHandler {
	return &overtimeHandler{
		overtimeService: overtimeService,
	}
}

type overtimeHandler struct {
	overtimeService service.OvertimeService
}

func (h *overtimeHandler) Create(c echo.Context) error {
	var req dto.CreateOvertimeRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.ValidationFailed(c, err)
	}

	overtime, err := h.overtimeService.Create(c.Request().Context(), &req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtime)
}

func (h *overtimeHandler) GetOvertimes(c echo.Context) error {
	var req dto.OvertimeQuery
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.ValidationFailed(c, err)
	}

	overtimes, err := h.overtimeService.GetOvertimes(c.Request().Context(), &req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtimes)
}

func (h *overtimeHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var req dto.UpdateOvertimeRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.ValidationFailed(c, err)
	}

	overtime, err := h.overtimeService.Update(c.Request().Context(), id, &req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtime)
}

func (h *overtimeHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	overtime, err := h.overtimeService.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtime)
}

func (h *overtimeHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.overtimeService.Delete(c.Request().Context(), id); err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, nil)
}

func (h *overtimeHandler) Submit(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	overtime, err := h.overtimeService.Submit(c.Request().Context(), id)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtime)
}

func (h *overtimeHandler) ProcessApproval(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var req dto.ApprovalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.ValidationFailed(c, err)
	}

	overtime, err := h.overtimeService.ProcessApproval(c.Request().Context(), id, &req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, overtime)
}
