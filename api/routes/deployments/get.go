package deployments

import (
	"context"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/fosshostorg/teardrop/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	client := db.Connect()

	deployments, err := client.Deployment.Query().WithDomains().All(context.Background())
	if err != nil {
		response.RespondError(c, *echo.ErrInternalServerError, "error querying deployments")
	}

	if len(deployments) > 0 {

		var deploymentsLiteral []ent.Deployment

		for _, d := range deployments {
			deploymentsLiteral = append(deploymentsLiteral, *d)
		}

		if err != nil {
			return response.RespondError(c, *echo.ErrInternalServerError)
		}

		return response.Respond(c, models.Response{
			Message: "successfully queried all deployments",
			Data:    deploymentsLiteral,
		})
	} else {
		return response.RespondError(c, *echo.ErrNotFound, "no deployments found")
	}

}
