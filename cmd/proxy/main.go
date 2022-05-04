package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/domain"
	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var proxyMap map[string]cachedProxy = map[string]cachedProxy{}

type cachedProxy struct {
	proxy     *httputil.ReverseProxy
	TTL       time.Duration
	createdAt time.Time
}

type baseHandle struct {
	DBClient *ent.Client
}

var ctx = context.Background()

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("x-forwarded-host")
	//host := r.Host

	opt, _ := redis.ParseURL("redis://localhost:6379")
	opt.DB = models.ProxiesDB
	redisdb := redis.NewClient(opt)

	var remoteUrl *url.URL

	if target, err := redisdb.Get(ctx, host).Result(); err == nil {
		remoteUrl, err = url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

	} else {
		proxies, err := h.DBClient.Domain.Query().Where(domain.DomainEQ(host)).All(ctx)
		if err != nil {
			log.Println(err)
		}
		var proxyValues []string
		for _, p := range proxies {
			proxyValues = append(proxyValues, *&p.Edges.Deployment.Address)
		}

		_, err = redisdb.Set(ctx, host, proxyValues[0], time.Minute*10).Result()

		if err != nil {

			log.Println(err)
			w.WriteHeader(403)
			w.Write([]byte("403: Host forbidden " + host))
		}

		a := *proxies[0]

		remoteUrl, err = url.Parse(a.Edges.Deployment.Address)

		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

	}

	var proxy *httputil.ReverseProxy

	if cached, ok := proxyMap[remoteUrl.Host]; ok {
		cached.proxy.ServeHTTP(w, r)
		return
	}

	proxy = newProxy(remoteUrl)
	proxyMap[remoteUrl.Host] = cachedProxy{
		proxy:     proxy,
		TTL:       time.Minute * 10,
		createdAt: time.Now(),
	}

	proxy.ServeHTTP(w, r)
	return

}

func main() {
	client, err := ent.Open("sqlite3", "file:lmao?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// if err := client.Schema.Create(
	// 	context.TODO(),
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithDropColumn(true),
	// ); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	// testing stuffs
	// proj := client.Project.Create().SetName("teardrop-test").SetGit("https://github.com/JcdeA/website.git").SetDefaultBranch("main").SaveX(context.TODO())
	// client.User.Create().SetName("babo").SetEmail("io@fosshost.org").AddProjects(proj).SaveX(context.TODO())
	// client.Proxy.Create().SetDomain("localhost:1323").SetOrigin("http://localhost:3000").SetProjects(proj).SaveX(context.TODO())

	e := echo.New()
	//e.Use(middleware.Recover())
	//e.Use(middleware.Logger())

	e.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler { return &baseHandle{DBClient: client} }))

	e.Logger.Fatal(e.Start(":1323"))
}
