package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("x-forwarded-host")

	

	if target, ok := "https://jcde.xyz", true; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}
		proxy := newProxy(remoteUrl)
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler { return &baseHandle{} }))

	e.Logger.Fatal(e.Start(":1323"))
}
