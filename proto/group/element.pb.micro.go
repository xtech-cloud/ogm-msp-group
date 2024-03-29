// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/group/element.proto

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

// Api Endpoints for Element service

func NewElementEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Element service

type ElementService interface {
	// 加入一个元素
	Add(ctx context.Context, in *ElementAddRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 更新一个元素
	Update(ctx context.Context, in *ElementUpdateRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 列举一个集合的元素
	List(ctx context.Context, in *ElementListRequest, opts ...client.CallOption) (*ElementListResponse, error)
	// 搜索一个集合中的元素
	Search(ctx context.Context, in *ElementSearchRequest, opts ...client.CallOption) (*ElementListResponse, error)
	// 移除一个元素
	Remove(ctx context.Context, in *ElementRemoveRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取一个元素信息
	Get(ctx context.Context, in *ElementGetRequest, opts ...client.CallOption) (*ElementGetResponse, error)
	// 获取所在的集合
	Where(ctx context.Context, in *ElementWhereRequest, opts ...client.CallOption) (*ElementWhereResponse, error)
}

type elementService struct {
	c    client.Client
	name string
}

func NewElementService(name string, c client.Client) ElementService {
	return &elementService{
		c:    c,
		name: name,
	}
}

func (c *elementService) Add(ctx context.Context, in *ElementAddRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Add", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) Update(ctx context.Context, in *ElementUpdateRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Update", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) List(ctx context.Context, in *ElementListRequest, opts ...client.CallOption) (*ElementListResponse, error) {
	req := c.c.NewRequest(c.name, "Element.List", in)
	out := new(ElementListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) Search(ctx context.Context, in *ElementSearchRequest, opts ...client.CallOption) (*ElementListResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Search", in)
	out := new(ElementListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) Remove(ctx context.Context, in *ElementRemoveRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Remove", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) Get(ctx context.Context, in *ElementGetRequest, opts ...client.CallOption) (*ElementGetResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Get", in)
	out := new(ElementGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elementService) Where(ctx context.Context, in *ElementWhereRequest, opts ...client.CallOption) (*ElementWhereResponse, error) {
	req := c.c.NewRequest(c.name, "Element.Where", in)
	out := new(ElementWhereResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Element service

type ElementHandler interface {
	// 加入一个元素
	Add(context.Context, *ElementAddRequest, *UuidResponse) error
	// 更新一个元素
	Update(context.Context, *ElementUpdateRequest, *UuidResponse) error
	// 列举一个集合的元素
	List(context.Context, *ElementListRequest, *ElementListResponse) error
	// 搜索一个集合中的元素
	Search(context.Context, *ElementSearchRequest, *ElementListResponse) error
	// 移除一个元素
	Remove(context.Context, *ElementRemoveRequest, *UuidResponse) error
	// 获取一个元素信息
	Get(context.Context, *ElementGetRequest, *ElementGetResponse) error
	// 获取所在的集合
	Where(context.Context, *ElementWhereRequest, *ElementWhereResponse) error
}

func RegisterElementHandler(s server.Server, hdlr ElementHandler, opts ...server.HandlerOption) error {
	type element interface {
		Add(ctx context.Context, in *ElementAddRequest, out *UuidResponse) error
		Update(ctx context.Context, in *ElementUpdateRequest, out *UuidResponse) error
		List(ctx context.Context, in *ElementListRequest, out *ElementListResponse) error
		Search(ctx context.Context, in *ElementSearchRequest, out *ElementListResponse) error
		Remove(ctx context.Context, in *ElementRemoveRequest, out *UuidResponse) error
		Get(ctx context.Context, in *ElementGetRequest, out *ElementGetResponse) error
		Where(ctx context.Context, in *ElementWhereRequest, out *ElementWhereResponse) error
	}
	type Element struct {
		element
	}
	h := &elementHandler{hdlr}
	return s.Handle(s.NewHandler(&Element{h}, opts...))
}

type elementHandler struct {
	ElementHandler
}

func (h *elementHandler) Add(ctx context.Context, in *ElementAddRequest, out *UuidResponse) error {
	return h.ElementHandler.Add(ctx, in, out)
}

func (h *elementHandler) Update(ctx context.Context, in *ElementUpdateRequest, out *UuidResponse) error {
	return h.ElementHandler.Update(ctx, in, out)
}

func (h *elementHandler) List(ctx context.Context, in *ElementListRequest, out *ElementListResponse) error {
	return h.ElementHandler.List(ctx, in, out)
}

func (h *elementHandler) Search(ctx context.Context, in *ElementSearchRequest, out *ElementListResponse) error {
	return h.ElementHandler.Search(ctx, in, out)
}

func (h *elementHandler) Remove(ctx context.Context, in *ElementRemoveRequest, out *UuidResponse) error {
	return h.ElementHandler.Remove(ctx, in, out)
}

func (h *elementHandler) Get(ctx context.Context, in *ElementGetRequest, out *ElementGetResponse) error {
	return h.ElementHandler.Get(ctx, in, out)
}

func (h *elementHandler) Where(ctx context.Context, in *ElementWhereRequest, out *ElementWhereResponse) error {
	return h.ElementHandler.Where(ctx, in, out)
}
