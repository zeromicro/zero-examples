package main

import (
	"flag"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/zero-examples/http/httpc/types"
)

var port = flag.Int("port", 3333, "the port to listen")

func handle(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, req)
}

func main() {
	flag.Parse()

	svr := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: *port,
	})
	defer svr.Stop()

	svr.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/nodes/:node",
		Handler: handle,
	})
	svr.Start()
}
