package handler

import (
	"context"
	pb "github.com/i-coder-robot/go-micro-action-user-api/proto/health"
)

// Health 用户结构
type Health struct {
}

// Health 用户是否存在
func (srv *Health) Health(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = true
	return nil
}

