package main

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type baseHandle struct {
	DBClient *ent.Client
}

var ctx = context.Background()

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("x-forwarded-host")

	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       models.HostsDB,
	})

	var remoteUrl *url.URL

	if target, err := redisdb.Get(ctx, host).Result(); err != nil {
		remoteUrl, err = url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}
	} else {

	}

	proxy := newProxy(remoteUrl)
	proxy.ServeHTTP(w, r)
	return

	w.Write([]byte("403: Host forbidden " + host))
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler { return &baseHandle{DBClient: client} }))

	e.Logger.Fatal(e.Start(":1323"))
}
