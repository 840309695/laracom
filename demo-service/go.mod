module github.com/840309695/laracom/demo-service

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/micro/go-micro v1.18.0
	github.com/nonfu/laracom/user-service v0.0.0-20200527135610-9e24a016004e
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933

)
