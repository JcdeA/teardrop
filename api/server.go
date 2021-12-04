package api

import (
	"log"
	"time"

	"github.com/fosshostorg/teardrop/api/deployments"

	"github.com/fosshostorg/teardrop/api/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
)

func StartAPI() {
	godotenv.Load(".env.testing")

	ghconfig, err := ReadConfig(".config.yml")
	if err != nil {
		panic(err)
	}

	_, err = githubapp.NewDefaultCachingClientCreator(
		ghconfig.Github,
		githubapp.WithClientUserAgent("example-app/1.0.0"),
		githubapp.WithClientTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}

	ghOAuthHandler := auth.GithubOAuthHandler(*ghconfig)

	e := echo.New()
	APIGroup := e.Group("/api")
	APIGroup.GET("/deployments/get", deployments.Get)

	APIGroup.Any("/auth/github", ghOAuthHandler)
	APIGroup.GET("/ping", ping)

	echo.NotFoundHandler = notfound

	//client, _ := deploy.NewNomadClient("http://127.0.0.1:4646")
	//client.NewDeployment(types.Run{Name: "hello-world", Started: time.Now(), Environment: types.Production, Project: types.Project{ContainerImage: "registry.hub.docker.com/library/redis"}}, 1)

	log.Fatal(e.Start("0.0.0.0:3000"))

}
