package main

import (
	"flag"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var port = flag.Int("port", 3333, "the port to listen")

type (
	AnotherService struct{}

	Request struct {
		User string `form:"user"`
	}
)

func (s *AnotherService) GetToken() string {
	return stringx.Rand()
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "static-middleware")
		next(w, r)
	}
}

func middlewareWithAnotherService(s *AnotherService) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Middleware", s.GetToken())
			next(w, r)
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "hello, "+req.User)
}

func main() {
	flag.Parse()

	server := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: *port,
	})
	defer server.Stop()

	server.Use(middleware)
	server.Use(middlewareWithAnotherService(new(AnotherService)))
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/greet",
		Handler: handle,
	})
	server.Start()
}
