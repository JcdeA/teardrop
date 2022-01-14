package api

import (
	"time"

	"github.com/fosshostorg/teardrop/api/deployments"
	"github.com/fosshostorg/teardrop/api/webhook"
	"github.com/fosshostorg/teardrop/deploy"
	"github.com/fosshostorg/teardrop/models"

	"github.com/fosshostorg/teardrop/api/auth"
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

	client, _ := deploy.NewNomadClient("http://10.8.0.101:4646")
	err = client.NewDeployment(models.Run{Name: "hellohttp-2", Started: time.Now(), Environment: models.Production, Project: models.Project{ContainerImage: "nginxdemos/hello"}}, 31)

	if err != nil {
		panic(err)
	}
	registerRoutes(e, ghconfig.Github)

	// log.Fatal(e.Start("0.0.0.0:3000"))

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
	APIGroup.GET("/deployments/get", deployments.Get)

	APIGroup.Any("/auth/github", auth.GithubOAuthHandler(c))

	APIGroup.GET("/webhook/github", echo.WrapHandler(
		githubapp.NewDefaultEventDispatcher(c,
			&webhook.PushHandler{Client: cc},
		)))
	APIGroup.GET("/ping", ping)

	echo.NotFoundHandler = notfound
}
