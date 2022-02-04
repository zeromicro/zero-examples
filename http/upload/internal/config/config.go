package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Path string `json:",default=."`
}
