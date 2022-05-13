package auth

import (
	"github.com/fosshostorg/teardrop/api/utils"
	"github.com/fosshostorg/teardrop/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

func GetSession(c echo.Context) error {
	user, err := utils.AuthenticateUser(c)

	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	return response.Respond(c, 200, user)
}
