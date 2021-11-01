package main

import (
	"flag"

	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/config"
	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/handler"
	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/svc"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	server.Start()
}
