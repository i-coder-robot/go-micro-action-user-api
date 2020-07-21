package handler

import (
	"context"
	client "github.com/i-coder-robot/go-micro-action-core/client"
	pb "github.com/i-coder-robot/go-micro-action-user-api/proto/auth"
)

// Permission 权限结构
type Permission struct {
	ServiceName string
}

// All 权限列表
func (srv *Permission) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.All", req, res)
}

// List 权限列表
func (srv *Permission) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.List", req, res)
}

// Get 获取权限
func (srv *Permission) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.Get", req, res)
}

// Create 创建权限
func (srv *Permission) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.Create", req, res)
}

// Update 更新权限
func (srv *Permission) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.Update", req, res)
}

// Delete 删除权限
func (srv *Permission) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Permissions.Delete", req, res)
}
