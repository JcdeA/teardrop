package api

import (
	"log"
	"net/url"
	"time"

	"github.com/fosshostorg/teardrop/api/routes/auth"
	"github.com/fosshostorg/teardrop/api/routes/deployments"
	"github.com/fosshostorg/teardrop/api/routes/domains"
	"github.com/fosshostorg/teardrop/api/routes/projects"
	"github.com/fosshostorg/teardrop/api/routes/webhook"
	"github.com/fosshostorg/teardrop/api/utils"
	"github.com/fosshostorg/teardrop/internal/pkg/db"
	"github.com/fosshostorg/teardrop/internal/pkg/proxyutils"
	"github.com/gorilla/csrf"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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

	//e.Use(middleware.Logger())
	registerRoutes(e, ghconfig.Github)

	e.Use(session.Middleware(auth.SessionStore))

	e.Use(echo.WrapMiddleware(
		csrf.Protect([]byte("8hmirhx6f9uuaqb5ym6tf9283ea2ibu6"), csrf.Secure(false)),
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

	APIGroup.POST("/projects", projects.New)

	APIGroup.GET("/projects", projects.GetAll)

	APIGroup.GET("/projects/:id", projects.GetProject)

	APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	APIGroup.GET("/webhook/github",
		echo.WrapHandler(
			githubapp.NewDefaultEventDispatcher(
				c,
				&webhook.PushHandler{Client: cc},
			),
		),
	)

	APIGroup.GET("/ping", utils.Ping)

	APIGroup.GET("/csrf", func(c echo.Context) error {
		c.Response().Header().Set("x-csrf-token", csrf.Token(c.Request()))

		return c.JSON(200, echo.Map{"csrfToken": csrf.Token(c.Request())})
	})

	APIGroup.GET("/auth/session", auth.GetSession)

	frontend, _ := url.Parse("http://localhost:4000")

	e.Any("/*", echo.WrapHandler(proxyutils.NewProxy(frontend)))

}
