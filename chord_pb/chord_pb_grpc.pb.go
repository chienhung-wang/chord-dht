// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: chord_pb/chord_pb.proto

package chord_pb

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

// ChordNodeClient is the client API for ChordNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChordNodeClient interface {
	Join(ctx context.Context, in *NodeAddr, opts ...grpc.CallOption) (*NodeAddr, error)
	FindSuccessor(ctx context.Context, in *NodeId, opts ...grpc.CallOption) (*FindFindSuccessorResp, error)
	GetPredecessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeAddr, error)
	GetFirstSuccessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeAddr, error)
	Notify(ctx context.Context, in *NodeAddr, opts ...grpc.CallOption) (*Empty, error)
	GetFingers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFingersResp, error)
	HelpFind(ctx context.Context, in *NodeId, opts ...grpc.CallOption) (*NodeAddr, error)
	CheckAlive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	KeyTransfer(ctx context.Context, in *KeyValueMap, opts ...grpc.CallOption) (*KeyTransferResp, error)
	MapGet(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyVal, error)
	MapPut(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*KeyVal, error)
	MapDelete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error)
}

type chordNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewChordNodeClient(cc grpc.ClientConnInterface) ChordNodeClient {
	return &chordNodeClient{cc}
}

func (c *chordNodeClient) Join(ctx context.Context, in *NodeAddr, opts ...grpc.CallOption) (*NodeAddr, error) {
	out := new(NodeAddr)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) FindSuccessor(ctx context.Context, in *NodeId, opts ...grpc.CallOption) (*FindFindSuccessorResp, error) {
	out := new(FindFindSuccessorResp)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) GetPredecessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeAddr, error) {
	out := new(NodeAddr)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/GetPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) GetFirstSuccessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeAddr, error) {
	out := new(NodeAddr)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/GetFirstSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) Notify(ctx context.Context, in *NodeAddr, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) GetFingers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFingersResp, error) {
	out := new(GetFingersResp)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/GetFingers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) HelpFind(ctx context.Context, in *NodeId, opts ...grpc.CallOption) (*NodeAddr, error) {
	out := new(NodeAddr)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/HelpFind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) CheckAlive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/CheckAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) KeyTransfer(ctx context.Context, in *KeyValueMap, opts ...grpc.CallOption) (*KeyTransferResp, error) {
	out := new(KeyTransferResp)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/KeyTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) MapGet(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyVal, error) {
	out := new(KeyVal)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/MapGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) MapPut(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*KeyVal, error) {
	out := new(KeyVal)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/MapPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordNodeClient) MapDelete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/chord_pb.ChordNode/MapDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordNodeServer is the server API for ChordNode service.
// All implementations must embed UnimplementedChordNodeServer
// for forward compatibility
type ChordNodeServer interface {
	Join(context.Context, *NodeAddr) (*NodeAddr, error)
	FindSuccessor(context.Context, *NodeId) (*FindFindSuccessorResp, error)
	GetPredecessor(context.Context, *Empty) (*NodeAddr, error)
	GetFirstSuccessor(context.Context, *Empty) (*NodeAddr, error)
	Notify(context.Context, *NodeAddr) (*Empty, error)
	GetFingers(context.Context, *Empty) (*GetFingersResp, error)
	HelpFind(context.Context, *NodeId) (*NodeAddr, error)
	CheckAlive(context.Context, *Empty) (*Empty, error)
	KeyTransfer(context.Context, *KeyValueMap) (*KeyTransferResp, error)
	MapGet(context.Context, *Key) (*KeyVal, error)
	MapPut(context.Context, *KeyVal) (*KeyVal, error)
	MapDelete(context.Context, *Key) (*Key, error)
	mustEmbedUnimplementedChordNodeServer()
}

// UnimplementedChordNodeServer must be embedded to have forward compatible implementations.
type UnimplementedChordNodeServer struct {
}

func (UnimplementedChordNodeServer) Join(context.Context, *NodeAddr) (*NodeAddr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedChordNodeServer) FindSuccessor(context.Context, *NodeId) (*FindFindSuccessorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (UnimplementedChordNodeServer) GetPredecessor(context.Context, *Empty) (*NodeAddr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}
func (UnimplementedChordNodeServer) GetFirstSuccessor(context.Context, *Empty) (*NodeAddr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirstSuccessor not implemented")
}
func (UnimplementedChordNodeServer) Notify(context.Context, *NodeAddr) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedChordNodeServer) GetFingers(context.Context, *Empty) (*GetFingersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFingers not implemented")
}
func (UnimplementedChordNodeServer) HelpFind(context.Context, *NodeId) (*NodeAddr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelpFind not implemented")
}
func (UnimplementedChordNodeServer) CheckAlive(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAlive not implemented")
}
func (UnimplementedChordNodeServer) KeyTransfer(context.Context, *KeyValueMap) (*KeyTransferResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyTransfer not implemented")
}
func (UnimplementedChordNodeServer) MapGet(context.Context, *Key) (*KeyVal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapGet not implemented")
}
func (UnimplementedChordNodeServer) MapPut(context.Context, *KeyVal) (*KeyVal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapPut not implemented")
}
func (UnimplementedChordNodeServer) MapDelete(context.Context, *Key) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapDelete not implemented")
}
func (UnimplementedChordNodeServer) mustEmbedUnimplementedChordNodeServer() {}

// UnsafeChordNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChordNodeServer will
// result in compilation errors.
type UnsafeChordNodeServer interface {
	mustEmbedUnimplementedChordNodeServer()
}

func RegisterChordNodeServer(s grpc.ServiceRegistrar, srv ChordNodeServer) {
	s.RegisterService(&ChordNode_ServiceDesc, srv)
}

func _ChordNode_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).Join(ctx, req.(*NodeAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).FindSuccessor(ctx, req.(*NodeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/GetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).GetPredecessor(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_GetFirstSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).GetFirstSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/GetFirstSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).GetFirstSuccessor(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).Notify(ctx, req.(*NodeAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_GetFingers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).GetFingers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/GetFingers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).GetFingers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_HelpFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).HelpFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/HelpFind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).HelpFind(ctx, req.(*NodeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_CheckAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).CheckAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/CheckAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).CheckAlive(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_KeyTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).KeyTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/KeyTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).KeyTransfer(ctx, req.(*KeyValueMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_MapGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).MapGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/MapGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).MapGet(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_MapPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyVal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).MapPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/MapPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).MapPut(ctx, req.(*KeyVal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordNode_MapDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordNodeServer).MapDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord_pb.ChordNode/MapDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordNodeServer).MapDelete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// ChordNode_ServiceDesc is the grpc.ServiceDesc for ChordNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChordNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chord_pb.ChordNode",
	HandlerType: (*ChordNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _ChordNode_Join_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _ChordNode_FindSuccessor_Handler,
		},
		{
			MethodName: "GetPredecessor",
			Handler:    _ChordNode_GetPredecessor_Handler,
		},
		{
			MethodName: "GetFirstSuccessor",
			Handler:    _ChordNode_GetFirstSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _ChordNode_Notify_Handler,
		},
		{
			MethodName: "GetFingers",
			Handler:    _ChordNode_GetFingers_Handler,
		},
		{
			MethodName: "HelpFind",
			Handler:    _ChordNode_HelpFind_Handler,
		},
		{
			MethodName: "CheckAlive",
			Handler:    _ChordNode_CheckAlive_Handler,
		},
		{
			MethodName: "KeyTransfer",
			Handler:    _ChordNode_KeyTransfer_Handler,
		},
		{
			MethodName: "MapGet",
			Handler:    _ChordNode_MapGet_Handler,
		},
		{
			MethodName: "MapPut",
			Handler:    _ChordNode_MapPut_Handler,
		},
		{
			MethodName: "MapDelete",
			Handler:    _ChordNode_MapDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord_pb/chord_pb.proto",
}
