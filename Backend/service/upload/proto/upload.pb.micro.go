// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: upload.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UploadService service

func NewUploadServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UploadService service

type UploadService interface {
	// 获取上传入口地址
	UploadEntry(ctx context.Context, in *ReqEntry, opts ...client.CallOption) (*ResEntry, error)
}

type uploadService struct {
	c    client.Client
	name string
}

func NewUploadService(name string, c client.Client) UploadService {
	return &uploadService{
		c:    c,
		name: name,
	}
}

func (c *uploadService) UploadEntry(ctx context.Context, in *ReqEntry, opts ...client.CallOption) (*ResEntry, error) {
	req := c.c.NewRequest(c.name, "UploadService.UploadEntry", in)
	out := new(ResEntry)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UploadService service

type UploadServiceHandler interface {
	// 获取上传入口地址
	UploadEntry(context.Context, *ReqEntry, *ResEntry) error
}

func RegisterUploadServiceHandler(s server.Server, hdlr UploadServiceHandler, opts ...server.HandlerOption) error {
	type uploadService interface {
		UploadEntry(ctx context.Context, in *ReqEntry, out *ResEntry) error
	}
	type UploadService struct {
		uploadService
	}
	h := &uploadServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UploadService{h}, opts...))
}

type uploadServiceHandler struct {
	UploadServiceHandler
}

func (h *uploadServiceHandler) UploadEntry(ctx context.Context, in *ReqEntry, out *ResEntry) error {
	return h.UploadServiceHandler.UploadEntry(ctx, in, out)
}
