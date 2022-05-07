package response

import (
	"fmt"
	"strings"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/labstack/echo/v4"
)

func Respond(c echo.Context, status int, data interface{}) error {

	return c.JSON(status, data)
}

func RespondError(c echo.Context, status echo.HTTPError, msg ...string) error {
	message := ""

	if len(msg) == 0 {
		message = fmt.Sprint(status.Message)
	} else {
		message = strings.Join(msg, " ")
	}

	return Respond(c, status.Code, models.ErrorResponse{
		Status:           status.Code,
		Message:          message,
		DocumentationUrl: "https://example.com"})
}
