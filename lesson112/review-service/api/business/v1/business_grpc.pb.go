// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: business/v1/business.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Business_ReplyReview_FullMethodName  = "/api.business.v1.Business/ReplyReview"
	Business_AppealReview_FullMethodName = "/api.business.v1.Business/AppealReview"
)

// BusinessClient is the client API for Business service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessClient interface {
	// 商家回复用户评价
	ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...grpc.CallOption) (*ReplyReviewReply, error)
	// 商家申诉用户评价
	AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...grpc.CallOption) (*AppealReviewReply, error)
}

type businessClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessClient(cc grpc.ClientConnInterface) BusinessClient {
	return &businessClient{cc}
}

func (c *businessClient) ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...grpc.CallOption) (*ReplyReviewReply, error) {
	out := new(ReplyReviewReply)
	err := c.cc.Invoke(ctx, Business_ReplyReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessClient) AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...grpc.CallOption) (*AppealReviewReply, error) {
	out := new(AppealReviewReply)
	err := c.cc.Invoke(ctx, Business_AppealReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServer is the server API for Business service.
// All implementations must embed UnimplementedBusinessServer
// for forward compatibility
type BusinessServer interface {
	// 商家回复用户评价
	ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error)
	// 商家申诉用户评价
	AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error)
	mustEmbedUnimplementedBusinessServer()
}

// UnimplementedBusinessServer must be embedded to have forward compatible implementations.
type UnimplementedBusinessServer struct {
}

func (UnimplementedBusinessServer) ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyReview not implemented")
}
func (UnimplementedBusinessServer) AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppealReview not implemented")
}
func (UnimplementedBusinessServer) mustEmbedUnimplementedBusinessServer() {}

// UnsafeBusinessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServer will
// result in compilation errors.
type UnsafeBusinessServer interface {
	mustEmbedUnimplementedBusinessServer()
}

func RegisterBusinessServer(s grpc.ServiceRegistrar, srv BusinessServer) {
	s.RegisterService(&Business_ServiceDesc, srv)
}

func _Business_ReplyReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).ReplyReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_ReplyReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).ReplyReview(ctx, req.(*ReplyReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_AppealReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppealReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).AppealReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_AppealReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).AppealReview(ctx, req.(*AppealReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Business_ServiceDesc is the grpc.ServiceDesc for Business service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Business_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.business.v1.Business",
	HandlerType: (*BusinessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplyReview",
			Handler:    _Business_ReplyReview_Handler,
		},
		{
			MethodName: "AppealReview",
			Handler:    _Business_AppealReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "business/v1/business.proto",
}
