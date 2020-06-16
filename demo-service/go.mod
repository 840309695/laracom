module laracom/demo-service

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/nonfu/laracom/demo-service v0.0.0-20200527135610-9e24a016004e
	github.com/nonfu/laracom/user-service v0.0.0-20200410143932-a6f4d5f4b264
	google.golang.org/protobuf v1.24.0 // indirect
)
