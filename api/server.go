package api

import (
	"log"
	"time"

	"github.com/fosshostorg/teardrop/api/routes/deployments"
	"github.com/fosshostorg/teardrop/api/routes/webhook"

	"github.com/fosshostorg/teardrop/api/routes/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
)

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

	APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	APIGroup.GET("/webhook/github", echo.WrapHandler(
		githubapp.NewDefaultEventDispatcher(c,
			&webhook.PushHandler{Client: cc},
		)))
	APIGroup.GET("/ping", ping)

}
