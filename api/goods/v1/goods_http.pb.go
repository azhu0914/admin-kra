// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type GoodsServiceHTTPServer interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error)
}

func RegisterGoodsServiceHTTPServer(s *http.Server, srv GoodsServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/login", _GoodsService_Login0_HTTP_Handler(srv))
	r.POST("/v1/goods/add", _GoodsService_CreateGoods0_HTTP_Handler(srv))
	r.PUT("/v1/goods/{id}", _GoodsService_UpdateGoods0_HTTP_Handler(srv))
	r.DELETE("/v1/goods/{id}", _GoodsService_DeleteGoods0_HTTP_Handler(srv))
	r.GET("/v1/goods/{id}", _GoodsService_GetGoods0_HTTP_Handler(srv))
	r.GET("/v1/goods/", _GoodsService_ListGoods0_HTTP_Handler(srv))
}

func _GoodsService_Login0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_CreateGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGoodsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/CreateGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGoods(ctx, req.(*CreateGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_UpdateGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGoodsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/UpdateGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGoods(ctx, req.(*UpdateGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_DeleteGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/DeleteGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteGoods(ctx, req.(*DeleteGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_GetGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/GetGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoods(ctx, req.(*GetGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_ListGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/goods.v1.GoodsService/ListGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGoods(ctx, req.(*ListGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGoodsReply)
		return ctx.Result(200, reply)
	}
}

type GoodsServiceHTTPClient interface {
	CreateGoods(ctx context.Context, req *CreateGoodsRequest, opts ...http.CallOption) (rsp *CreateGoodsReply, err error)
	DeleteGoods(ctx context.Context, req *DeleteGoodsRequest, opts ...http.CallOption) (rsp *DeleteGoodsReply, err error)
	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http.CallOption) (rsp *GetGoodsReply, err error)
	ListGoods(ctx context.Context, req *ListGoodsRequest, opts ...http.CallOption) (rsp *ListGoodsReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	UpdateGoods(ctx context.Context, req *UpdateGoodsRequest, opts ...http.CallOption) (rsp *UpdateGoodsReply, err error)
}

type GoodsServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewGoodsServiceHTTPClient(client *http.Client) GoodsServiceHTTPClient {
	return &GoodsServiceHTTPClientImpl{client}
}

func (c *GoodsServiceHTTPClientImpl) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...http.CallOption) (*CreateGoodsReply, error) {
	var out CreateGoodsReply
	pattern := "/v1/goods/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/CreateGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...http.CallOption) (*DeleteGoodsReply, error) {
	var out DeleteGoodsReply
	pattern := "/v1/goods/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/DeleteGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http.CallOption) (*GetGoodsReply, error) {
	var out GetGoodsReply
	pattern := "/v1/goods/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/GetGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...http.CallOption) (*ListGoodsReply, error) {
	var out ListGoodsReply
	pattern := "/v1/goods/"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/ListGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...http.CallOption) (*UpdateGoodsReply, error) {
	var out UpdateGoodsReply
	pattern := "/v1/goods/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/goods.v1.GoodsService/UpdateGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
