// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/crank.proto

package crank

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
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

// Api Endpoints for Crank service

func NewCrankEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Crank service

type CrankService interface {
	CreateCrank(ctx context.Context, in *CreateCrankRequest, opts ...client.CallOption) (*CreateCrankResponse, error)
}

type crankService struct {
	c    client.Client
	name string
}

func NewCrankService(name string, c client.Client) CrankService {
	return &crankService{
		c:    c,
		name: name,
	}
}

func (c *crankService) CreateCrank(ctx context.Context, in *CreateCrankRequest, opts ...client.CallOption) (*CreateCrankResponse, error) {
	req := c.c.NewRequest(c.name, "Crank.CreateCrank", in)
	out := new(CreateCrankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Crank service

type CrankHandler interface {
	CreateCrank(context.Context, *CreateCrankRequest, *CreateCrankResponse) error
}

func RegisterCrankHandler(s server.Server, hdlr CrankHandler, opts ...server.HandlerOption) error {
	type crank interface {
		CreateCrank(ctx context.Context, in *CreateCrankRequest, out *CreateCrankResponse) error
	}
	type Crank struct {
		crank
	}
	h := &crankHandler{hdlr}
	return s.Handle(s.NewHandler(&Crank{h}, opts...))
}

type crankHandler struct {
	CrankHandler
}

func (h *crankHandler) CreateCrank(ctx context.Context, in *CreateCrankRequest, out *CreateCrankResponse) error {
	return h.CrankHandler.CreateCrank(ctx, in, out)
}
