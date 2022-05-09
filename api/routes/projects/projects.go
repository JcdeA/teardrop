package projects

import (
	"context"
	"net/http"

	"github.com/fosshostorg/teardrop/api/utils"
	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/project"
	"github.com/fosshostorg/teardrop/ent/user"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type newProjectRequest struct {
	Name          string `json:"name"`
	Git           string `json:"git"`
	DefaultBranch string `json:"defaultBranch"`
}

func New(c echo.Context) error {
	client := db.Connect()

	user, err := utils.AuthenticateUser(c)
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	pr := new(newProjectRequest)

	if err := c.Bind(pr); err != nil {
		return response.RespondError(c, *echo.ErrBadRequest)
	}

	project, err := client.Project.Create().
		AddUserIDs(user.Id).
		SetName(pr.Name).
		SetGit(pr.Git).
		SetDefaultBranch("main").
		Save(context.Background())
	if err != nil {
		println(err.Error())
		return response.RespondError(c, *echo.ErrBadRequest, err.Error())
	}
	return response.Respond(c, 200, *project)

}

func GetAll(c echo.Context) error {
	client := db.Connect()

	sessionUser, err := utils.AuthenticateUser(c)
	if err != nil {
		response.RespondError(c, *echo.ErrUnauthorized)
	}

	var projects []*ent.Project

	includeUser := utils.ParseIncludeQuery(c, "user")
	if includeUser {
		projects, err = client.Project.Query().
			Where(project.HasUsersWith(user.IDEQ(sessionUser.Id))).
			WithUsers().
			All(context.Background())

	} else {
		projects, err = client.Project.Query().
			Where(project.HasUsersWith(user.IDEQ(sessionUser.Id))).
			All(context.Background())

	}
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

func GetProject(c echo.Context) error {
	client := db.Connect()
	sessionUser, err := utils.AuthenticateUser(c)
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	id, _ := uuid.Parse(c.Param("id"))

	user, err := client.User.Get(context.Background(), sessionUser.Id)
	if err != nil {
		return response.RespondError(c, *echo.ErrUnauthorized)
	}

	proj, err := user.QueryProjects().Where(project.IDEQ(id)).First(context.Background())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return response.RespondError(c, *echo.ErrNotFound)
		default:
			return response.RespondError(c, *echo.ErrInternalServerError)
		}
	}

	return response.Respond(c, 200, *proj)
}
