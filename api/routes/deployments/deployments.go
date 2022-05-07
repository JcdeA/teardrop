package deployments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/predicate"
	"github.com/fosshostorg/teardrop/ent/user"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/response"
	"github.com/google/uuid"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	client := db.Connect()
	sess, _ := session.Get("session", c)
	userId, err := uuid.Parse(fmt.Sprint(sess.Values["userId"]))
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	deployments, err := client.Deployment.Query().WithDomains().Where(predicate.Deployment(user.ID(userId))).All(context.Background())
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

		return response.Respond(c, http.StatusOK, deploymentsLiteral)
	} else {
		return response.RespondError(c, *echo.ErrNotFound, "no deployments found")
	}

}
