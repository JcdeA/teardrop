package api

import (
	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	return c.String(200, "ping!")
}
