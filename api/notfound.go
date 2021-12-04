package api

import (
	"github.com/fosshostorg/teardrop/api/response"
	"github.com/fosshostorg/teardrop/models"
	"github.com/labstack/echo/v4"
)

func notfound(c echo.Context) error {
	return response.Respond(c, models.Response{Status: 404, Message: "not found", DocumentationUrl: "https://example.com"})
}
