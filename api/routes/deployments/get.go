package deployments

import (
	"fmt"

	"github.com/fosshostorg/teardrop/api/response"
	"github.com/fosshostorg/teardrop/models"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	return response.Respond(c, models.Response{Message: fmt.Sprintf("No deployments found for user %v", "placeholder@example.com")})
}
