package deployments

import (
	"github.com/fosshostorg/teardrop/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	return response.RespondError(c, *echo.ErrNotFound)
}
