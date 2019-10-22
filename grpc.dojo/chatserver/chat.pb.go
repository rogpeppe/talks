// Code generated by protoc-gen-go.
// source: chat.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	Username
	NewUserResponse
	Server2User
	User2Server
	Empty
	User
	UserList
	ServerAnnouncement
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Username holds a desired username.
type Username struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *Username) Reset()                    { *m = Username{} }
func (m *Username) String() string            { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()               {}
func (*Username) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// NewUserResponse holds the response to a NewUser
// RPC call.
type NewUserResponse struct {
	// Token holds a random token that will act as
	// a password for authenticating the user.
	// If you don't save this, you won't be able
	// to act as the user.
	Token string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	// IPAddr holds the IP address (without port)
	// of the originating machine.
	IPAddr string `protobuf:"bytes,2,opt,name=IPAddr" json:"IPAddr,omitempty"`
}

func (m *NewUserResponse) Reset()                    { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string            { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()               {}
func (*NewUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Server2User holds a chat message send from the
// server to a listener.
type Server2User struct {
	// Username holds the name of the user that
	// sent the message.
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	// Text holds the text of the message.
	Text string `protobuf:"bytes,2,opt,name=Text" json:"Text,omitempty"`
}

func (m *Server2User) Reset()                    { *m = Server2User{} }
func (m *Server2User) String() string            { return proto.CompactTextString(m) }
func (*Server2User) ProtoMessage()               {}
func (*Server2User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// User2Server holds a chat message sent from
// a user to the server. It will be broadcast to
// all listeners.
type User2Server struct {
	// Username holds the name of the user that's
	// sending the message.
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	// Token holds the user's token (this acts as a
	// password).
	Token string `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
	// Text holds the text of the message.
	Text string `protobuf:"bytes,3,opt,name=Text" json:"Text,omitempty"`
}

func (m *User2Server) Reset()                    { *m = User2Server{} }
func (m *User2Server) String() string            { return proto.CompactTextString(m) }
func (*User2Server) ProtoMessage()               {}
func (*User2Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Empty is used to signify an empty response.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type User struct {
	// Name holds the user's name.
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// ServerAddr holds the address and port of the
	// server (e.g. 10.3.4.6:5677)
	ServerAddr string `protobuf:"bytes,2,opt,name=ServerAddr" json:"ServerAddr,omitempty"`
	// ServerProtocol holds the .proto RPC description
	// served by the server.
	ServerProtocol string `protobuf:"bytes,3,opt,name=ServerProtocol" json:"ServerProtocol,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// UserList holds a list of all the users.
type UserList struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UserList) Reset()                    { *m = UserList{} }
func (m *UserList) String() string            { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()               {}
func (*UserList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ServerAnnouncement struct {
	// Username holds the name of the user that's
	// running the server.
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	// Token holds the user's token (this acts as a
	// password).
	Token string `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
	// ServerAddr holds the address and port of the
	// server (e.g. 10.3.4.6:5677)
	ServerAddr string `protobuf:"bytes,3,opt,name=ServerAddr" json:"ServerAddr,omitempty"`
	// ServerProtocol holds the .proto RPC description
	// served by the server.
	ServerProtocol string `protobuf:"bytes,4,opt,name=ServerProtocol" json:"ServerProtocol,omitempty"`
}

func (m *ServerAnnouncement) Reset()                    { *m = ServerAnnouncement{} }
func (m *ServerAnnouncement) String() string            { return proto.CompactTextString(m) }
func (*ServerAnnouncement) ProtoMessage()               {}
func (*ServerAnnouncement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Username)(nil), "main.Username")
	proto.RegisterType((*NewUserResponse)(nil), "main.NewUserResponse")
	proto.RegisterType((*Server2User)(nil), "main.Server2User")
	proto.RegisterType((*User2Server)(nil), "main.User2Server")
	proto.RegisterType((*Empty)(nil), "main.Empty")
	proto.RegisterType((*User)(nil), "main.User")
	proto.RegisterType((*UserList)(nil), "main.UserList")
	proto.RegisterType((*ServerAnnouncement)(nil), "main.ServerAnnouncement")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	// NewUser creates a new user account.
	// It returns a token that acts as a password for
	// the user and the IP address of the user client.
	NewUser(ctx context.Context, in *Username, opts ...grpc.CallOption) (*NewUserResponse, error)
	// Say sends a chat message to the server.
	// The message is authenticated by a user token
	// as returned by NewUser.
	Say(ctx context.Context, in *User2Server, opts ...grpc.CallOption) (*Empty, error)
	// Listen listens for any messages sent by anyone.
	Listen(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chat_ListenClient, error)
	// Who returns a list of all the users.
	Who(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error)
	// AnnounceRPCServer announces that a particular
	// user has an RPC server running.
	// The message is authenticated by a user token
	// as returned by NewUser.
	AnnounceRPCServer(ctx context.Context, in *ServerAnnouncement, opts ...grpc.CallOption) (*Empty, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) NewUser(ctx context.Context, in *Username, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := grpc.Invoke(ctx, "/main.Chat/NewUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Say(ctx context.Context, in *User2Server, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/main.Chat/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Listen(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chat_ListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/main.Chat/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ListenClient interface {
	Recv() (*Server2User, error)
	grpc.ClientStream
}

type chatListenClient struct {
	grpc.ClientStream
}

func (x *chatListenClient) Recv() (*Server2User, error) {
	m := new(Server2User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) Who(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := grpc.Invoke(ctx, "/main.Chat/Who", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AnnounceRPCServer(ctx context.Context, in *ServerAnnouncement, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/main.Chat/AnnounceRPCServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatServer interface {
	// NewUser creates a new user account.
	// It returns a token that acts as a password for
	// the user and the IP address of the user client.
	NewUser(context.Context, *Username) (*NewUserResponse, error)
	// Say sends a chat message to the server.
	// The message is authenticated by a user token
	// as returned by NewUser.
	Say(context.Context, *User2Server) (*Empty, error)
	// Listen listens for any messages sent by anyone.
	Listen(*Empty, Chat_ListenServer) error
	// Who returns a list of all the users.
	Who(context.Context, *Empty) (*UserList, error)
	// AnnounceRPCServer announces that a particular
	// user has an RPC server running.
	// The message is authenticated by a user token
	// as returned by NewUser.
	AnnounceRPCServer(context.Context, *ServerAnnouncement) (*Empty, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Chat/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).NewUser(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User2Server)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Chat/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Say(ctx, req.(*User2Server))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).Listen(m, &chatListenServer{stream})
}

type Chat_ListenServer interface {
	Send(*Server2User) error
	grpc.ServerStream
}

type chatListenServer struct {
	grpc.ServerStream
}

func (x *chatListenServer) Send(m *Server2User) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_Who_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Who(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Chat/Who",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Who(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AnnounceRPCServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerAnnouncement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AnnounceRPCServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Chat/AnnounceRPCServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AnnounceRPCServer(ctx, req.(*ServerAnnouncement))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _Chat_NewUser_Handler,
		},
		{
			MethodName: "Say",
			Handler:    _Chat_Say_Handler,
		},
		{
			MethodName: "Who",
			Handler:    _Chat_Who_Handler,
		},
		{
			MethodName: "AnnounceRPCServer",
			Handler:    _Chat_AnnounceRPCServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Chat_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x93, 0x26, 0x4d, 0x9f, 0x67, 0x02, 0x95, 0x0e, 0x2a, 0x21, 0x87, 0x52, 0x16, 0xd1,
	0x1e, 0x4a, 0x91, 0x78, 0x55, 0x44, 0x8b, 0x07, 0x41, 0x4a, 0x69, 0x2b, 0x82, 0xb7, 0xb4, 0x5d,
	0x68, 0xd1, 0xec, 0x96, 0x64, 0x7d, 0xe9, 0x97, 0xf0, 0xdb, 0x7a, 0x97, 0x7d, 0xa9, 0xdd, 0x86,
	0x22, 0xe2, 0x6d, 0x67, 0x76, 0xf6, 0x3f, 0xbf, 0x79, 0x59, 0x80, 0xe9, 0x3c, 0x15, 0xdd, 0x65,
	0xce, 0x05, 0x47, 0x3f, 0x4b, 0x17, 0x8c, 0x34, 0xe1, 0xdf, 0x7d, 0x41, 0x73, 0x96, 0x66, 0x14,
	0x11, 0xfc, 0x7e, 0x9a, 0xd1, 0xc8, 0x6d, 0xb9, 0xed, 0xff, 0x43, 0x75, 0x26, 0x97, 0xb0, 0xd7,
	0xa7, 0x6f, 0x32, 0x64, 0x48, 0x8b, 0x25, 0x67, 0x05, 0xc5, 0x7d, 0xa8, 0x8e, 0xf9, 0x13, 0x65,
	0x26, 0x4e, 0x1b, 0x78, 0x08, 0xc1, 0xed, 0xe0, 0x6a, 0x36, 0xcb, 0xa3, 0x8a, 0x72, 0x1b, 0x8b,
	0x5c, 0x40, 0x38, 0xa2, 0xf9, 0x2b, 0xcd, 0x13, 0x29, 0x82, 0xf1, 0x26, 0x9f, 0x79, 0xbf, 0x95,
	0x7f, 0x4c, 0xdf, 0x85, 0x11, 0x50, 0x67, 0x32, 0x82, 0x50, 0xde, 0x27, 0x5a, 0xe3, 0xc7, 0xe7,
	0xdf, 0x5c, 0x15, 0x9b, 0x6b, 0x2d, 0xea, 0x59, 0xa2, 0x35, 0xa8, 0xde, 0x64, 0x4b, 0xb1, 0x22,
	0x13, 0xf0, 0x15, 0xd5, 0x8e, 0xca, 0xb1, 0x09, 0xa0, 0x93, 0x5a, 0x45, 0x59, 0x1e, 0x3c, 0x86,
	0xba, 0xb6, 0x06, 0xb2, 0x9d, 0x53, 0xfe, 0x6c, 0x52, 0x94, 0xbc, 0xa4, 0xa3, 0x91, 0xef, 0x16,
	0x85, 0xc0, 0x16, 0x54, 0xe5, 0xb9, 0x88, 0xdc, 0x96, 0xd7, 0x0e, 0x13, 0xe8, 0xca, 0x19, 0x74,
	0x55, 0x77, 0xf5, 0x05, 0xf9, 0x70, 0x01, 0x4d, 0x12, 0xc6, 0xf8, 0x0b, 0x9b, 0xd2, 0x8c, 0x32,
	0xf1, 0x87, 0xba, 0xb7, 0xf1, 0xbd, 0x5f, 0xe0, 0xfb, 0xbb, 0xf0, 0x93, 0x4f, 0x17, 0xfc, 0xde,
	0x3c, 0x15, 0x98, 0x40, 0xcd, 0x6c, 0x02, 0xd6, 0x37, 0xdc, 0x92, 0x20, 0x3e, 0xd0, 0x76, 0x69,
	0x51, 0x88, 0x83, 0x27, 0xe0, 0x8d, 0xd2, 0x15, 0x36, 0x36, 0xf1, 0x66, 0x90, 0x71, 0xa8, 0x5d,
	0x7a, 0x0c, 0x0e, 0x76, 0x20, 0x90, 0x0d, 0xa2, 0x0c, 0xed, 0x8b, 0xd8, 0x3c, 0xb4, 0x16, 0x88,
	0x38, 0xa7, 0x2e, 0x1e, 0x81, 0xf7, 0x30, 0xe7, 0xdb, 0xa1, 0x16, 0x93, 0x54, 0x22, 0x0e, 0x9e,
	0x43, 0x63, 0xdd, 0xc3, 0xe1, 0xa0, 0x67, 0x16, 0x28, 0xb2, 0x15, 0xed, 0x16, 0x97, 0x88, 0xae,
	0x83, 0x47, 0xf5, 0x41, 0x26, 0x81, 0xfa, 0x2d, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xfb, 0x72, 0xa9, 0x3b, 0x03, 0x00, 0x00,
}
