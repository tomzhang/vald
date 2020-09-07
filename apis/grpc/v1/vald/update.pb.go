//
// Copyright (C) 2019-2020 Vdaas.org Vald team ( kpango, rinx, kmrmt )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package vald

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	payload "github.com/vdaas/vald/apis/grpc/v1/payload"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("apis/proto/v1/vald/update.proto", fileDescriptor_a564bbf4b2600403) }

var fileDescriptor_a564bbf4b2600403 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc8, 0x2c,
	0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x33, 0xd4, 0x2f, 0x4b, 0xcc, 0x49, 0xd1, 0x2f,
	0x2d, 0x48, 0x49, 0x2c, 0x49, 0xd5, 0x03, 0x0b, 0x0a, 0xb1, 0x80, 0x84, 0xa4, 0x2c, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x52, 0x12, 0x13, 0x8b, 0x21, 0x2a,
	0x51, 0x35, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0xc0, 0x68, 0x88, 0x01, 0x52, 0x32, 0xe9,
	0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0x20, 0xb5, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99,
	0xf9, 0x79, 0xc5, 0x10, 0x59, 0xa3, 0x37, 0x8c, 0x5c, 0x6c, 0xa1, 0x60, 0xfb, 0x84, 0xfc, 0xe1,
	0x2c, 0x71, 0x3d, 0x98, 0x11, 0x10, 0x01, 0xbd, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29,
	0x09, 0xb8, 0x84, 0x7f, 0x52, 0x56, 0x6a, 0x72, 0x89, 0x9e, 0x4f, 0x7e, 0x32, 0xd8, 0x38, 0x25,
	0xa1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x28, 0xb1, 0x43, 0x5d, 0x6f, 0xc5, 0xa8, 0x25, 0xe4,
	0xce, 0xc5, 0x13, 0x5c, 0x52, 0x94, 0x9a, 0x98, 0x4b, 0xbe, 0xb1, 0x0c, 0x1a, 0x8c, 0x06, 0x8c,
	0x42, 0x1e, 0x5c, 0xdc, 0xbe, 0xa5, 0x39, 0x25, 0x99, 0x50, 0x73, 0x64, 0xd0, 0xcd, 0x01, 0x4b,
	0xc2, 0x0c, 0x93, 0xc4, 0x65, 0x58, 0xb1, 0x12, 0x83, 0x53, 0xf4, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x25, 0x95, 0x5f, 0x94, 0xae, 0x07, 0x0e, 0x4e,
	0x3d, 0x50, 0x70, 0xea, 0x25, 0x16, 0x64, 0xea, 0x95, 0x19, 0x82, 0xd9, 0x4e, 0xd0, 0xb0, 0x08,
	0x60, 0x8c, 0x52, 0xc7, 0x13, 0xf6, 0xe9, 0x45, 0x05, 0xc9, 0xb0, 0x78, 0x4b, 0x62, 0x03, 0x07,
	0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1e, 0x2b, 0xff, 0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdateClient is the client API for Update service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateClient interface {
	Update(ctx context.Context, in *payload.Update_Request, opts ...grpc.CallOption) (*payload.Object_Location, error)
	StreamUpdate(ctx context.Context, opts ...grpc.CallOption) (Update_StreamUpdateClient, error)
	MultiUpdate(ctx context.Context, in *payload.Update_MultiRequest, opts ...grpc.CallOption) (*payload.Object_Locations, error)
}

type updateClient struct {
	cc *grpc.ClientConn
}

func NewUpdateClient(cc *grpc.ClientConn) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) Update(ctx context.Context, in *payload.Update_Request, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.Update/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) StreamUpdate(ctx context.Context, opts ...grpc.CallOption) (Update_StreamUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Update_serviceDesc.Streams[0], "/vald.Update/StreamUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamUpdateClient{stream}
	return x, nil
}

type Update_StreamUpdateClient interface {
	Send(*payload.Update_Request) error
	Recv() (*payload.Object_Location, error)
	grpc.ClientStream
}

type updateStreamUpdateClient struct {
	grpc.ClientStream
}

func (x *updateStreamUpdateClient) Send(m *payload.Update_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *updateStreamUpdateClient) Recv() (*payload.Object_Location, error) {
	m := new(payload.Object_Location)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateClient) MultiUpdate(ctx context.Context, in *payload.Update_MultiRequest, opts ...grpc.CallOption) (*payload.Object_Locations, error) {
	out := new(payload.Object_Locations)
	err := c.cc.Invoke(ctx, "/vald.Update/MultiUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServer is the server API for Update service.
type UpdateServer interface {
	Update(context.Context, *payload.Update_Request) (*payload.Object_Location, error)
	StreamUpdate(Update_StreamUpdateServer) error
	MultiUpdate(context.Context, *payload.Update_MultiRequest) (*payload.Object_Locations, error)
}

// UnimplementedUpdateServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateServer struct {
}

func (*UnimplementedUpdateServer) Update(ctx context.Context, req *payload.Update_Request) (*payload.Object_Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUpdateServer) StreamUpdate(srv Update_StreamUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUpdate not implemented")
}
func (*UnimplementedUpdateServer) MultiUpdate(ctx context.Context, req *payload.Update_MultiRequest) (*payload.Object_Locations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiUpdate not implemented")
}

func RegisterUpdateServer(s *grpc.Server, srv UpdateServer) {
	s.RegisterService(&_Update_serviceDesc, srv)
}

func _Update_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Update_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.Update/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Update(ctx, req.(*payload.Update_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_StreamUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpdateServer).StreamUpdate(&updateStreamUpdateServer{stream})
}

type Update_StreamUpdateServer interface {
	Send(*payload.Object_Location) error
	Recv() (*payload.Update_Request, error)
	grpc.ServerStream
}

type updateStreamUpdateServer struct {
	grpc.ServerStream
}

func (x *updateStreamUpdateServer) Send(m *payload.Object_Location) error {
	return x.ServerStream.SendMsg(m)
}

func (x *updateStreamUpdateServer) Recv() (*payload.Update_Request, error) {
	m := new(payload.Update_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Update_MultiUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Update_MultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).MultiUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.Update/MultiUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).MultiUpdate(ctx, req.(*payload.Update_MultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Update_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vald.Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Update_Update_Handler,
		},
		{
			MethodName: "MultiUpdate",
			Handler:    _Update_MultiUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUpdate",
			Handler:       _Update_StreamUpdate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apis/proto/v1/vald/update.proto",
}