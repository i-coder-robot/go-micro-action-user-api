package handler

import (
	"context"
	"fmt"
	client "github.com/i-coder-robot/go-micro-action-core/client"

	pb "github.com/i-coder-robot/go-micro-action-user-api/proto/auth"
)

// Auth 授权服务处理
type Auth struct {
	ServiceName string
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	fmt.Println("user-api,auth.auth")
	fmt.Println("ServiceName:"+srv.ServiceName)
	fmt.Println(req)
	err = client.Call(ctx, srv.ServiceName, "Auth.Auth", req, res)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Logout 登录退出
// 后期开发
// token 增加前期认证根据根据存储的 token 进行认证 和过期查询
// 通过删除 存储 token 来登出
// 通过存储多种类型的 token 来实现多端登录
// token id type type=登录类型
// 过期时间默认还是在 jwt token 中存储
func (srv *Auth) Logout(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Auth.ValidateToken", req, res)
}
