// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package daymenu_shipping_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// 用户
type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Status               int64    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Token                string   `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Mobile               string   `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

// 请求
type Request struct {
	PageSize             int64    `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Id                   string   `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Request) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Request) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// 响应
type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	Token                *Token   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Count                int64    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Code                 int64    `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Response) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

// token
type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Descriptioin         string   `protobuf:"bytes,2,opt,name=descriptioin,proto3" json:"descriptioin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescriptioin() string {
	if m != nil {
		return m.Descriptioin
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "daymenu.shipping.srv.user.User")
	proto.RegisterType((*Request)(nil), "daymenu.shipping.srv.user.Request")
	proto.RegisterType((*Response)(nil), "daymenu.shipping.srv.user.Response")
	proto.RegisterType((*Token)(nil), "daymenu.shipping.srv.user.Token")
	proto.RegisterType((*Error)(nil), "daymenu.shipping.srv.user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x7f, 0xd6, 0x09, 0x53, 0xe0, 0xb0, 0x0a, 0x68, 0xe9, 0x85, 0xc8, 0x5c, 0x7a, 0x32,
	0x52, 0x2b, 0x10, 0x37, 0x0e, 0x08, 0x21, 0xa4, 0x1e, 0x8a, 0xd3, 0x70, 0x77, 0xed, 0x21, 0xac,
	0x88, 0x77, 0xcd, 0xee, 0x3a, 0xa8, 0xbc, 0x02, 0x37, 0x9e, 0x87, 0x87, 0xe0, 0x91, 0xd0, 0x8e,
	0xed, 0x34, 0x1c, 0x88, 0x23, 0xd4, 0x4b, 0x32, 0xdf, 0xe4, 0x9b, 0x99, 0x6f, 0x7e, 0x36, 0xf0,
	0xa8, 0x31, 0xda, 0xe9, 0xe7, 0xad, 0x45, 0x43, 0x1f, 0x19, 0x61, 0xfe, 0xa4, 0x2a, 0x6e, 0x6a,
	0x54, 0x6d, 0x66, 0x3f, 0xcb, 0xa6, 0x91, 0x6a, 0x95, 0x59, 0xb3, 0xc9, 0x3c, 0x21, 0xfd, 0x15,
	0x40, 0xbc, 0xb4, 0x68, 0xf8, 0x43, 0x08, 0x65, 0x25, 0x82, 0x79, 0x70, 0x1a, 0xe5, 0xa1, 0xac,
	0x38, 0x87, 0x58, 0x15, 0x35, 0x8a, 0x70, 0x1e, 0x9c, 0xde, 0xcb, 0xc9, 0xe6, 0x02, 0x26, 0xa5,
	0xae, 0x9b, 0x42, 0xdd, 0x88, 0x88, 0xdc, 0x03, 0xe4, 0x33, 0x60, 0x58, 0x17, 0x72, 0x2d, 0x62,
	0xf2, 0x77, 0x80, 0x9f, 0xc0, 0xb4, 0x29, 0xac, 0xfd, 0xa6, 0x4d, 0x25, 0x18, 0xfd, 0xb0, 0xc5,
	0xfc, 0x31, 0x24, 0xd6, 0x15, 0xae, 0xb5, 0x22, 0xa1, 0x9a, 0x3d, 0xf2, 0x99, 0x9c, 0xfe, 0x82,
	0x4a, 0x4c, 0xba, 0x4c, 0x04, 0x3c, 0xbb, 0xd6, 0xd7, 0x72, 0x8d, 0x62, 0x4a, 0xee, 0x1e, 0xa5,
	0x3f, 0x02, 0x98, 0xe4, 0xf8, 0xb5, 0x45, 0xeb, 0xba, 0x6a, 0x2b, 0x5c, 0xc8, 0xef, 0xd8, 0xf7,
	0xb1, 0xc5, 0xbe, 0x1b, 0x6f, 0x53, 0x37, 0x51, 0x4e, 0xf6, 0xad, 0xe6, 0x68, 0x57, 0xf3, 0xd0,
	0x77, 0xbc, 0xd3, 0xf7, 0x6d, 0x75, 0xb6, 0x5b, 0xbd, 0x9f, 0x59, 0x42, 0xbe, 0x50, 0x56, 0xe9,
	0xcf, 0x10, 0xa6, 0x39, 0xda, 0x46, 0x2b, 0x8b, 0xfc, 0x1c, 0x62, 0x3f, 0x61, 0x92, 0x72, 0x7c,
	0xf6, 0x34, 0xfb, 0xe7, 0x0e, 0x32, 0x3f, 0xff, 0x9c, 0xc8, 0xfc, 0x05, 0x30, 0xff, 0x6d, 0x45,
	0x38, 0x8f, 0x0e, 0x89, 0xea, 0xd8, 0xfc, 0x15, 0x24, 0x68, 0x8c, 0x36, 0x56, 0x44, 0x14, 0x37,
	0xdf, 0x13, 0xf7, 0xd6, 0x13, 0xf3, 0x9e, 0xcf, 0x5f, 0x0e, 0xe3, 0x8e, 0x49, 0xe6, 0xbe, 0xc0,
	0x2b, 0xcf, 0x1b, 0x16, 0x32, 0x03, 0x56, 0xea, 0x56, 0x39, 0x9a, 0x48, 0x94, 0x77, 0xc0, 0x0f,
	0xaf, 0xd4, 0x15, 0xf6, 0x2b, 0x25, 0x3b, 0xad, 0x81, 0x5d, 0x0d, 0x21, 0x5d, 0xa9, 0x60, 0x77,
	0xb3, 0x33, 0x60, 0x9b, 0x62, 0x2d, 0x2b, 0x5a, 0xcd, 0x34, 0xef, 0xc0, 0xff, 0x37, 0x94, 0xbe,
	0x06, 0x46, 0x8e, 0xad, 0x16, 0x5f, 0x8d, 0x75, 0x5a, 0x78, 0x0a, 0xf7, 0x2b, 0xb4, 0xa5, 0x91,
	0x8d, 0x93, 0x5a, 0xaa, 0xfe, 0xb8, 0xff, 0xf2, 0x9d, 0xfd, 0x8e, 0xe1, 0xd8, 0xcf, 0x76, 0x81,
	0x66, 0x23, 0x4b, 0xe4, 0x97, 0x90, 0xbc, 0x31, 0x58, 0x38, 0xe4, 0x63, 0xdb, 0x38, 0x79, 0xb6,
	0x87, 0x30, 0xdc, 0x45, 0x7a, 0xe4, 0x33, 0x2e, 0x9b, 0xea, 0x6e, 0x33, 0x46, 0xef, 0xd0, 0xf1,
	0x74, 0x2f, 0x9b, 0x5e, 0xc9, 0xa1, 0x19, 0x3f, 0x40, 0x7c, 0xe9, 0x1f, 0xc9, 0x1d, 0xa6, 0xbc,
	0x00, 0x76, 0xa1, 0x57, 0x52, 0x8d, 0x77, 0x3d, 0x7a, 0x85, 0xe9, 0x11, 0x5f, 0xc0, 0xd4, 0x73,
	0xdf, 0xab, 0x4f, 0x9a, 0x8f, 0xf2, 0x0f, 0x95, 0xb8, 0x84, 0x07, 0x1f, 0xfd, 0xfd, 0x15, 0x0e,
	0xbb, 0x9b, 0x1d, 0xcf, 0x7c, 0x80, 0xd6, 0xeb, 0x84, 0xfe, 0x86, 0xcf, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x42, 0xac, 0x18, 0x6c, 0x9f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Page(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Login(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	UserInfo(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "daymenu.shipping.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Page(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Page", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Login", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UserInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Update(context.Context, *User, *Response) error
	Get(context.Context, *Request, *Response) error
	Page(context.Context, *Request, *Response) error
	Login(context.Context, *User, *Token) error
	UserInfo(context.Context, *Token, *Response) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Update(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) Page(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Page(ctx, in, out)
}

func (h *UserService) Login(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *UserService) UserInfo(ctx context.Context, in *Token, out *Response) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
