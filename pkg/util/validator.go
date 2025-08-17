package util

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("role", role)
	return &CustomValidator{validator: v}
}

func (c *CustomValidator) ValidationError(err error) map[string]string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		errsMap := make(map[string]string)
		for i := range errs {
			switch errs[i].ActualTag() {
			default:
				errsMap[errs[i].Field()] = fmt.Sprintf("%s, %s", errs[i].Tag(), errs[i].Param())
			}
		}

		return errsMap
	}
	return nil
}

var role validator.Func = func(fl validator.FieldLevel) bool {
	role := types.StringToRole(fl.Field().String())
	return role != types.Unknown_Role
}

func (c *CustomValidator) Validate(in any) error {
	if err := c.validator.Struct(in); err != nil {
		errMsgs := c.ValidationError(err)
		return echo.NewHTTPError(http.StatusBadRequest, errMsgs)
	}
	return nil
}
