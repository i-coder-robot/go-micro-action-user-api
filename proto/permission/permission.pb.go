// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/permission/permission.proto

package permission

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Permission struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Service     string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method      string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Auth        bool   `protobuf:"varint,6,opt,name=auth,proto3" json:"auth,omitempty"`
	Policy      bool   `protobuf:"varint,7,opt,name=policy,proto3" json:"policy,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{0}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return m.Size()
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Permission) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Permission) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Permission) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Permission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Permission) GetAuth() bool {
	if m != nil {
		return m.Auth
	}
	return false
}

func (m *Permission) GetPolicy() bool {
	if m != nil {
		return m.Policy
	}
	return false
}

func (m *Permission) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Permission) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort  string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Where string `protobuf:"bytes,4,opt,name=where,proto3" json:"where,omitempty"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{1}
}
func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return m.Size()
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

type Request struct {
	ListQuery   *ListQuery    `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Permission  *Permission   `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Request) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Response struct {
	Valid       bool          `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Total       int64         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Permission  *Permission   `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	Permissions []*Permission `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Response) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*Permission)(nil), "permission.Permission")
	proto.RegisterType((*ListQuery)(nil), "permission.ListQuery")
	proto.RegisterType((*Request)(nil), "permission.Request")
	proto.RegisterType((*Response)(nil), "permission.Response")
}

func init() { proto.RegisterFile("proto/permission/permission.proto", fileDescriptor_fe0f260493db3286) }

var fileDescriptor_fe0f260493db3286 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0xad, 0x9b, 0x5e, 0xdb, 0x7c, 0x91, 0x18, 0xcc, 0x71, 0xb2, 0x90, 0x88, 0x42, 0xa6, 0x4e,
	0x87, 0x94, 0x03, 0xc4, 0x5a, 0x18, 0x58, 0x18, 0xc0, 0x88, 0xb9, 0x0a, 0x89, 0x45, 0x2d, 0xb9,
	0x71, 0xce, 0x76, 0x0f, 0xf5, 0x5f, 0xf0, 0x33, 0x58, 0xf9, 0x17, 0x2c, 0x48, 0x37, 0x32, 0xa2,
	0x76, 0xe7, 0x37, 0x20, 0x3b, 0x6e, 0x63, 0x18, 0x10, 0x61, 0xfb, 0xde, 0x7b, 0x7e, 0xf1, 0x7b,
	0xb6, 0x15, 0x78, 0xd8, 0x2a, 0x69, 0xe4, 0xa3, 0x96, 0xa9, 0x0d, 0xd7, 0x9a, 0xcb, 0x26, 0x18,
	0x2f, 0x9d, 0x86, 0xa1, 0x67, 0xf2, 0x9f, 0x08, 0xe0, 0xf5, 0x09, 0xe2, 0x3b, 0x30, 0xe6, 0x35,
	0x41, 0x19, 0x5a, 0x44, 0x74, 0xcc, 0x6b, 0x4c, 0x60, 0xa6, 0x99, 0xba, 0xe1, 0x15, 0x23, 0xe3,
	0x0c, 0x2d, 0x62, 0x7a, 0x84, 0xf8, 0x02, 0xa6, 0x1b, 0x66, 0xd6, 0xb2, 0x26, 0x91, 0x13, 0x3c,
	0xc2, 0x18, 0x26, 0x4d, 0xb9, 0x61, 0x64, 0xe2, 0x58, 0x37, 0xe3, 0x0c, 0x92, 0x9a, 0xe9, 0x4a,
	0xf1, 0xd6, 0x70, 0xd9, 0x90, 0x33, 0x27, 0x85, 0x94, 0x75, 0x95, 0x5b, 0xb3, 0x26, 0xd3, 0x0c,
	0x2d, 0xe6, 0xd4, 0xcd, 0x76, 0x87, 0x56, 0x0a, 0x5e, 0xed, 0xc8, 0xcc, 0xb1, 0x1e, 0xe1, 0x07,
	0x00, 0x95, 0x62, 0xa5, 0x61, 0xf5, 0xaa, 0x34, 0x64, 0xee, 0x3e, 0x16, 0x7b, 0x66, 0x69, 0xac,
	0xbc, 0x6d, 0xeb, 0xa3, 0x1c, 0x77, 0xb2, 0x67, 0x96, 0x26, 0x5f, 0x41, 0xfc, 0x8a, 0x6b, 0xf3,
	0x66, 0xcb, 0xd4, 0x0e, 0x9f, 0xc3, 0x99, 0xe0, 0x1b, 0x6e, 0x7c, 0xe3, 0x0e, 0xd8, 0x30, 0x6d,
	0xf9, 0xa1, 0x6b, 0x1c, 0x51, 0x37, 0x5b, 0x4e, 0x4b, 0x65, 0x7c, 0x59, 0x37, 0x5b, 0xf7, 0xc7,
	0x35, 0x53, 0xc7, 0xae, 0x1d, 0xc8, 0xbf, 0x20, 0x98, 0x51, 0x76, 0xbd, 0x65, 0xda, 0xe0, 0xc7,
	0x00, 0x82, 0x6b, 0xb3, 0xba, 0xb6, 0xbb, 0xb9, 0x4d, 0x92, 0xe2, 0xde, 0x65, 0x70, 0x21, 0xa7,
	0x28, 0x34, 0x16, 0xa7, 0x54, 0x4f, 0x21, 0xb8, 0x21, 0x97, 0x22, 0x29, 0x2e, 0x42, 0x57, 0x7f,
	0x61, 0x34, 0x58, 0x89, 0x9f, 0x41, 0xd2, 0x23, 0x4d, 0xa2, 0x2c, 0xfa, 0x8b, 0x31, 0x5c, 0x9a,
	0x7f, 0x46, 0x30, 0xa7, 0x4c, 0xb7, 0xb2, 0xd1, 0xcc, 0xd6, 0xba, 0x29, 0x85, 0x7f, 0x06, 0x73,
	0xda, 0x01, 0xcb, 0x1a, 0x69, 0x4a, 0xe1, 0x4f, 0xa5, 0x03, 0x7f, 0x44, 0x8d, 0xfe, 0x37, 0xea,
	0xe4, 0x9f, 0xa3, 0x16, 0xdf, 0xc6, 0x90, 0xf4, 0x9a, 0xc6, 0x05, 0x44, 0x4b, 0x21, 0xf0, 0xdd,
	0xd0, 0xeb, 0x8f, 0xff, 0xfe, 0xf9, 0xef, 0x64, 0xd7, 0x2f, 0x1f, 0xe1, 0x2b, 0x98, 0xd8, 0x83,
	0x1f, 0x66, 0x2a, 0x20, 0x7a, 0xc9, 0x06, 0x7a, 0x9e, 0xc0, 0xf4, 0x85, 0x7b, 0x98, 0x83, 0x6d,
	0xef, 0xdc, 0x83, 0x1d, 0x5c, 0xeb, 0xed, 0xae, 0xa9, 0x06, 0x99, 0x9e, 0x93, 0xaf, 0xfb, 0x14,
	0xdd, 0xee, 0x53, 0xf4, 0x63, 0x9f, 0xa2, 0x4f, 0x87, 0x74, 0x74, 0x7b, 0x48, 0x47, 0xdf, 0x0f,
	0xe9, 0xe8, 0xfd, 0xd4, 0xfd, 0x2d, 0xae, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x35, 0xfa, 0xdc,
	0xa8, 0x52, 0x04, 0x00, 0x00,
}

func (m *Permission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Permission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x42
	}
	if m.Policy {
		i--
		if m.Policy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Auth {
		i--
		if m.Auth {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPermission(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Where) > 0 {
		i -= len(m.Where)
		copy(dAtA[i:], m.Where)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Where)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sort) > 0 {
		i -= len(m.Sort)
		copy(dAtA[i:], m.Sort)
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Sort)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Page != 0 {
		i = encodeVarintPermission(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if m.Limit != 0 {
		i = encodeVarintPermission(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPermission(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Permission != nil {
		{
			size, err := m.Permission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPermission(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ListQuery != nil {
		{
			size, err := m.ListQuery.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPermission(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPermission(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Permission != nil {
		{
			size, err := m.Permission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPermission(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Total != 0 {
		i = encodeVarintPermission(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPermission(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Permission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPermission(uint64(m.Id))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.Auth {
		n += 2
	}
	if m.Policy {
		n += 2
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	return n
}

func (m *ListQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovPermission(uint64(m.Limit))
	}
	if m.Page != 0 {
		n += 1 + sovPermission(uint64(m.Page))
	}
	l = len(m.Sort)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.Where)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListQuery != nil {
		l = m.ListQuery.Size()
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.Permission != nil {
		l = m.Permission.Size()
		n += 1 + l + sovPermission(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovPermission(uint64(l))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	if m.Total != 0 {
		n += 1 + sovPermission(uint64(m.Total))
	}
	if m.Permission != nil {
		l = m.Permission.Size()
		n += 1 + l + sovPermission(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovPermission(uint64(l))
		}
	}
	return n
}

func sovPermission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermission(x uint64) (n int) {
	return sovPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Permission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Permission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auth", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Auth = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Policy = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Where", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Where = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListQuery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListQuery == nil {
				m.ListQuery = &ListQuery{}
			}
			if err := m.ListQuery.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permission == nil {
				m.Permission = &Permission{}
			}
			if err := m.Permission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &Permission{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permission == nil {
				m.Permission = &Permission{}
			}
			if err := m.Permission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &Permission{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPermission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPermission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPermission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPermission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPermission = fmt.Errorf("proto: unexpected end of group")
)
