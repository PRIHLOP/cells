// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: object.proto

package object

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

// Api Endpoints for ObjectsEndpoint service

func NewObjectsEndpointEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ObjectsEndpoint service

type ObjectsEndpointService interface {
	GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, opts ...client.CallOption) (*GetMinioConfigResponse, error)
	StorageStats(ctx context.Context, in *StorageStatsRequest, opts ...client.CallOption) (*StorageStatsResponse, error)
}

type objectsEndpointService struct {
	c    client.Client
	name string
}

func NewObjectsEndpointService(name string, c client.Client) ObjectsEndpointService {
	return &objectsEndpointService{
		c:    c,
		name: name,
	}
}

func (c *objectsEndpointService) GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, opts ...client.CallOption) (*GetMinioConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ObjectsEndpoint.GetMinioConfig", in)
	out := new(GetMinioConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectsEndpointService) StorageStats(ctx context.Context, in *StorageStatsRequest, opts ...client.CallOption) (*StorageStatsResponse, error) {
	req := c.c.NewRequest(c.name, "ObjectsEndpoint.StorageStats", in)
	out := new(StorageStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ObjectsEndpoint service

type ObjectsEndpointHandler interface {
	GetMinioConfig(context.Context, *GetMinioConfigRequest, *GetMinioConfigResponse) error
	StorageStats(context.Context, *StorageStatsRequest, *StorageStatsResponse) error
}

func RegisterObjectsEndpointHandler(s server.Server, hdlr ObjectsEndpointHandler, opts ...server.HandlerOption) error {
	type objectsEndpoint interface {
		GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, out *GetMinioConfigResponse) error
		StorageStats(ctx context.Context, in *StorageStatsRequest, out *StorageStatsResponse) error
	}
	type ObjectsEndpoint struct {
		objectsEndpoint
	}
	h := &objectsEndpointHandler{hdlr}
	return s.Handle(s.NewHandler(&ObjectsEndpoint{h}, opts...))
}

type objectsEndpointHandler struct {
	ObjectsEndpointHandler
}

func (h *objectsEndpointHandler) GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, out *GetMinioConfigResponse) error {
	return h.ObjectsEndpointHandler.GetMinioConfig(ctx, in, out)
}

func (h *objectsEndpointHandler) StorageStats(ctx context.Context, in *StorageStatsRequest, out *StorageStatsResponse) error {
	return h.ObjectsEndpointHandler.StorageStats(ctx, in, out)
}

// Api Endpoints for DataSourceEndpoint service

func NewDataSourceEndpointEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DataSourceEndpoint service

type DataSourceEndpointService interface {
	GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, opts ...client.CallOption) (*GetDataSourceConfigResponse, error)
}

type dataSourceEndpointService struct {
	c    client.Client
	name string
}

func NewDataSourceEndpointService(name string, c client.Client) DataSourceEndpointService {
	return &dataSourceEndpointService{
		c:    c,
		name: name,
	}
}

func (c *dataSourceEndpointService) GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, opts ...client.CallOption) (*GetDataSourceConfigResponse, error) {
	req := c.c.NewRequest(c.name, "DataSourceEndpoint.GetDataSourceConfig", in)
	out := new(GetDataSourceConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataSourceEndpoint service

type DataSourceEndpointHandler interface {
	GetDataSourceConfig(context.Context, *GetDataSourceConfigRequest, *GetDataSourceConfigResponse) error
}

func RegisterDataSourceEndpointHandler(s server.Server, hdlr DataSourceEndpointHandler, opts ...server.HandlerOption) error {
	type dataSourceEndpoint interface {
		GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, out *GetDataSourceConfigResponse) error
	}
	type DataSourceEndpoint struct {
		dataSourceEndpoint
	}
	h := &dataSourceEndpointHandler{hdlr}
	return s.Handle(s.NewHandler(&DataSourceEndpoint{h}, opts...))
}

type dataSourceEndpointHandler struct {
	DataSourceEndpointHandler
}

func (h *dataSourceEndpointHandler) GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, out *GetDataSourceConfigResponse) error {
	return h.DataSourceEndpointHandler.GetDataSourceConfig(ctx, in, out)
}

// Api Endpoints for ResourceCleanerEndpoint service

func NewResourceCleanerEndpointEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ResourceCleanerEndpoint service

type ResourceCleanerEndpointService interface {
	CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, opts ...client.CallOption) (*CleanResourcesResponse, error)
}

type resourceCleanerEndpointService struct {
	c    client.Client
	name string
}

func NewResourceCleanerEndpointService(name string, c client.Client) ResourceCleanerEndpointService {
	return &resourceCleanerEndpointService{
		c:    c,
		name: name,
	}
}

func (c *resourceCleanerEndpointService) CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, opts ...client.CallOption) (*CleanResourcesResponse, error) {
	req := c.c.NewRequest(c.name, "ResourceCleanerEndpoint.CleanResourcesBeforeDelete", in)
	out := new(CleanResourcesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceCleanerEndpoint service

type ResourceCleanerEndpointHandler interface {
	CleanResourcesBeforeDelete(context.Context, *CleanResourcesRequest, *CleanResourcesResponse) error
}

func RegisterResourceCleanerEndpointHandler(s server.Server, hdlr ResourceCleanerEndpointHandler, opts ...server.HandlerOption) error {
	type resourceCleanerEndpoint interface {
		CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, out *CleanResourcesResponse) error
	}
	type ResourceCleanerEndpoint struct {
		resourceCleanerEndpoint
	}
	h := &resourceCleanerEndpointHandler{hdlr}
	return s.Handle(s.NewHandler(&ResourceCleanerEndpoint{h}, opts...))
}

type resourceCleanerEndpointHandler struct {
	ResourceCleanerEndpointHandler
}

func (h *resourceCleanerEndpointHandler) CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, out *CleanResourcesResponse) error {
	return h.ResourceCleanerEndpointHandler.CleanResourcesBeforeDelete(ctx, in, out)
}