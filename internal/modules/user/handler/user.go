package handler

import (
	"errors"
	"strconv"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/modules/user/service"
	"github.com/ilhamosaurus/HRIS/pkg/response"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler interface {
	CreateUser(echo.Context) error
	GetUserByID(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
	GetUserByUsername(echo.Context) error
	ListUsers(echo.Context) error
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

type userHandler struct {
	userService service.UserService
}

func (h *userHandler) CreateUser(c echo.Context) error {
	authInfo := util.GetUserAuth(c)
	if authInfo.Role == types.Employee {
		return response.Unauthorized(c, "unauthorized")
	}

	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.userService.Create(c.Request().Context(), &req); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.BadRequest(c, "duplicated credentials")
		}
		return response.InternalServerError(c, err.Error())
	}
	return response.SuccessReponse(c, nil)
}

func (h *userHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid user ID")
	}

	authInfo := util.GetUserAuth(c)
	if authInfo.Role == types.Employee && authInfo.ID != id {
		return response.Unauthorized(c, "unauthorized")
	}

	user, err := h.userService.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, user)
}

func (h *userHandler) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")
	authInfo := util.GetUserAuth(c)
	if authInfo.Role == types.Employee && authInfo.Username != username {
		return response.Unauthorized(c, "unauthorized")
	}

	user, err := h.userService.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}
	return response.SuccessReponse(c, user)
}

func (h *userHandler) ListUsers(c echo.Context) error {
	authInfo := util.GetUserAuth(c)
	if authInfo.Role == types.Employee {
		return response.Unauthorized(c, "unauthorized")
	}

	var req dto.UserQuery
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	userList, err := h.userService.List(c.Request().Context(), &req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, userList)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid user ID")
	}
	authInfo := util.GetUserAuth(c)
	if authInfo.Role == types.Employee && authInfo.ID != id {
		return response.Unauthorized(c, "unauthorized")
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	req.ID = id
	if err := h.userService.Update(c.Request().Context(), &req); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.BadRequest(c, "duplicated credentials")
		}
		return response.InternalServerError(c, err.Error())
	}

	return response.SuccessReponse(c, nil)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid user ID")
	}
	authInfo := util.GetUserAuth(c)
	if authInfo.Role != types.Admin {
		return response.Unauthorized(c, "unauthorized")
	}

	if err := h.userService.Delete(c.Request().Context(), id); err != nil {
		return response.InternalServerError(c, err.Error())
	}
	return response.SuccessReponse(c, nil)
}
