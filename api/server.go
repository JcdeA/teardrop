package api

import (
	"log"
	"os"

	"github.com/fosshostorg/teardrop/api/routes/deployments"
	"github.com/fosshostorg/teardrop/api/routes/domains"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"

	_ "github.com/mattn/go-sqlite3"
)

var (
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)


func StartAPI() {
	// Initialize DB Client; DB connection is a global in the internal/pkg/db package
	db.Connect()

	godotenv.Load()

	// ghconfig, err := ReadConfig(".config.yml")
	// if err != nil {
	// 	panic(err)
	// }
	e := echo.New()

	registerRoutes(e, githubapp.Config{})

	e.Use(session.Middleware(SessionStore))

	e.Use(echo.WrapMiddleware(
		csrf.Protect([]byte("8hmirhx6f9uuaqb5ym6tf9283ea2ibu6")),
	))

	log.Fatal(e.Start("0.0.0.0:3000"))
}

func registerRoutes(e *echo.Echo, c githubapp.Config) {
	// cc, err := githubapp.NewDefaultCachingClientCreator(
	// 	c,
	// 	githubapp.WithClientUserAgent("example-app/1.0.0"),
	// 	githubapp.WithClientTimeout(3*time.Second),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	APIGroup := e.Group("/api")
	APIGroup.GET("/deployments", deployments.Get)

	APIGroup.POST("/domains", domains.Add)

	// APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	// APIGroup.GET("/webhook/github", echo.WrapHandler(
	// 	githubapp.NewDefaultEventDispatcher(c,
	// 		&webhook.PushHandler{Client: cc},
	// 	)))
	APIGroup.GET("/ping", ping)

}
