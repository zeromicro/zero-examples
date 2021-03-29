module github.com/zeromicro/zero-examples

go 1.15

require (
	github.com/dchest/siphash v1.2.2
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.4.2
	github.com/google/gops v0.3.14
	github.com/gorilla/websocket v1.4.2
	github.com/stretchr/testify v1.7.0
	github.com/tal-tech/go-zero v1.1.6
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/grpc v1.31.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
