// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/group/collection.proto

package group

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

// Api Endpoints for Collection service

func NewCollectionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Collection service

type CollectionService interface {
	// 创建一个集合
	Make(ctx context.Context, in *CollectionMakeRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 更新一个集合
	Update(ctx context.Context, in *CollectionUpdateRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 列举集合
	List(ctx context.Context, in *CollectionListRequest, opts ...client.CallOption) (*CollectionListResponse, error)
	// 搜索集合
	Search(ctx context.Context, in *CollectionSearchRequest, opts ...client.CallOption) (*CollectionListResponse, error)
	// 删除集合
	Remove(ctx context.Context, in *CollectionRemoveRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取集合信息
	Get(ctx context.Context, in *CollectionGetRequest, opts ...client.CallOption) (*CollectionGetResponse, error)
}

type collectionService struct {
	c    client.Client
	name string
}

func NewCollectionService(name string, c client.Client) CollectionService {
	return &collectionService{
		c:    c,
		name: name,
	}
}

func (c *collectionService) Make(ctx context.Context, in *CollectionMakeRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Make", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Update(ctx context.Context, in *CollectionUpdateRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Update", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) List(ctx context.Context, in *CollectionListRequest, opts ...client.CallOption) (*CollectionListResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.List", in)
	out := new(CollectionListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Search(ctx context.Context, in *CollectionSearchRequest, opts ...client.CallOption) (*CollectionListResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Search", in)
	out := new(CollectionListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Remove(ctx context.Context, in *CollectionRemoveRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Remove", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Get(ctx context.Context, in *CollectionGetRequest, opts ...client.CallOption) (*CollectionGetResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Get", in)
	out := new(CollectionGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collection service

type CollectionHandler interface {
	// 创建一个集合
	Make(context.Context, *CollectionMakeRequest, *UuidResponse) error
	// 更新一个集合
	Update(context.Context, *CollectionUpdateRequest, *UuidResponse) error
	// 列举集合
	List(context.Context, *CollectionListRequest, *CollectionListResponse) error
	// 搜索集合
	Search(context.Context, *CollectionSearchRequest, *CollectionListResponse) error
	// 删除集合
	Remove(context.Context, *CollectionRemoveRequest, *UuidResponse) error
	// 获取集合信息
	Get(context.Context, *CollectionGetRequest, *CollectionGetResponse) error
}

func RegisterCollectionHandler(s server.Server, hdlr CollectionHandler, opts ...server.HandlerOption) error {
	type collection interface {
		Make(ctx context.Context, in *CollectionMakeRequest, out *UuidResponse) error
		Update(ctx context.Context, in *CollectionUpdateRequest, out *UuidResponse) error
		List(ctx context.Context, in *CollectionListRequest, out *CollectionListResponse) error
		Search(ctx context.Context, in *CollectionSearchRequest, out *CollectionListResponse) error
		Remove(ctx context.Context, in *CollectionRemoveRequest, out *UuidResponse) error
		Get(ctx context.Context, in *CollectionGetRequest, out *CollectionGetResponse) error
	}
	type Collection struct {
		collection
	}
	h := &collectionHandler{hdlr}
	return s.Handle(s.NewHandler(&Collection{h}, opts...))
}

type collectionHandler struct {
	CollectionHandler
}

func (h *collectionHandler) Make(ctx context.Context, in *CollectionMakeRequest, out *UuidResponse) error {
	return h.CollectionHandler.Make(ctx, in, out)
}

func (h *collectionHandler) Update(ctx context.Context, in *CollectionUpdateRequest, out *UuidResponse) error {
	return h.CollectionHandler.Update(ctx, in, out)
}

func (h *collectionHandler) List(ctx context.Context, in *CollectionListRequest, out *CollectionListResponse) error {
	return h.CollectionHandler.List(ctx, in, out)
}

func (h *collectionHandler) Search(ctx context.Context, in *CollectionSearchRequest, out *CollectionListResponse) error {
	return h.CollectionHandler.Search(ctx, in, out)
}

func (h *collectionHandler) Remove(ctx context.Context, in *CollectionRemoveRequest, out *UuidResponse) error {
	return h.CollectionHandler.Remove(ctx, in, out)
}

func (h *collectionHandler) Get(ctx context.Context, in *CollectionGetRequest, out *CollectionGetResponse) error {
	return h.CollectionHandler.Get(ctx, in, out)
}
