package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

type Request struct {
	User string `form:"user"`
}

func first(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "first")
		next(w, r)
	}
}

func second(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "second")
		next(w, r)
	}
}

func handleHomeHello(w http.ResponseWriter, r *http.Request) {
	ddl, ok := r.Context().Deadline()
	fmt.Println(ok)
	fmt.Printf("%#v\n", ddl)
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(time.Millisecond * 1500)
	httpx.OkJson(w, "welcome home, "+req.User)
}

func handleApiHello(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(time.Millisecond * 1500)
	httpx.OkJson(w, "hello, "+req.User)
}

func handleApiHi(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "hi, "+req.User)
}

func main() {
	flag.Parse()
	logx.DisableStat()

	var c rest.RestConf
	conf.MustLoad(*configFile, &c)
	srv := rest.MustNewServer(c)
	defer srv.Stop()

	srv.Use(first)
	srv.Use(second)

	srv.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: handleHomeHello,
	}, rest.WithTimeout(time.Second))
	srv.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/hello",
			Handler: handleApiHello,
		},
		{
			Method:  http.MethodGet,
			Path:    "/hi",
			Handler: handleApiHi,
		},
	}, rest.WithPrefix("/api"), rest.WithTimeout(time.Second*5))

	srv.Start()
}
