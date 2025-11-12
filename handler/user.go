package handler

import (
	"net/http"

	"github.com/ilhamosaurus/HRIS/model"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	var req types.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	user := h.model.GetUserByUsername(req.Username)
	if user.Name != req.Username || !h.hasher.VerifySHAHash(req.Password, user.Password) {
		return c.JSON(types.GenerateReponse(http.StatusUnauthorized, "invalid credentials", nil))
	}

	token, err := util.GeneratoeJWTToken(user.Name, user.UserRole)
	if err != nil {
		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", token))
}

func (h *Handler) SetUser(c echo.Context) error {
	auth := util.GetUserAuth(c)

	if auth.Role != types.Admin {
		return c.JSON(types.GenerateReponse(http.StatusUnauthorized, "unauthorized", nil))
	}

	var req types.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	user := ToModelUser(req)
	if req.ID == nil {
		hashedPassword := h.hasher.GenerateSHAHash(user.Password)
		user.Password = hashedPassword
		if err := h.model.AddUser(user); err != nil {
			return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
		}
		return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
	}

	if req.Password != nil {
		hashedPassword := h.hasher.GenerateSHAHash(user.Password)
		user.Password = hashedPassword
	}
	if err := h.model.UpdateUser(user); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
}

func (h *Handler) GetUserInfo(c echo.Context) error {
	auth := util.GetUserAuth(c)
	user := h.model.GetUserByUsername(auth.Username)
	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", ToTypeUser(user)))
}

func (h *Handler) ChangePassword(c echo.Context) error {
	auth := util.GetUserAuth(c)
	var req types.ChangePasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
	}

	user := h.model.GetUserByUsername(auth.Username)
	if !h.hasher.VerifySHAHash(req.OldPassword, user.Password) {
		return c.JSON(types.GenerateReponse(http.StatusUnauthorized, "invalid credentials", nil))
	}

	hashedPassword := h.hasher.GenerateSHAHash(req.NewPassword)
	user.Password = hashedPassword
	if err := h.model.UpdateUser(user); err != nil {
		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
}

func ToModelUser(req types.User) model.User {
	user := model.User{
		Name:     req.Username,
		UserRole: req.Role,
		Salary:   req.GetSalary(),
		Active:   req.IsActive(),
	}

	if req.Password != nil {
		user.Password = req.GetPassword()
	}

	if req.ID != nil {
		user.ID = req.GetID()
	}

	return user
}

func ToTypeUser(user model.User) *types.User {
	return &types.User{
		ID:       &user.ID,
		Username: user.Name,
		Role:     user.UserRole,
		Salary:   &user.Salary,
		Active:   &user.Active,
	}
}
