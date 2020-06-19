module github.com/840309695/laracom/demo-service

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/nonfu/laracom/user-service v0.0.0-20200527135610-9e24a016004e
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933

)
