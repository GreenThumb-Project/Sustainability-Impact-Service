// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: Sustainability-impact-service.proto

package sustainability

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

// SustainabilityimpactServiceClient is the client API for SustainabilityimpactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SustainabilityimpactServiceClient interface {
	LogImpact(ctx context.Context, in *LogImpactRequest, opts ...grpc.CallOption) (*LogImpactResponse, error)
	GetUserImpact(ctx context.Context, in *GetUserImpactRequest, opts ...grpc.CallOption) (*GetUserImpactResponse, error)
	GetCommunityImpact(ctx context.Context, in *GetCommunityImpactRequest, opts ...grpc.CallOption) (*GetCommunityImpactResponse, error)
	GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error)
	JoinChallenge(ctx context.Context, in *JoinChallengeRequest, opts ...grpc.CallOption) (*JoinChallengeResponse, error)
	UpdateChallengeProgress(ctx context.Context, in *UpdateChallengeProgressRequest, opts ...grpc.CallOption) (*UpdateChallengeProgressResponse, error)
	GetUserChallenges(ctx context.Context, in *GetUserChallengesRequest, opts ...grpc.CallOption) (*GetUserChallengesResponse, error)
	GetUsersLeaderboard(ctx context.Context, in *GetUsersLeaderboardRequest, opts ...grpc.CallOption) (*GetUsersLeaderboardResponse, error)
	GetCommunitiesLeaderboard(ctx context.Context, in *GetCommunitiesLeaderboardRequest, opts ...grpc.CallOption) (*GetCommunitiesLeaderboardResponse, error)
}

type sustainabilityimpactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSustainabilityimpactServiceClient(cc grpc.ClientConnInterface) SustainabilityimpactServiceClient {
	return &sustainabilityimpactServiceClient{cc}
}

func (c *sustainabilityimpactServiceClient) LogImpact(ctx context.Context, in *LogImpactRequest, opts ...grpc.CallOption) (*LogImpactResponse, error) {
	out := new(LogImpactResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/LogImpact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetUserImpact(ctx context.Context, in *GetUserImpactRequest, opts ...grpc.CallOption) (*GetUserImpactResponse, error) {
	out := new(GetUserImpactResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetUserImpact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetCommunityImpact(ctx context.Context, in *GetCommunityImpactRequest, opts ...grpc.CallOption) (*GetCommunityImpactResponse, error) {
	out := new(GetCommunityImpactResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetCommunityImpact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error) {
	out := new(GetChallengesResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetChallenges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) JoinChallenge(ctx context.Context, in *JoinChallengeRequest, opts ...grpc.CallOption) (*JoinChallengeResponse, error) {
	out := new(JoinChallengeResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/JoinChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) UpdateChallengeProgress(ctx context.Context, in *UpdateChallengeProgressRequest, opts ...grpc.CallOption) (*UpdateChallengeProgressResponse, error) {
	out := new(UpdateChallengeProgressResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/UpdateChallengeProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetUserChallenges(ctx context.Context, in *GetUserChallengesRequest, opts ...grpc.CallOption) (*GetUserChallengesResponse, error) {
	out := new(GetUserChallengesResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetUserChallenges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetUsersLeaderboard(ctx context.Context, in *GetUsersLeaderboardRequest, opts ...grpc.CallOption) (*GetUsersLeaderboardResponse, error) {
	out := new(GetUsersLeaderboardResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetUsersLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityimpactServiceClient) GetCommunitiesLeaderboard(ctx context.Context, in *GetCommunitiesLeaderboardRequest, opts ...grpc.CallOption) (*GetCommunitiesLeaderboardResponse, error) {
	out := new(GetCommunitiesLeaderboardResponse)
	err := c.cc.Invoke(ctx, "/SustainabilityImpactService.SustainabilityimpactService/GetCommunitiesLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SustainabilityimpactServiceServer is the server API for SustainabilityimpactService service.
// All implementations must embed UnimplementedSustainabilityimpactServiceServer
// for forward compatibility
type SustainabilityimpactServiceServer interface {
	LogImpact(context.Context, *LogImpactRequest) (*LogImpactResponse, error)
	GetUserImpact(context.Context, *GetUserImpactRequest) (*GetUserImpactResponse, error)
	GetCommunityImpact(context.Context, *GetCommunityImpactRequest) (*GetCommunityImpactResponse, error)
	GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error)
	JoinChallenge(context.Context, *JoinChallengeRequest) (*JoinChallengeResponse, error)
	UpdateChallengeProgress(context.Context, *UpdateChallengeProgressRequest) (*UpdateChallengeProgressResponse, error)
	GetUserChallenges(context.Context, *GetUserChallengesRequest) (*GetUserChallengesResponse, error)
	GetUsersLeaderboard(context.Context, *GetUsersLeaderboardRequest) (*GetUsersLeaderboardResponse, error)
	GetCommunitiesLeaderboard(context.Context, *GetCommunitiesLeaderboardRequest) (*GetCommunitiesLeaderboardResponse, error)
	mustEmbedUnimplementedSustainabilityimpactServiceServer()
}

// UnimplementedSustainabilityimpactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSustainabilityimpactServiceServer struct {
}

func (UnimplementedSustainabilityimpactServiceServer) LogImpact(context.Context, *LogImpactRequest) (*LogImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogImpact not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetUserImpact(context.Context, *GetUserImpactRequest) (*GetUserImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserImpact not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetCommunityImpact(context.Context, *GetCommunityImpactRequest) (*GetCommunityImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunityImpact not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenges not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) JoinChallenge(context.Context, *JoinChallengeRequest) (*JoinChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChallenge not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) UpdateChallengeProgress(context.Context, *UpdateChallengeProgressRequest) (*UpdateChallengeProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChallengeProgress not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetUserChallenges(context.Context, *GetUserChallengesRequest) (*GetUserChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserChallenges not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetUsersLeaderboard(context.Context, *GetUsersLeaderboardRequest) (*GetUsersLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersLeaderboard not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) GetCommunitiesLeaderboard(context.Context, *GetCommunitiesLeaderboardRequest) (*GetCommunitiesLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunitiesLeaderboard not implemented")
}
func (UnimplementedSustainabilityimpactServiceServer) mustEmbedUnimplementedSustainabilityimpactServiceServer() {
}

// UnsafeSustainabilityimpactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SustainabilityimpactServiceServer will
// result in compilation errors.
type UnsafeSustainabilityimpactServiceServer interface {
	mustEmbedUnimplementedSustainabilityimpactServiceServer()
}

func RegisterSustainabilityimpactServiceServer(s grpc.ServiceRegistrar, srv SustainabilityimpactServiceServer) {
	s.RegisterService(&SustainabilityimpactService_ServiceDesc, srv)
}

func _SustainabilityimpactService_LogImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).LogImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/LogImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).LogImpact(ctx, req.(*LogImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetUserImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetUserImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetUserImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetUserImpact(ctx, req.(*GetUserImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetCommunityImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetCommunityImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetCommunityImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetCommunityImpact(ctx, req.(*GetCommunityImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetChallenges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetChallenges(ctx, req.(*GetChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_JoinChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).JoinChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/JoinChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).JoinChallenge(ctx, req.(*JoinChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_UpdateChallengeProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChallengeProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).UpdateChallengeProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/UpdateChallengeProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).UpdateChallengeProgress(ctx, req.(*UpdateChallengeProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetUserChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetUserChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetUserChallenges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetUserChallenges(ctx, req.(*GetUserChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetUsersLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetUsersLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetUsersLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetUsersLeaderboard(ctx, req.(*GetUsersLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityimpactService_GetCommunitiesLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunitiesLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityimpactServiceServer).GetCommunitiesLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SustainabilityImpactService.SustainabilityimpactService/GetCommunitiesLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityimpactServiceServer).GetCommunitiesLeaderboard(ctx, req.(*GetCommunitiesLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SustainabilityimpactService_ServiceDesc is the grpc.ServiceDesc for SustainabilityimpactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SustainabilityimpactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SustainabilityImpactService.SustainabilityimpactService",
	HandlerType: (*SustainabilityimpactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogImpact",
			Handler:    _SustainabilityimpactService_LogImpact_Handler,
		},
		{
			MethodName: "GetUserImpact",
			Handler:    _SustainabilityimpactService_GetUserImpact_Handler,
		},
		{
			MethodName: "GetCommunityImpact",
			Handler:    _SustainabilityimpactService_GetCommunityImpact_Handler,
		},
		{
			MethodName: "GetChallenges",
			Handler:    _SustainabilityimpactService_GetChallenges_Handler,
		},
		{
			MethodName: "JoinChallenge",
			Handler:    _SustainabilityimpactService_JoinChallenge_Handler,
		},
		{
			MethodName: "UpdateChallengeProgress",
			Handler:    _SustainabilityimpactService_UpdateChallengeProgress_Handler,
		},
		{
			MethodName: "GetUserChallenges",
			Handler:    _SustainabilityimpactService_GetUserChallenges_Handler,
		},
		{
			MethodName: "GetUsersLeaderboard",
			Handler:    _SustainabilityimpactService_GetUsersLeaderboard_Handler,
		},
		{
			MethodName: "GetCommunitiesLeaderboard",
			Handler:    _SustainabilityimpactService_GetCommunitiesLeaderboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Sustainability-impact-service.proto",
}
