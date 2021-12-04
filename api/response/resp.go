package response

import (
	"github.com/fosshostorg/teardrop/models"
	"github.com/labstack/echo/v4"
)

func Respond(c echo.Context, resp models.Response) error {
	return c.JSON(resp.Status, resp.ToMap())
}

func RespondUnauthorized(c echo.Context) error {
	return Respond(c, models.Response{
		Status:           echo.ErrUnauthorized.Code,
		Message:          "unauthorized",
		DocumentationUrl: "https://example.com"})
}
