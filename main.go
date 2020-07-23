package main

import (
	"github.com/i-coder-robot/go-micro-action-user-api/config"
	"github.com/i-coder-robot/go-micro-action-user-api/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	var Conf = config.Conf
	service := micro.NewService(
		micro.Name(Conf.Name),
		micro.Version(Conf.Version),
		micro.WrapHandler(Conf.Middleware().Wrapper), //验证权限
	)
	service.Init()

	// 注册服务
	h := handler.Handler{
		Server: service.Server(),
	}
	h.Register()
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
