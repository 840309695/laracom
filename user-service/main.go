package main

import (
	"fmt"
	"log"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-micro/broker/nats"
	database "github.com/840309695/laracom/user-service/db"
	pb "github.com/840309695/laracom/user-service/proto/user"
	repository "github.com/840309695/laracom/user-service/repo"
	"github.com/840309695/laracom/user-service/handler"
	"github.com/840309695/laracom/user-service/service"
	"github.com/840309695/laracom/user-service/model"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// 启动 HTTP 服务监听 Prometheus 客户端监控数据采集
func prometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}
func main() {

	// 创建数据库连接，程序退出时断开连接
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 和 Laravel 数据库迁移类似
	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.PasswordReset{})

	// 初始化 Repo 实例用于后续数据库操作
	repo := &repository.UserRepository{db}
	// 初始化 token service
	token := &service.TokenService{repo}
	resetRepo := &repository.PasswordResetRepository{db}

	// 以下是 Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.service.user"),
		micro.Version("latest"),  // 新增接口版本参数
	)
	srv.Init()
	// 获取 Broker 实例
	pubSub := srv.Server().Options().Broker

	// 注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo, resetRepo,token,pubSub})

	// 采集监控数据
	prometheusBoot()
	// 启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}