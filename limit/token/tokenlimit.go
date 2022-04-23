package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	burst   = 100
	rate    = 100
	seconds = 5
)

var (
	rdx     = flag.String("redis", "localhost:6379", "the redis, default localhost:6379")
	rdxKey  = flag.String("redisKey", "rate", "the redis key, default rate")
	rdxPass = flag.String("redisPass", "", "the redis password")
	threads = flag.Int("threads", runtime.NumCPU(), "the concurrent threads, default to cores")
)

func main() {
	flag.Parse()

	store := redis.New(*rdx, redis.WithPass(*rdxPass))
	fmt.Println(store.Ping())
	limit := limit.NewTokenLimiter(rate, burst, store, *rdxKey)
	timer := time.NewTimer(time.Second * seconds)
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	var allowed, denied int32
	var wait sync.WaitGroup
	for i := 0; i < *threads; i++ {
		wait.Add(1)
		go func() {
			for {
				select {
				case <-quit:
					wait.Done()
					return
				default:
					if limit.Allow() {
						atomic.AddInt32(&allowed, 1)
					} else {
						atomic.AddInt32(&denied, 1)
					}
				}
			}
		}()
	}

	wait.Wait()
	fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)
}
