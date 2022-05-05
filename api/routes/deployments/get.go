package deployments

import (
	"context"
	"encoding/json"

	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/fosshostorg/teardrop/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	client := db.Connect()

	deployments, err := client.Deployment.Query().All(context.Background())
	if err != nil {
		response.RespondError(c, *echo.ErrInternalServerError, "error querying deployments")
	}

	if len(deployments) > 0 {

		data, err := json.Marshal(deployments)
		if err != nil {
			return response.RespondError(c, *echo.ErrInternalServerError)
		}

		return response.Respond(c, models.Response{
			Message: "successfully queried all deployments",
			Data:    data,
		})
	} else {
		return response.RespondError(c, *echo.ErrNotFound, "no deployments found")
	}

}
