package handler

import (
	"context"
	"github.com/go-log/log"
	"github.com/i-coder-robot/go-micro-action-core/client"
	"github.com/i-coder-robot/go-micro-action-user-api/config"
	authPB "github.com/i-coder-robot/go-micro-action-user-api/proto/auth"
	casbinPB "github.com/i-coder-robot/go-micro-action-user-api/proto/casbin"
	frontPermitPB "github.com/i-coder-robot/go-micro-action-user-api/proto/frontPermit"
	healthPB "github.com/i-coder-robot/go-micro-action-user-api/proto/health"
	permissionPB "github.com/i-coder-robot/go-micro-action-user-api/proto/permission"
	rolePB "github.com/i-coder-robot/go-micro-action-user-api/proto/role"
	userPB "github.com/i-coder-robot/go-micro-action-user-api/proto/user"
	PB "github.com/i-coder-robot/go-micro-action-user/user/proto/permission"
	"github.com/micro/go-micro/v2/server"
	"time"
)

// Handler 注册方法
type Handler struct {
	Server server.Server
}

var Conf = config.Conf

// Register 注册
func (srv *Handler) Register() {
	userPB.RegisterUsersHandler(srv.Server, &User{Conf.Service["user"]})
	authPB.RegisterAuthHandler(srv.Server, &Auth{Conf.Service["user"]})
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server, &FrontPermit{Conf.Service["user"]})
	permissionPB.RegisterPermissionsHandler(srv.Server, &Permission{Conf.Service["user"]})
	rolePB.RegisterRolesHandler(srv.Server, &Role{Conf.Service["user"]})
	casbinPB.RegisterCasbinHandler(srv.Server, &Casbin{Conf.Service["user"]}) // 权限管理服务实现
	healthPB.RegisterHealthHandler(srv.Server, &Health{})

	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.Service["user"], "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
