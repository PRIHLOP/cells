// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: install.proto

package install

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Install service

func NewInstallEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Install service

type InstallService interface {
	GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...client.CallOption) (*GetDefaultsResponse, error)
	Install(ctx context.Context, in *InstallRequest, opts ...client.CallOption) (*InstallResponse, error)
	PerformCheck(ctx context.Context, in *PerformCheckRequest, opts ...client.CallOption) (*PerformCheckResponse, error)
}

type installService struct {
	c    client.Client
	name string
}

func NewInstallService(name string, c client.Client) InstallService {
	return &installService{
		c:    c,
		name: name,
	}
}

func (c *installService) GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...client.CallOption) (*GetDefaultsResponse, error) {
	req := c.c.NewRequest(c.name, "Install.GetDefaults", in)
	out := new(GetDefaultsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installService) Install(ctx context.Context, in *InstallRequest, opts ...client.CallOption) (*InstallResponse, error) {
	req := c.c.NewRequest(c.name, "Install.Install", in)
	out := new(InstallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installService) PerformCheck(ctx context.Context, in *PerformCheckRequest, opts ...client.CallOption) (*PerformCheckResponse, error) {
	req := c.c.NewRequest(c.name, "Install.PerformCheck", in)
	out := new(PerformCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Install service

type InstallHandler interface {
	GetDefaults(context.Context, *GetDefaultsRequest, *GetDefaultsResponse) error
	Install(context.Context, *InstallRequest, *InstallResponse) error
	PerformCheck(context.Context, *PerformCheckRequest, *PerformCheckResponse) error
}

func RegisterInstallHandler(s server.Server, hdlr InstallHandler, opts ...server.HandlerOption) error {
	type install interface {
		GetDefaults(ctx context.Context, in *GetDefaultsRequest, out *GetDefaultsResponse) error
		Install(ctx context.Context, in *InstallRequest, out *InstallResponse) error
		PerformCheck(ctx context.Context, in *PerformCheckRequest, out *PerformCheckResponse) error
	}
	type Install struct {
		install
	}
	h := &installHandler{hdlr}
	return s.Handle(s.NewHandler(&Install{h}, opts...))
}

type installHandler struct {
	InstallHandler
}

func (h *installHandler) GetDefaults(ctx context.Context, in *GetDefaultsRequest, out *GetDefaultsResponse) error {
	return h.InstallHandler.GetDefaults(ctx, in, out)
}

func (h *installHandler) Install(ctx context.Context, in *InstallRequest, out *InstallResponse) error {
	return h.InstallHandler.Install(ctx, in, out)
}

func (h *installHandler) PerformCheck(ctx context.Context, in *PerformCheckRequest, out *PerformCheckResponse) error {
	return h.InstallHandler.PerformCheck(ctx, in, out)
}