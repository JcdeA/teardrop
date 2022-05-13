package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(200, "ping!")
}

func AuthenticateUser(c echo.Context) (models.Session, error) {
	sess, _ := session.Get("session", c)
	log.Println(sess.Values["session"])
	sessionData, ok := sess.Values["session"].(models.Session)
	if !ok {
		return models.Session{}, errors.New("could not type cast user in session to models.Session - value is probably nil")
	}

	return sessionData, nil
}

func ParseIncludeQuery(c echo.Context, value string) bool {
	include := c.QueryParam("include")
	return strings.Contains(include, value)
}
