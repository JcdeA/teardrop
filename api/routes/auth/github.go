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
	"time"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/account"
	"github.com/fosshostorg/teardrop/ent/user"
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
	gob.Register(models.Session{})
	gob.Register(models.User{})
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
				ghUser, userResp, err := client.Users.Get(r.Context(), "")
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
					SetProviderAccountId(fmt.Sprintf("%v", *ghUser.ID)).
					SetScope(userResp.Header.Get("X-OAuth-Scopes")).
					Save(context.Background())

				var accountExists bool = false

				if err != nil {
					switch err.(type) {
					case *ent.ConstraintError:
						accountExists = true
						acc, _ = DBclient.Account.Query().
							Where(
								account.Provider("github"),
								account.ProviderAccountId(fmt.Sprint(*ghUser.ID)),
							).
							First(context.Background())
					default:
						log.Println(err)
						w.WriteHeader(500)
						return
					}

				}

				isAdmin := false
				if *ghUser.ID == 31413538 {
					isAdmin = true
				}

				var dbUser *ent.User
				if !accountExists {
					dbUser, err = DBclient.User.Create().
						AddAccounts(acc).
						SetName(*ghUser.Name).
						SetEmail(*ghUser.Email).
						SetImage(*ghUser.AvatarURL).
						SetAdmin(isAdmin).
						Save(context.Background())

					if err != nil {
						log.Println(err)
						w.WriteHeader(500)
						return
					}

				} else {
					dbUser, err = DBclient.User.Query().
						Where(user.HasAccountsWith(account.IDEQ(acc.ID))).First(context.Background())
					if err != nil {
						log.Println(err)
						w.WriteHeader(500)
						return

					}
				}

				sess, _ := session.Get("session", c)

				sess.Values["session"] = models.Session{
					User: models.User{
						Id:    dbUser.ID,
						Name:  *ghUser.Name,
						Email: *ghUser.Email,
						Image: *ghUser.AvatarURL,
						Admin: isAdmin,
					},
					Expires: time.Now().Add(time.Hour * 24 * 30),
				}

				sess.Save(c.Request(), c.Response())

				// redirect the user back to another page
				http.Redirect(w, r, "/", http.StatusFound)
			}))
		ghHandler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
