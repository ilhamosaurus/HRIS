package handler

import (
	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/modules/user/service"
	"github.com/ilhamosaurus/HRIS/pkg/response"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

type authHandler struct {
	authService service.AuthService
}

func (h *authHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	token, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return response.Unauthorized(c, err.Error())
		}
	}
	return response.SuccessReponse(c, token)
}
