package response

import "github.com/labstack/echo/v4"

func Error(c echo.Context, status int, err error) error {
	return c.JSON(status, map[string]any{
		"error": err.Error(),
	})
}
