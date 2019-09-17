// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetSmsCd/GetSmsCd.proto

package go_micro_srv_GetSmsCd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GetSmsCd service

type GetSmsCdService interface {
	CallGetSmsCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type getSmsCdService struct {
	c    client.Client
	name string
}

func NewGetSmsCdService(name string, c client.Client) GetSmsCdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.GetSmsCd"
	}
	return &getSmsCdService{
		c:    c,
		name: name,
	}
}

func (c *getSmsCdService) CallGetSmsCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetSmsCd.CallGetSmsCd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetSmsCd service

type GetSmsCdHandler interface {
	CallGetSmsCd(context.Context, *Request, *Response) error
}

func RegisterGetSmsCdHandler(s server.Server, hdlr GetSmsCdHandler, opts ...server.HandlerOption) error {
	type getSmsCd interface {
		CallGetSmsCd(ctx context.Context, in *Request, out *Response) error
	}
	type GetSmsCd struct {
		getSmsCd
	}
	h := &getSmsCdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetSmsCd{h}, opts...))
}

type getSmsCdHandler struct {
	GetSmsCdHandler
}

func (h *getSmsCdHandler) CallGetSmsCd(ctx context.Context, in *Request, out *Response) error {
	return h.GetSmsCdHandler.CallGetSmsCd(ctx, in, out)
}
