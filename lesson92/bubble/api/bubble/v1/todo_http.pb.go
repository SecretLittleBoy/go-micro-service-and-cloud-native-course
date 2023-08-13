// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.20.1
// source: bubble/v1/todo.proto

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

const OperationTodoCreateTodo = "/api.bubble.v1.Todo/CreateTodo"
const OperationTodoDeleteTodo = "/api.bubble.v1.Todo/DeleteTodo"
const OperationTodoGetTodo = "/api.bubble.v1.Todo/GetTodo"
const OperationTodoListTodo = "/api.bubble.v1.Todo/ListTodo"
const OperationTodoUpdateTodo = "/api.bubble.v1.Todo/UpdateTodo"

type TodoHTTPServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoReply, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoReply, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoReply, error)
	ListTodo(context.Context, *ListTodoRequest) (*ListTodoReply, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoReply, error)
}

func RegisterTodoHTTPServer(s *http.Server, srv TodoHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/todo", _Todo_CreateTodo0_HTTP_Handler(srv))
	r.PUT("/v1/todo/{id}", _Todo_UpdateTodo0_HTTP_Handler(srv))
	r.DELETE("/v1/todo/{id}", _Todo_DeleteTodo0_HTTP_Handler(srv))
	r.GET("/v1/todo/{id}", _Todo_GetTodo0_HTTP_Handler(srv))
	r.GET("/v1/todos", _Todo_ListTodo0_HTTP_Handler(srv))
}

func _Todo_CreateTodo0_HTTP_Handler(srv TodoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTodoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTodoCreateTodo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTodo(ctx, req.(*CreateTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todo_UpdateTodo0_HTTP_Handler(srv TodoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTodoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTodoUpdateTodo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTodo(ctx, req.(*UpdateTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todo_DeleteTodo0_HTTP_Handler(srv TodoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTodoDeleteTodo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTodo(ctx, req.(*DeleteTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todo_GetTodo0_HTTP_Handler(srv TodoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTodoGetTodo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTodo(ctx, req.(*GetTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todo_ListTodo0_HTTP_Handler(srv TodoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTodoListTodo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTodo(ctx, req.(*ListTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTodoReply)
		return ctx.Result(200, reply)
	}
}

type TodoHTTPClient interface {
	CreateTodo(ctx context.Context, req *CreateTodoRequest, opts ...http.CallOption) (rsp *CreateTodoReply, err error)
	DeleteTodo(ctx context.Context, req *DeleteTodoRequest, opts ...http.CallOption) (rsp *DeleteTodoReply, err error)
	GetTodo(ctx context.Context, req *GetTodoRequest, opts ...http.CallOption) (rsp *GetTodoReply, err error)
	ListTodo(ctx context.Context, req *ListTodoRequest, opts ...http.CallOption) (rsp *ListTodoReply, err error)
	UpdateTodo(ctx context.Context, req *UpdateTodoRequest, opts ...http.CallOption) (rsp *UpdateTodoReply, err error)
}

type TodoHTTPClientImpl struct {
	cc *http.Client
}

func NewTodoHTTPClient(client *http.Client) TodoHTTPClient {
	return &TodoHTTPClientImpl{client}
}

func (c *TodoHTTPClientImpl) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...http.CallOption) (*CreateTodoReply, error) {
	var out CreateTodoReply
	pattern := "/v1/todo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTodoCreateTodo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodoHTTPClientImpl) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...http.CallOption) (*DeleteTodoReply, error) {
	var out DeleteTodoReply
	pattern := "/v1/todo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTodoDeleteTodo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodoHTTPClientImpl) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...http.CallOption) (*GetTodoReply, error) {
	var out GetTodoReply
	pattern := "/v1/todo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTodoGetTodo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodoHTTPClientImpl) ListTodo(ctx context.Context, in *ListTodoRequest, opts ...http.CallOption) (*ListTodoReply, error) {
	var out ListTodoReply
	pattern := "/v1/todos"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTodoListTodo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodoHTTPClientImpl) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...http.CallOption) (*UpdateTodoReply, error) {
	var out UpdateTodoReply
	pattern := "/v1/todo/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTodoUpdateTodo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
