package response

import "github.com/labstack/echo/v4"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessReponse(c echo.Context, data any) error {
	return c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

func BadRequest(c echo.Context, message string) error {
	return c.JSON(400, Response{
		Code:    400,
		Message: message,
	})
}

func Unauthorized(c echo.Context, message string) error {
	return c.JSON(401, Response{
		Code:    401,
		Message: message,
	})
}

func InternalServerError(c echo.Context, message string) error {
	return c.JSON(500, Response{
		Code:    500,
		Message: message,
	})
}
