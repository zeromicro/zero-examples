package main

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	Message struct {
		Name     string
		Password string
		Message  string
	}

	SensitiveLogger struct {
		logx.Writer
	}
)

func NewSensitiveLogger(writer logx.Writer) *SensitiveLogger {
	return &SensitiveLogger{
		Writer: writer,
	}
}

func (l *SensitiveLogger) Info(msg interface{}, fields ...logx.LogField) {
	if m, ok := msg.(Message); ok {
		l.Writer.Info(Message{
			Name:     m.Name,
			Password: "******",
			Message:  m.Message,
		}, fields...)
	} else {
		l.Writer.Info(msg, fields...)
	}
}

func main() {
	logx.SetUp(logx.LogConf{
		Mode: "console",
	})
	originalWriter := logx.Reset()
	writer := NewSensitiveLogger(originalWriter)
	logx.SetWriter(writer)

	logx.Infow("infow foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Infov(Message{
		Name:     "foo",
		Password: "shouldNotAppear",
		Message:  "bar",
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
