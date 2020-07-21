package handler

import (
	"context"
	client "github.com/i-coder-robot/go-micro-action-core/client"
	pb "github.com/i-coder-robot/go-micro-action-user-api/proto/auth"
)

// Role 角色结构
type Role struct {
	ServiceName string
}

// All 权限列表
func (srv *Role) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.All", req, res)
}

// List 角色列表
func (srv *Role) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.List", req, res)
}

// Get 获取角色
func (srv *Role) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.Get", req, res)
}

// Create 创建角色
func (srv *Role) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.Create", req, res)
}

// Update 更新角色
func (srv *Role) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.Update", req, res)
}

// Delete 删除角色
func (srv *Role) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Roles.Delete", req, res)
}

