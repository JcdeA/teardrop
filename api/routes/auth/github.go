package auth

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/palantir/go-githubapp/oauth2"
)

var (
	DefaultSessionKey = "oauth2.state"
	scopes            = []string{"user:email"}
	SessionStore      = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)

type SessionStateStore struct {
	Sessions *sessions.CookieStore
}

// https://stackoverflow.com/questions/44817570/set-array-struct-to-session-in-golang
func init() {
	gob.Register(&github.User{})
	gob.Register(uuid.UUID{})
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

func GithubOAuthHandler(config githubapp.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		ghHandler := oauth2.NewHandler(oauth2.GetConfig(config,
			scopes),

			oauth2.WithStore(&SessionStateStore{Sessions: SessionStore}),
			oauth2.OnLogin(func(w http.ResponseWriter, r *http.Request, login *oauth2.Login) {
				client := github.NewClient(login.Client)
				user, userResp, err := client.Users.Get(r.Context(), "")
				if err != nil {
					w.WriteHeader(401)

					resp := models.ErrorResponse{Status: 401, Message: "unauthorized", DocumentationUrl: "https://example.com"}

					respRaw, _ := resp.ToByteArray()
					w.Write(respRaw)
				}

				DBclient := db.Connect()

				acc, err := DBclient.Account.Create().
					SetProvider("github").
					SetAccessToken(login.Token.AccessToken).
					SetRefreshToken(login.Token.RefreshToken).
					SetTokenType(login.Token.TokenType).
					SetExpiresAt(login.Token.Expiry).
					SetProviderAccountId(fmt.Sprintf("%v", *user.ID)).
					SetScope(userResp.Header.Get("X-OAuth-Scopes")).
					Save(context.Background())

				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}

				dbUser, err := DBclient.User.Create().
					AddAccounts(acc).
					SetName(*user.Name).
					SetEmail(*user.Email).
					SetImage(*user.AvatarURL).
					Save(context.Background())

				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
					return
				}

				sess, _ := session.Get("session", c)
				sess.Values["userId"] = dbUser.ID
				sess.Values["user"] = *user
				sess.Save(c.Request(), c.Response())

				// redirect the user back to another page
				http.Redirect(w, r, "/api/deployments", http.StatusFound)
			}))
		ghHandler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
