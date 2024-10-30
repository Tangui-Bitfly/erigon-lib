// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: txpool/mining.proto

package txpoolproto

import (
	context "context"
	typesproto "github.com/Tangui-Bitfly/erigon-lib/gointerfaces/typesproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Mining_Version_FullMethodName        = "/txpool.Mining/Version"
	Mining_OnPendingBlock_FullMethodName = "/txpool.Mining/OnPendingBlock"
	Mining_OnMinedBlock_FullMethodName   = "/txpool.Mining/OnMinedBlock"
	Mining_OnPendingLogs_FullMethodName  = "/txpool.Mining/OnPendingLogs"
	Mining_GetWork_FullMethodName        = "/txpool.Mining/GetWork"
	Mining_SubmitWork_FullMethodName     = "/txpool.Mining/SubmitWork"
	Mining_SubmitHashRate_FullMethodName = "/txpool.Mining/SubmitHashRate"
	Mining_HashRate_FullMethodName       = "/txpool.Mining/HashRate"
	Mining_Mining_FullMethodName         = "/txpool.Mining/Mining"
)

// MiningClient is the client API for Mining service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiningClient interface {
	// Version returns the service version number
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*typesproto.VersionReply, error)
	// subscribe to pending blocks event
	OnPendingBlock(ctx context.Context, in *OnPendingBlockRequest, opts ...grpc.CallOption) (Mining_OnPendingBlockClient, error)
	// subscribe to mined blocks event
	OnMinedBlock(ctx context.Context, in *OnMinedBlockRequest, opts ...grpc.CallOption) (Mining_OnMinedBlockClient, error)
	// subscribe to pending blocks event
	OnPendingLogs(ctx context.Context, in *OnPendingLogsRequest, opts ...grpc.CallOption) (Mining_OnPendingLogsClient, error)
	// GetWork returns a work package for external miner.
	//
	// The work package consists of 3 strings:
	//
	//	result[0] - 32 bytes hex encoded current block header pow-hash
	//	result[1] - 32 bytes hex encoded seed hash used for DAG
	//	result[2] - 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	//	result[3] - hex encoded block number
	GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkReply, error)
	// SubmitWork can be used by external miner to submit their POW solution.
	// It returns an indication if the work was accepted.
	// Note either an invalid solution, a stale work a non-existent work will return false.
	SubmitWork(ctx context.Context, in *SubmitWorkRequest, opts ...grpc.CallOption) (*SubmitWorkReply, error)
	// SubmitHashRate can be used for remote miners to submit their hash rate.
	// This enables the node to report the combined hash rate of all miners
	// which submit work through this node.
	//
	// It accepts the miner hash rate and an identifier which must be unique
	// between nodes.
	SubmitHashRate(ctx context.Context, in *SubmitHashRateRequest, opts ...grpc.CallOption) (*SubmitHashRateReply, error)
	// HashRate returns the current hashrate for local CPU miner and remote miner.
	HashRate(ctx context.Context, in *HashRateRequest, opts ...grpc.CallOption) (*HashRateReply, error)
	// Mining returns an indication if this node is currently mining and its mining configuration
	Mining(ctx context.Context, in *MiningRequest, opts ...grpc.CallOption) (*MiningReply, error)
}

type miningClient struct {
	cc grpc.ClientConnInterface
}

func NewMiningClient(cc grpc.ClientConnInterface) MiningClient {
	return &miningClient{cc}
}

func (c *miningClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*typesproto.VersionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(typesproto.VersionReply)
	err := c.cc.Invoke(ctx, Mining_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningClient) OnPendingBlock(ctx context.Context, in *OnPendingBlockRequest, opts ...grpc.CallOption) (Mining_OnPendingBlockClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Mining_ServiceDesc.Streams[0], Mining_OnPendingBlock_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &miningOnPendingBlockClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mining_OnPendingBlockClient interface {
	Recv() (*OnPendingBlockReply, error)
	grpc.ClientStream
}

type miningOnPendingBlockClient struct {
	grpc.ClientStream
}

func (x *miningOnPendingBlockClient) Recv() (*OnPendingBlockReply, error) {
	m := new(OnPendingBlockReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *miningClient) OnMinedBlock(ctx context.Context, in *OnMinedBlockRequest, opts ...grpc.CallOption) (Mining_OnMinedBlockClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Mining_ServiceDesc.Streams[1], Mining_OnMinedBlock_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &miningOnMinedBlockClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mining_OnMinedBlockClient interface {
	Recv() (*OnMinedBlockReply, error)
	grpc.ClientStream
}

type miningOnMinedBlockClient struct {
	grpc.ClientStream
}

func (x *miningOnMinedBlockClient) Recv() (*OnMinedBlockReply, error) {
	m := new(OnMinedBlockReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *miningClient) OnPendingLogs(ctx context.Context, in *OnPendingLogsRequest, opts ...grpc.CallOption) (Mining_OnPendingLogsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Mining_ServiceDesc.Streams[2], Mining_OnPendingLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &miningOnPendingLogsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mining_OnPendingLogsClient interface {
	Recv() (*OnPendingLogsReply, error)
	grpc.ClientStream
}

type miningOnPendingLogsClient struct {
	grpc.ClientStream
}

func (x *miningOnPendingLogsClient) Recv() (*OnPendingLogsReply, error) {
	m := new(OnPendingLogsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *miningClient) GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorkReply)
	err := c.cc.Invoke(ctx, Mining_GetWork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningClient) SubmitWork(ctx context.Context, in *SubmitWorkRequest, opts ...grpc.CallOption) (*SubmitWorkReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitWorkReply)
	err := c.cc.Invoke(ctx, Mining_SubmitWork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningClient) SubmitHashRate(ctx context.Context, in *SubmitHashRateRequest, opts ...grpc.CallOption) (*SubmitHashRateReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitHashRateReply)
	err := c.cc.Invoke(ctx, Mining_SubmitHashRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningClient) HashRate(ctx context.Context, in *HashRateRequest, opts ...grpc.CallOption) (*HashRateReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HashRateReply)
	err := c.cc.Invoke(ctx, Mining_HashRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningClient) Mining(ctx context.Context, in *MiningRequest, opts ...grpc.CallOption) (*MiningReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MiningReply)
	err := c.cc.Invoke(ctx, Mining_Mining_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiningServer is the server API for Mining service.
// All implementations must embed UnimplementedMiningServer
// for forward compatibility
type MiningServer interface {
	// Version returns the service version number
	Version(context.Context, *emptypb.Empty) (*typesproto.VersionReply, error)
	// subscribe to pending blocks event
	OnPendingBlock(*OnPendingBlockRequest, Mining_OnPendingBlockServer) error
	// subscribe to mined blocks event
	OnMinedBlock(*OnMinedBlockRequest, Mining_OnMinedBlockServer) error
	// subscribe to pending blocks event
	OnPendingLogs(*OnPendingLogsRequest, Mining_OnPendingLogsServer) error
	// GetWork returns a work package for external miner.
	//
	// The work package consists of 3 strings:
	//
	//	result[0] - 32 bytes hex encoded current block header pow-hash
	//	result[1] - 32 bytes hex encoded seed hash used for DAG
	//	result[2] - 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	//	result[3] - hex encoded block number
	GetWork(context.Context, *GetWorkRequest) (*GetWorkReply, error)
	// SubmitWork can be used by external miner to submit their POW solution.
	// It returns an indication if the work was accepted.
	// Note either an invalid solution, a stale work a non-existent work will return false.
	SubmitWork(context.Context, *SubmitWorkRequest) (*SubmitWorkReply, error)
	// SubmitHashRate can be used for remote miners to submit their hash rate.
	// This enables the node to report the combined hash rate of all miners
	// which submit work through this node.
	//
	// It accepts the miner hash rate and an identifier which must be unique
	// between nodes.
	SubmitHashRate(context.Context, *SubmitHashRateRequest) (*SubmitHashRateReply, error)
	// HashRate returns the current hashrate for local CPU miner and remote miner.
	HashRate(context.Context, *HashRateRequest) (*HashRateReply, error)
	// Mining returns an indication if this node is currently mining and its mining configuration
	Mining(context.Context, *MiningRequest) (*MiningReply, error)
	mustEmbedUnimplementedMiningServer()
}

// UnimplementedMiningServer must be embedded to have forward compatible implementations.
type UnimplementedMiningServer struct {
}

func (UnimplementedMiningServer) Version(context.Context, *emptypb.Empty) (*typesproto.VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedMiningServer) OnPendingBlock(*OnPendingBlockRequest, Mining_OnPendingBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method OnPendingBlock not implemented")
}
func (UnimplementedMiningServer) OnMinedBlock(*OnMinedBlockRequest, Mining_OnMinedBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method OnMinedBlock not implemented")
}
func (UnimplementedMiningServer) OnPendingLogs(*OnPendingLogsRequest, Mining_OnPendingLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method OnPendingLogs not implemented")
}
func (UnimplementedMiningServer) GetWork(context.Context, *GetWorkRequest) (*GetWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWork not implemented")
}
func (UnimplementedMiningServer) SubmitWork(context.Context, *SubmitWorkRequest) (*SubmitWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitWork not implemented")
}
func (UnimplementedMiningServer) SubmitHashRate(context.Context, *SubmitHashRateRequest) (*SubmitHashRateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitHashRate not implemented")
}
func (UnimplementedMiningServer) HashRate(context.Context, *HashRateRequest) (*HashRateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashRate not implemented")
}
func (UnimplementedMiningServer) Mining(context.Context, *MiningRequest) (*MiningReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mining not implemented")
}
func (UnimplementedMiningServer) mustEmbedUnimplementedMiningServer() {}

// UnsafeMiningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiningServer will
// result in compilation errors.
type UnsafeMiningServer interface {
	mustEmbedUnimplementedMiningServer()
}

func RegisterMiningServer(s grpc.ServiceRegistrar, srv MiningServer) {
	s.RegisterService(&Mining_ServiceDesc, srv)
}

func _Mining_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mining_OnPendingBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OnPendingBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MiningServer).OnPendingBlock(m, &miningOnPendingBlockServer{ServerStream: stream})
}

type Mining_OnPendingBlockServer interface {
	Send(*OnPendingBlockReply) error
	grpc.ServerStream
}

type miningOnPendingBlockServer struct {
	grpc.ServerStream
}

func (x *miningOnPendingBlockServer) Send(m *OnPendingBlockReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Mining_OnMinedBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OnMinedBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MiningServer).OnMinedBlock(m, &miningOnMinedBlockServer{ServerStream: stream})
}

type Mining_OnMinedBlockServer interface {
	Send(*OnMinedBlockReply) error
	grpc.ServerStream
}

type miningOnMinedBlockServer struct {
	grpc.ServerStream
}

func (x *miningOnMinedBlockServer) Send(m *OnMinedBlockReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Mining_OnPendingLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OnPendingLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MiningServer).OnPendingLogs(m, &miningOnPendingLogsServer{ServerStream: stream})
}

type Mining_OnPendingLogsServer interface {
	Send(*OnPendingLogsReply) error
	grpc.ServerStream
}

type miningOnPendingLogsServer struct {
	grpc.ServerStream
}

func (x *miningOnPendingLogsServer) Send(m *OnPendingLogsReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Mining_GetWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).GetWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_GetWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).GetWork(ctx, req.(*GetWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mining_SubmitWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).SubmitWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_SubmitWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).SubmitWork(ctx, req.(*SubmitWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mining_SubmitHashRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitHashRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).SubmitHashRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_SubmitHashRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).SubmitHashRate(ctx, req.(*SubmitHashRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mining_HashRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).HashRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_HashRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).HashRate(ctx, req.(*HashRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mining_Mining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningServer).Mining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mining_Mining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningServer).Mining(ctx, req.(*MiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mining_ServiceDesc is the grpc.ServiceDesc for Mining service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mining_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "txpool.Mining",
	HandlerType: (*MiningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Mining_Version_Handler,
		},
		{
			MethodName: "GetWork",
			Handler:    _Mining_GetWork_Handler,
		},
		{
			MethodName: "SubmitWork",
			Handler:    _Mining_SubmitWork_Handler,
		},
		{
			MethodName: "SubmitHashRate",
			Handler:    _Mining_SubmitHashRate_Handler,
		},
		{
			MethodName: "HashRate",
			Handler:    _Mining_HashRate_Handler,
		},
		{
			MethodName: "Mining",
			Handler:    _Mining_Mining_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnPendingBlock",
			Handler:       _Mining_OnPendingBlock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnMinedBlock",
			Handler:       _Mining_OnMinedBlock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnPendingLogs",
			Handler:       _Mining_OnPendingLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "txpool/mining.proto",
}
