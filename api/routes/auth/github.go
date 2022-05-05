package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/fosshostorg/teardrop/api"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/palantir/go-githubapp/oauth2"
)

var (
	DefaultSessionKey = "oauth2.state"
	scopes            = []string{"user:email"}
)

type SessionStateStore struct {
	Sessions *sessions.CookieStore
}

func (s *SessionStateStore) GenerateState(w http.ResponseWriter, r *http.Request) (string, error) {
	sess, _ := s.Sessions.Get(r, "session")

	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return "", errors.Wrap(err, "failed to generate state value")
	}

	state := hex.EncodeToString(b)
	sess.Values[DefaultSessionKey] = state

	return state, sess.Save(r, w)
}

func (s *SessionStateStore) VerifyState(r *http.Request, expected string) (bool, error) {
	sess, _ := s.Sessions.Get(r, "session")

	state := fmt.Sprint(sess.Values[DefaultSessionKey])

	return subtle.ConstantTimeCompare([]byte(expected), []byte(state)) == 1, nil
}

func GithubOAuthHandler(c githubapp.Config) echo.HandlerFunc {
	ghHandler := oauth2.NewHandler(oauth2.GetConfig(c,
		scopes),
		oauth2.WithStore(&SessionStateStore{Sessions: api.SessionStore}),

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
