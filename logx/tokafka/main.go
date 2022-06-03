package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaWriter struct {
	Pusher *kq.Pusher
}

func NewKafkaWriter(pusher *kq.Pusher) *KafkaWriter {
	return &KafkaWriter{
		Pusher: pusher,
	}
}

func (w *KafkaWriter) Write(p []byte) (n int, err error) {
	// writing log with newlines, trim them.
	if err := w.Pusher.Push(strings.TrimSpace(string(p))); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return len(p), nil
}

func main() {
	pusher := kq.NewPusher([]string{"localhost:9092"}, "go-zero")
	defer pusher.Close()

	writer := logx.NewWriter(NewKafkaWriter(pusher))
	logx.SetWriter(writer)

	for {
		time.Sleep(time.Second * 5)
		logx.Infow("infow foo",
			logx.Field("url", "http://localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
		logx.Errorw("errorw foo",
			logx.Field("url", "http://localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
		logx.Sloww("sloww foo",
			logx.Field("url", "http://localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
		logx.Error("error")
		logx.Infov(map[string]interface{}{
			"url":     "localhost:8080/hello",
			"attempt": 3,
			"backoff": time.Second,
			"value":   "foo",
		})
		logx.WithDuration(1100*time.Microsecond).Infow("infow withduration",
			logx.Field("url", "localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
		logx.WithContext(context.Background()).WithDuration(1100*time.Microsecond).Errorw(
			"errorw withcontext withduration",
			logx.Field("url", "localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
		logx.WithDuration(1100*time.Microsecond).WithContext(context.Background()).Errorw(
			"errorw withduration withcontext",
			logx.Field("url", "localhost:8080/hello"),
			logx.Field("attempt", 3),
			logx.Field("backoff", time.Second),
		)
	}
}
