package auth

import (
	"net/http"

	"github.com/fosshostorg/teardrop/models"
	"github.com/google/go-github/github"

	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/palantir/go-githubapp/oauth2"
)

func GithubOAuthHandler(c githubapp.Config) echo.HandlerFunc {
	ghHandler := oauth2.NewHandler(oauth2.GetConfig(c,
		[]string{"user:email"}),
		oauth2.OnLogin(func(w http.ResponseWriter, r *http.Request, login *oauth2.Login) {

			client := github.NewClient(login.Client)
			user, _, err := client.Users.Get(r.Context(), "")
			if err != nil {
				w.WriteHeader(401)

				resp := models.Response{Status: 401, Message: "unauthorized", DocumentationUrl: "https://example.com"}

				respRaw, _ := resp.ToByteArray()
				w.Write(respRaw)
			}
			println(user.Email)

			// redirect the user back to another page
			http.Redirect(w, r, "/deployments", http.StatusFound)
		}))
	return func(c echo.Context) error {
		ghHandler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
