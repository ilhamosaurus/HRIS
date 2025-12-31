package util

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamosaurus/HRIS/pkg/types"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("role", role)
	v.RegisterValidation("password", password)
	v.RegisterValidation("date", date)
	v.RegisterValidation("status", status)
	return &CustomValidator{validator: v}
}

func (c *CustomValidator) ValidationError(err error) map[string]string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		errsMap := make(map[string]string)
		for i := range errs {
			switch errs[i].ActualTag() {
			case "required":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s is %s", errs[i].Tag(), errs[i].Param())
			case "gte":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s must at least %s", errs[i].Tag(), errs[i].Param())
			case "role":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s is not a valid role", errs[i].Value())
			case "password":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s must contain at least one uppercase letter, one lowercase letter, one number, and one special character", errs[i].Field())
			case "date":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s is not a valid date, please use YYYYMMDD", errs[i].Value())
			case "status":
				errsMap[errs[i].Field()] = fmt.Sprintf("%s is not a valid status", errs[i].Value())
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

var (
	reUpper                  = regexp.MustCompile(`[A-Z]`)
	reLower                  = regexp.MustCompile(`[a-z]`)
	reNumber                 = regexp.MustCompile(`\d`)
	reSpecial                = regexp.MustCompile(`[^A-Za-z0-9]`)
	password  validator.Func = func(fl validator.FieldLevel) bool {
		pwd := fl.Field().String()

		if len(pwd) < 8 {
			return false
		}

		if !reUpper.MatchString(pwd) || !reLower.MatchString(pwd) || !reNumber.MatchString(pwd) || !reSpecial.MatchString(pwd) {
			return false
		}

		return true
	}
	dateRe                = regexp.MustCompile(`((19|20)[0-9]{2})(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])`)
	date   validator.Func = func(fl validator.FieldLevel) bool {
		return dateRe.MatchString(fl.Field().String())
	}
	status validator.Func = func(fl validator.FieldLevel) bool {
		status := types.StringToStatus(fl.Field().String())
		return status != types.Unknown_Status
	}
)

func (c *CustomValidator) Validate(in any) error {
	if err := c.validator.Struct(in); err != nil {
		errMsgs := c.ValidationError(err)
		return fmt.Errorf("%s", PrintToString(errMsgs))
	}
	return nil
}
