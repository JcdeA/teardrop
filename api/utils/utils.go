package utils

import (
	"errors"
	"strings"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(200, "ping!")
}

func AuthenticateUser(c echo.Context) (models.SessionUser, error) {
	sess, _ := session.Get("session", c)
	user, ok := sess.Values["user"].(models.SessionUser)
	if !ok {
		return models.SessionUser{}, errors.New("could not type cast user in session to models.SessionUser - value is probably nil")
	}

	return user, nil
}

func ParseIncludeQuery(c echo.Context, value string) bool {
	include := c.QueryParam("include")
	return strings.Contains(include, value)
}
