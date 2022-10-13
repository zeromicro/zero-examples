package main

import (
	"bufio"
	"context"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

type MultiWriter struct {
	writer        logx.Writer
	consoleWriter logx.Writer
}

func NewMultiWriter(writer logx.Writer) (logx.Writer, error) {
	return &MultiWriter{
		writer:        writer,
		consoleWriter: logx.NewWriter(bufio.NewWriter(os.Stdout)),
	}, nil
}

func (w *MultiWriter) Alert(v interface{}) {
	w.consoleWriter.Alert(v)
	w.writer.Alert(v)
}

func (w *MultiWriter) Close() error {
	w.consoleWriter.Close()
	return w.writer.Close()
}

func (w *MultiWriter) Error(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Error(v, fields...)
	w.writer.Error(v, fields...)
}

func (w *MultiWriter) Info(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Info(v, fields...)
	w.writer.Info(v, fields...)
}

func (w *MultiWriter) Severe(v interface{}) {
	w.consoleWriter.Severe(v)
	w.writer.Severe(v)
}

func (w *MultiWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Slow(v, fields...)
	w.writer.Slow(v, fields...)
}

func (w *MultiWriter) Stack(v interface{}) {
	w.consoleWriter.Stack(v)
	w.writer.Stack(v)
}

func (w *MultiWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Stat(v, fields...)
	w.writer.Stat(v, fields...)
}

func main() {
	var c logx.LogConf
	conf.MustLoad("config.toml", &c)
	logx.MustSetup(c)

	fileWriter := logx.Reset()
	writer, err := NewMultiWriter(fileWriter)
	logx.Must(err)
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
