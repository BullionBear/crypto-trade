// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/proto/feed.proto

package feed

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Feed_GetConfig_FullMethodName           = "/feed.Feed/GetConfig"
	Feed_GetStatus_FullMethodName           = "/feed.Feed/GetStatus"
	Feed_GetSubscriber_FullMethodName       = "/feed.Feed/GetSubscriber"
	Feed_SubscribeKline_FullMethodName      = "/feed.Feed/SubscribeKline"
	Feed_ReadHistoricalKline_FullMethodName = "/feed.Feed/ReadHistoricalKline"
)

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConfigResponse, error)
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	GetSubscriber(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SubscriberResponse, error)
	SubscribeKline(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Feed_SubscribeKlineClient, error)
	ReadHistoricalKline(ctx context.Context, in *ReadKlineRequest, opts ...grpc.CallOption) (Feed_ReadHistoricalKlineClient, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, Feed_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Feed_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) GetSubscriber(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SubscriberResponse, error) {
	out := new(SubscriberResponse)
	err := c.cc.Invoke(ctx, Feed_GetSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) SubscribeKline(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Feed_SubscribeKlineClient, error) {
	stream, err := c.cc.NewStream(ctx, &Feed_ServiceDesc.Streams[0], Feed_SubscribeKline_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &feedSubscribeKlineClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Feed_SubscribeKlineClient interface {
	Recv() (*KlineResponse, error)
	grpc.ClientStream
}

type feedSubscribeKlineClient struct {
	grpc.ClientStream
}

func (x *feedSubscribeKlineClient) Recv() (*KlineResponse, error) {
	m := new(KlineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedClient) ReadHistoricalKline(ctx context.Context, in *ReadKlineRequest, opts ...grpc.CallOption) (Feed_ReadHistoricalKlineClient, error) {
	stream, err := c.cc.NewStream(ctx, &Feed_ServiceDesc.Streams[1], Feed_ReadHistoricalKline_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &feedReadHistoricalKlineClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Feed_ReadHistoricalKlineClient interface {
	Recv() (*KlineResponse, error)
	grpc.ClientStream
}

type feedReadHistoricalKlineClient struct {
	grpc.ClientStream
}

func (x *feedReadHistoricalKlineClient) Recv() (*KlineResponse, error) {
	m := new(KlineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	GetConfig(context.Context, *emptypb.Empty) (*ConfigResponse, error)
	GetStatus(context.Context, *emptypb.Empty) (*StatusResponse, error)
	GetSubscriber(context.Context, *emptypb.Empty) (*SubscriberResponse, error)
	SubscribeKline(*emptypb.Empty, Feed_SubscribeKlineServer) error
	ReadHistoricalKline(*ReadKlineRequest, Feed_ReadHistoricalKlineServer) error
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) GetConfig(context.Context, *emptypb.Empty) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedFeedServer) GetStatus(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedFeedServer) GetSubscriber(context.Context, *emptypb.Empty) (*SubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriber not implemented")
}
func (UnimplementedFeedServer) SubscribeKline(*emptypb.Empty, Feed_SubscribeKlineServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeKline not implemented")
}
func (UnimplementedFeedServer) ReadHistoricalKline(*ReadKlineRequest, Feed_ReadHistoricalKlineServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadHistoricalKline not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_GetSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_GetSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetSubscriber(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_SubscribeKline_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedServer).SubscribeKline(m, &feedSubscribeKlineServer{stream})
}

type Feed_SubscribeKlineServer interface {
	Send(*KlineResponse) error
	grpc.ServerStream
}

type feedSubscribeKlineServer struct {
	grpc.ServerStream
}

func (x *feedSubscribeKlineServer) Send(m *KlineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Feed_ReadHistoricalKline_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadKlineRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedServer).ReadHistoricalKline(m, &feedReadHistoricalKlineServer{stream})
}

type Feed_ReadHistoricalKlineServer interface {
	Send(*KlineResponse) error
	grpc.ServerStream
}

type feedReadHistoricalKlineServer struct {
	grpc.ServerStream
}

func (x *feedReadHistoricalKlineServer) Send(m *KlineResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Feed_GetConfig_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Feed_GetStatus_Handler,
		},
		{
			MethodName: "GetSubscriber",
			Handler:    _Feed_GetSubscriber_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeKline",
			Handler:       _Feed_SubscribeKline_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadHistoricalKline",
			Handler:       _Feed_ReadHistoricalKline_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/feed.proto",
}