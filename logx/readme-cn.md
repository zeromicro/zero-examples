<IMG align="right" width="150px" src="https://raw.githubusercontent.com/zeromicro/zero-doc/main/doc/images/go-zero.png">

# logx examples

[English](readme.md) | 简体中文

## 第三方日志库集成

- zap
  - 实现：[https://github.com/zeromicro/zero-contrib/blob/main/logx/zapx/zap.go](https://github.com/zeromicro/zero-contrib/blob/main/logx/zapx/zap.go)
  - 使用示例：[https://github.com/zeromicro/zero-examples/blob/main/logx/zaplog/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/zaplog/main.go)
- logrus
  - 实现：[https://github.com/zeromicro/zero-contrib/blob/main/logx/logrusx/logrus.go](https://github.com/zeromicro/zero-contrib/blob/main/logx/logrusx/logrus.go)
  - 使用示例：[https://github.com/zeromicro/zero-examples/blob/main/logx/logrus/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/logrus/main.go)

对于其它的日志库，请参考上面示例实现，并欢迎提交 `PR` 到 [https://github.com/zeromicro/zero-contrib](https://github.com/zeromicro/zero-contrib)

## 将日志写到指定的存储

`logx`定义了两个接口，方便自定义 `logx`，将日志写入任何存储。

- `logx.NewWriter(w io.Writer)`
- `logx.SetWriter(write logx.Writer)`

例如，如果我们想把日志写进kafka，而不是控制台或文件，我们可以像下面这样做。

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

完整代码：[https://github.com/zeromicro/zero-examples/blob/main/logx/tokafka/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/tokafka/main.go)

## 过滤敏感字段

如果我们需要防止  `password` 字段被记录下来，我们可以像下面这样实现。

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

完整代码：[https://github.com/zeromicro/zero-examples/blob/main/logx/filterfields/main.go](https://github.com/zeromicro/zero-examples/blob/main/logx/filterfields/main.go)

## 更多示例

[https://github.com/zeromicro/zero-examples/tree/main/logx](https://github.com/zeromicro/zero-examples/tree/main/logx)

## Give a Star! ⭐

如果你正在使用或者觉得这个项目对你有帮助，请 **star** 支持，感谢！
