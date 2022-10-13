<img align="right" width="150px" src="https://raw.githubusercontent.com/zeromicro/zero-doc/main/doc/images/go-zero.png">

# logx examples

English | [简体中文](readme-cn.md)

## Integrating with third-party logging libs

- zap
  - implementation: [https://github.com/zeromicro/zero-contrib/blob/main/logx/zapx/zap.go](https://github.com/zeromicro/zero-contrib/blob/main/logx/zapx/zap.go)
  - usage example: [https://github.com/zeromicro/zero-examples/blob/main/logx/zaplog/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/zaplog/main.go)
- logrus
  - implementation: [https://github.com/zeromicro/zero-contrib/blob/main/logx/logrusx/logrus.go](https://github.com/zeromicro/zero-contrib/blob/main/logx/logrusx/logrus.go)
  - usage example: [https://github.com/zeromicro/zero-examples/blob/main/logx/logrus/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/logrus/main.go)
- zerolog
  - implementation：[https://github.com/zeromicro/zero-contrib/blob/main/logx/zerologx/zerolog.go](https://github.com/zeromicro/zero-contrib/blob/main/logx/zerologx/zerolog.go)
  - usage example: [https://github.com/zeromicro/zero-examples/blob/main/logx/zerolog/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/zerolog/main.go)

For more libs, please implement and PR to [https://github.com/zeromicro/zero-contrib](https://github.com/zeromicro/zero-contrib)

## Write the logs to specific stores

`logx` defined two interfaces to let you customize `logx` to write logs into any stores.

- `logx.NewWriter(w io.Writer)`
- `logx.SetWriter(writer logx.Writer)`

For example, if we want to write the logs into kafka instead of console or files, we can do it like below:

```go
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
		return 0, err
	}

	return len(p), nil
}

func main() {
	pusher := kq.NewPusher([]string{"localhost:9092"}, "go-zero")
	defer pusher.Close()

	writer := logx.NewWriter(NewKafkaWriter(pusher))
	logx.SetWriter(writer)
  
	// more code
}
```

Complete code: [https://github.com/zeromicro/zero-examples/blob/main/logx/tokafka/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/tokafka/main.go)

## Filtering sensitive fields

If we need to prevent the `password` fields from logging, we can do it like below:

```go
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
	// setup logx to make sure originalWriter not nil,
	// the injected writer is only for filtering, like a middleware.

	originalWriter := logx.Reset()
	writer := NewSensitiveLogger(originalWriter)
	logx.SetWriter(writer)

	logx.Infov(Message{
		Name:     "foo",
		Password: "shouldNotAppear",
		Message:  "bar",
	})
  
	// more code
}
```

Complete code: [https://github.com/zeromicro/zero-examples/blob/main/logx/filterfields/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/filterfields/main.go)

## More examples

[https://github.com/zeromicro/zero-examples/tree/main/logx](https://github.com/zeromicro/zero-examples/tree/main/logx)

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!