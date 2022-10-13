package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Host         string `json:",default=0.0.0.0"`
	Port         int
	LogMode      string        `json:",options=[file,console]"`
	Verbose      bool          `json:",optional"`
	MaxConns     int           `json:",default=10000"`
	MaxBytes     int64         `json:",default=1048576"`
	Timeout      time.Duration `json:",default=3s"`
	CpuThreshold int64         `json:",default=900,range=[0:1000]"`
}

func main() {
	var c Config
	conf.MustLoad("config.toml", &c)
	spew.Dump(c)
}
