// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Users service

func NewUsersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Users service

type UsersService interface {
	// 用过 用户名、邮箱、手机 查询用户是否存在
	Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 绑定手机 更新同时更新其他用户信息
	MobileBuild(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 用户通过 token 自己更新数据 只可以更改 用户名、昵称、头像
	SelfUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取用户
	Info(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取用户列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取用户
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建用户
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新用户
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除用户
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Exist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) MobileBuild(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.MobileBuild", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) SelfUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.SelfUpdate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Info(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Info", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	// 用过 用户名、邮箱、手机 查询用户是否存在
	Exist(context.Context, *Request, *Response) error
	// 绑定手机 更新同时更新其他用户信息
	MobileBuild(context.Context, *Request, *Response) error
	// 用户通过 token 自己更新数据 只可以更改 用户名、昵称、头像
	SelfUpdate(context.Context, *Request, *Response) error
	// 根据 唯一 获取用户
	Info(context.Context, *Request, *Response) error
	// 获取用户列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取用户
	Get(context.Context, *Request, *Response) error
	// 创建用户
	Create(context.Context, *Request, *Response) error
	// 更新用户
	Update(context.Context, *Request, *Response) error
	// 删除用户
	Delete(context.Context, *Request, *Response) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		Exist(ctx context.Context, in *Request, out *Response) error
		MobileBuild(ctx context.Context, in *Request, out *Response) error
		SelfUpdate(ctx context.Context, in *Request, out *Response) error
		Info(ctx context.Context, in *Request, out *Response) error
		List(ctx context.Context, in *Request, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
		Create(ctx context.Context, in *Request, out *Response) error
		Update(ctx context.Context, in *Request, out *Response) error
		Delete(ctx context.Context, in *Request, out *Response) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) Exist(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Exist(ctx, in, out)
}

func (h *usersHandler) MobileBuild(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.MobileBuild(ctx, in, out)
}

func (h *usersHandler) SelfUpdate(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.SelfUpdate(ctx, in, out)
}

func (h *usersHandler) Info(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Info(ctx, in, out)
}

func (h *usersHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.List(ctx, in, out)
}

func (h *usersHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Get(ctx, in, out)
}

func (h *usersHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Create(ctx, in, out)
}

func (h *usersHandler) Update(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Update(ctx, in, out)
}

func (h *usersHandler) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Delete(ctx, in, out)
}
