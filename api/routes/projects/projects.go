package projects

import (
	"context"
	"fmt"
	"log"
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

type newProjectRequest struct {
	Name          string `json:"name"`
	Git           string `json:"git"`
	DefaultBranch string `json:"defaultBranch"`
}

func New(c echo.Context) error {
	client := db.Connect()
	sess, _ := session.Get("session", c)
	userId, err := uuid.Parse(fmt.Sprint(sess.Values["userId"]))
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	pr := new(newProjectRequest)

	if err := c.Bind(pr); err != nil {
		return response.RespondError(c, *echo.ErrBadRequest)
	}

	project, err := client.Project.Create().
		AddUserIDs(userId).
		SetName(pr.Name).
		SetGit(pr.Git).
		SetDefaultBranch("main").
		Save(context.Background())

	log.Println(err.Error())
	return response.Respond(c, 200, *project)

}

func Get(c echo.Context) error {
	client := db.Connect()
	sess, _ := session.Get("session", c)
	userId, err := uuid.Parse(fmt.Sprint(sess.Values["userId"]))
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}
	println(userId.String())

	projects, err := client.Project.Query().Where(predicate.Project(user.IDEQ(userId))).All(context.Background())
	if err != nil {
		response.RespondError(c, *echo.ErrInternalServerError, "error querying deployments")
	}
	if len(projects) > 0 {

		var projectsLiteral []ent.Project

		for _, d := range projects {
			projectsLiteral = append(projectsLiteral, *d)
		}

		if err != nil {
			return response.RespondError(c, *echo.ErrInternalServerError)
		}

		return response.Respond(c, http.StatusOK, projectsLiteral)
	} else {
		return response.RespondError(c, *echo.ErrNotFound, "no projects found")
	}
}
