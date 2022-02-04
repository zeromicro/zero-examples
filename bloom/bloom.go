package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	store := redis.NewRedis("localhost:6379", "node")
	filter := bloom.New(store, "testbloom", 64)
	filter.Add([]byte("kevin"))
	filter.Add([]byte("wan"))
	fmt.Println(filter.Exists([]byte("kevin")))
	fmt.Println(filter.Exists([]byte("wan")))
	fmt.Println(filter.Exists([]byte("nothing")))
}
