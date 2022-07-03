package main

import (
	"flag"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var config = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c rest.RestConf
	conf.MustLoad(*config, &c)

	svr := rest.MustNewServer(c)
	defer svr.Stop()

	svr.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/:file",
		Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))).ServeHTTP,
	})
	svr.Start()
}
