// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adventar/v1/adventar.proto

package adventar_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("adventar/v1/adventar.proto", fileDescriptor_c3b1639048342dcd) }

var fileDescriptor_c3b1639048342dcd = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0x87, 0x77, 0xf3, 0x8e, 0x97, 0x14, 0x77, 0x11, 0xd0, 0x8b, 0x0c, 0xfc, 0xd3, 0x0f, 0x90,
	0x32, 0xbd, 0x17, 0x64, 0x1b, 0x2a, 0xa8, 0x17, 0x13, 0xc1, 0x3b, 0xc9, 0xd6, 0x63, 0x19, 0x74,
	0x49, 0x4c, 0xd2, 0xc2, 0xbe, 0x9b, 0x1f, 0x4e, 0x6c, 0x9a, 0x92, 0x23, 0xeb, 0x7a, 0xd7, 0x9e,
	0xe7, 0x77, 0x1e, 0xce, 0xc9, 0x21, 0x4c, 0xe4, 0x35, 0x48, 0x27, 0x4c, 0x56, 0xcf, 0xb2, 0xf0,
	0xcd, 0xb5, 0x51, 0x4e, 0xd1, 0xa4, 0xfb, 0xaf, 0x67, 0x6c, 0x5a, 0x28, 0x55, 0x94, 0x90, 0x35,
	0x68, 0x5d, 0x7d, 0x66, 0xb0, 0xd3, 0x6e, 0xef, 0x93, 0x6c, 0x1a, 0x5b, 0x0c, 0x58, 0x55, 0x99,
	0x0d, 0xd8, 0x16, 0x9e, 0x23, 0xa8, 0x37, 0x1f, 0x3b, 0xb0, 0x56, 0x14, 0x81, 0x5f, 0x7f, 0xff,
	0x23, 0xff, 0xef, 0xda, 0x08, 0x7d, 0x27, 0x27, 0x4f, 0x5b, 0xeb, 0xe6, 0xa2, 0x04, 0x99, 0x0b,
	0x63, 0xe9, 0x15, 0x8f, 0xa6, 0xe0, 0x88, 0xad, 0xe0, 0xab, 0x02, 0xeb, 0x58, 0x7a, 0x2c, 0x62,
	0xb5, 0x92, 0x16, 0xd2, 0x11, 0x5d, 0x91, 0xe4, 0x1e, 0x3a, 0x42, 0x2f, 0x50, 0x53, 0x44, 0x82,
	0xf5, 0xb2, 0x3f, 0xd0, 0x39, 0x9f, 0xc9, 0x64, 0x6e, 0x40, 0x38, 0xe8, 0xb4, 0x78, 0x16, 0x0c,
	0x83, 0xf9, 0x14, 0x67, 0x5a, 0xea, 0x75, 0x6f, 0x3a, 0xef, 0xd7, 0x61, 0x38, 0xa8, 0x7b, 0x21,
	0x93, 0x05, 0x94, 0xd0, 0xab, 0xc3, 0x30, 0xe8, 0xce, 0xb8, 0xbf, 0x34, 0x0f, 0x97, 0xe6, 0xcb,
	0xdf, 0x4b, 0xa7, 0x23, 0xba, 0x20, 0x89, 0x5f, 0x68, 0x29, 0x9d, 0xd9, 0xff, 0x79, 0xc1, 0x88,
	0x04, 0x13, 0x45, 0x81, 0x06, 0x79, 0x8b, 0xdf, 0xe3, 0x90, 0x25, 0x22, 0xc7, 0x2d, 0x0f, 0x24,
	0xf1, 0xe3, 0x1f, 0xb2, 0x44, 0x64, 0x78, 0xab, 0x5b, 0x32, 0x7e, 0xdd, 0x16, 0xf2, 0x51, 0x52,
	0x86, 0x24, 0xbe, 0x38, 0xd8, 0xbf, 0x1e, 0x37, 0x95, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x34, 0x7a, 0xcb, 0x96, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdventarClient is the client API for Adventar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdventarClient interface {
	ListCalendars(ctx context.Context, in *ListCalendarsRequest, opts ...grpc.CallOption) (*ListCalendarsResponse, error)
	GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error)
	CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type adventarClient struct {
	cc *grpc.ClientConn
}

func NewAdventarClient(cc *grpc.ClientConn) AdventarClient {
	return &adventarClient{cc}
}

func (c *adventarClient) ListCalendars(ctx context.Context, in *ListCalendarsRequest, opts ...grpc.CallOption) (*ListCalendarsResponse, error) {
	out := new(ListCalendarsResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/ListCalendars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error) {
	out := new(GetCalendarResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/GetCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/CreateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/UpdateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/DeleteCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventarServer is the server API for Adventar service.
type AdventarServer interface {
	ListCalendars(context.Context, *ListCalendarsRequest) (*ListCalendarsResponse, error)
	GetCalendar(context.Context, *GetCalendarRequest) (*GetCalendarResponse, error)
	CreateCalendar(context.Context, *CreateCalendarRequest) (*Calendar, error)
	UpdateCalendar(context.Context, *UpdateCalendarRequest) (*Calendar, error)
	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*empty.Empty, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*Entry, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*Entry, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*empty.Empty, error)
	SignIn(context.Context, *SignInRequest) (*empty.Empty, error)
}

func RegisterAdventarServer(s *grpc.Server, srv AdventarServer) {
	s.RegisterService(&_Adventar_serviceDesc, srv)
}

func _Adventar_ListCalendars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCalendarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).ListCalendars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/ListCalendars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).ListCalendars(ctx, req.(*ListCalendarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_GetCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).GetCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/GetCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).GetCalendar(ctx, req.(*GetCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_CreateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).CreateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/CreateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).CreateCalendar(ctx, req.(*CreateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_UpdateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).UpdateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/UpdateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).UpdateCalendar(ctx, req.(*UpdateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_DeleteCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).DeleteCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/DeleteCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).DeleteCalendar(ctx, req.(*DeleteCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adventar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adventar.v1.Adventar",
	HandlerType: (*AdventarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCalendars",
			Handler:    _Adventar_ListCalendars_Handler,
		},
		{
			MethodName: "GetCalendar",
			Handler:    _Adventar_GetCalendar_Handler,
		},
		{
			MethodName: "CreateCalendar",
			Handler:    _Adventar_CreateCalendar_Handler,
		},
		{
			MethodName: "UpdateCalendar",
			Handler:    _Adventar_UpdateCalendar_Handler,
		},
		{
			MethodName: "DeleteCalendar",
			Handler:    _Adventar_DeleteCalendar_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _Adventar_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Adventar_UpdateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Adventar_DeleteEntry_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Adventar_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adventar/v1/adventar.proto",
}
