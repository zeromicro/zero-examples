package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/tracing/remote/user"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.json", "the config file")

type UserServer struct {
	lock     sync.Mutex
	alive    bool
	downTime time.Time
}

func NewUserServer() *UserServer {
	return &UserServer{
		alive: true,
	}
}

func (gs *UserServer) GetGrade(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	fmt.Println("=>", req)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Response: "hello from " + hostname,
	}, nil
}

func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, NewUserServer())
	})
	server.Start()
}
