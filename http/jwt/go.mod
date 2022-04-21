module github.com/zeromicro/zero-examples/jwt

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/zeromicro/go-zero v1.3.2
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
