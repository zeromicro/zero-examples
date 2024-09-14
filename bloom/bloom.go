package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	//store := redis.New("192.168.8.6:6379") //过时的
	redconf := redis.RedisConf{Host: "192.168.8.6:6379",
	Pass: "redis_6X6RSk", Type: "node"}
	store := redis.MustNewRedis(redconf)
	filter := bloom.New(store, "testbloom", 64)
	filter.Add([]byte("kevin"))
	filter.Add([]byte("wan"))
	fmt.Println(filter.Exists([]byte("kevin")))
	fmt.Println(filter.Exists([]byte("wan")))
	fmt.Println(filter.Exists([]byte("nothing")))
}
