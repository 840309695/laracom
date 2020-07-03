package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "github.com/840309695/laracom/demo-service/proto/demo"
	"github.com/micro/go-micro/metadata"
	userpb "github.com/nonfu/laracom/user-service/proto/user"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/840309695/laracom/demo-service/trace"
	"github.com/opentracing/opentracing-go"
	"log"
)

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	// 从微服务上下文中获取追踪信息
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// 创建新的 Span 并将其绑定到微服务上下文
	sp = opentracing.StartSpan("SayHello", opentracing.ChildOf(wireContext))
	// 记录请求
	sp.SetTag("req", req)
	defer func() {
		// 记录响应
		sp.SetTag("res", rsp)
		// 在函数返回 stop span 之前，统计函数执行时间
		sp.Finish()
	}()

	rsp.Text = "你好, " + req.Name
	return nil
}
func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {
	service := micro.NewService()
	client := userpb.NewUserServiceClient("laracom.service.user", service.Client())
	resp, err := client.GetById(context.TODO(), &userpb.User{Id: req.Id})
	if err != nil {
		return err
	}
	rsp.Text = "你好, " + resp.User.Name
	return nil
}

func main()  {
	// 初始化全局服务追踪
	t, io, err := trace.NewTracer("laracom.service.demo", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(micro.Name("laracom.demo.service"))
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}