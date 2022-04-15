package response

import (
	"fmt"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/labstack/echo/v4"
)

func Respond(c echo.Context, resp models.Response) error {
	return c.JSON(resp.Status, resp.ToMap())
}

func RespondError(c echo.Context, status echo.HTTPError) error {
	return Respond(c, models.Response{
		Status:           status.Code,
		Message:          fmt.Sprintf("%v", status.Message),
		DocumentationUrl: "https://example.com"})
}
