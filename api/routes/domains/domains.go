package domains

import (
	"context"
	"fmt"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/fosshostorg/teardrop/internal/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type newDomainRequest struct {
	DeploymentId uuid.UUID `json:"deploymentID"`
	Domain       string    `json:"domain"`
}

func Add(c echo.Context) error {
	client := db.Connect()
	dr := new(newDomainRequest)

	if err := c.Bind(dr); err != nil {
		return response.RespondError(c, *echo.ErrBadRequest)
	}

	deployment, err := client.Deployment.Get(context.Background(), dr.DeploymentId)
	switch err.(type) {
	case *ent.NotFoundError:
		return response.RespondError(c, *echo.ErrNotFound)
	default:
		response.RespondError(c, *echo.ErrInternalServerError)
	}

	err = client.Domain.Create().SetDeployment(deployment).SetDomain(dr.Domain).Exec(context.Background())
	if err != nil {
		response.RespondError(c, *echo.ErrInternalServerError, fmt.Sprintf("error: %v", err.Error()))
	}

	return response.Respond(c, 200, models.ErrorResponse{
		Message: "successfully created domain",
	})

}
