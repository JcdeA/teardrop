package api

import (
	"log"
	"time"

	"github.com/fosshostorg/teardrop/api/routes/deployments"
	"github.com/fosshostorg/teardrop/api/routes/domains"
	"github.com/fosshostorg/teardrop/api/routes/webhook"
	"github.com/fosshostorg/teardrop/ent"
	"github.com/gorilla/csrf"

	"github.com/fosshostorg/teardrop/api/routes/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
)

var DB *ent.Client

func StartAPI() {
	godotenv.Load()

	ghconfig, err := ReadConfig(".config.yml")
	if err != nil {
		panic(err)
	}
	e := echo.New()

	if err != nil {
		panic(err)
	}
	registerRoutes(e, ghconfig.Github)

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

	APIGroup.GET("/domains", domains.Get)

	APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	APIGroup.GET("/webhook/github", echo.WrapHandler(
		githubapp.NewDefaultEventDispatcher(c,
			&webhook.PushHandler{Client: cc},
		)))
	APIGroup.GET("/ping", ping)

}
