package api

import (
	"log"
	"time"

	"github.com/fosshostorg/teardrop/api/routes/auth"
	"github.com/fosshostorg/teardrop/api/routes/deployments"
	"github.com/fosshostorg/teardrop/api/routes/domains"
	"github.com/fosshostorg/teardrop/api/routes/projects"
	"github.com/fosshostorg/teardrop/api/routes/webhook"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/gorilla/csrf"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/palantir/go-githubapp/githubapp"

	_ "github.com/mattn/go-sqlite3"
)

func StartAPI() {
	// Initialize DB Client; DB connection is a global in the internal/pkg/db package
	db.Connect()

	godotenv.Load()

	ghconfig, err := ReadEnv()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	registerRoutes(e, ghconfig.Github)

	e.Use(session.Middleware(auth.SessionStore))

	e.Use(echo.WrapMiddleware(
		csrf.Protect([]byte("8hmirhx6f9uuaqb5ym6tf9283ea2ibu6")),
	))

	log.Fatal(e.Start("0.0.0.0:3000"))
}

func registerRoutes(e *echo.Echo, c githubapp.Config) {
	cc, err := githubapp.NewDefaultCachingClientCreator(
		c,
		githubapp.WithClientUserAgent("example-app/1.0.0"),
		githubapp.WithClientTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}

	APIGroup := e.Group("/api")
	APIGroup.GET("/deployments", deployments.Get)

	APIGroup.POST("/domains", domains.Add)

	APIGroup.GET("/projects/new", projects.New)

	APIGroup.GET("/projects", projects.Get)

	APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	APIGroup.GET("/webhook/github",
		echo.WrapHandler(
			githubapp.NewDefaultEventDispatcher(
				c,
				&webhook.PushHandler{Client: cc},
			),
		),
	)

	APIGroup.GET("/ping", ping)

}
