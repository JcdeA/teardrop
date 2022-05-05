package main

import (
	"context"
	"log"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/domain"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/fosshostorg/teardrop/internal/pkg/response"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func checkDomainHandler(c echo.Context) error {
	host := c.Request().URL.Query().Get("domain")
	exists, err := domainExists(host)
	if err != nil {
		return response.RespondError(c, *echo.ErrInternalServerError)
	}

	if exists {
		return response.Respond(c, models.Response{Message: "doman linked to deployment"})
	} else {
		return response.Respond(c, models.Response{Message: "domain not linked to deployment", Status: 404})
	}

}

func domainExists(host string) (bool, error) {
	client := db.Connect()
	_, err := client.Domain.Query().Where(domain.DomainEQ(host)).First(context.Background())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func main() {
	db.Connect()

	godotenv.Load()

	e := echo.New()

	e.GET("/check", checkDomainHandler)

	log.Fatal(e.Start("0.0.0.0:3000"))
}
